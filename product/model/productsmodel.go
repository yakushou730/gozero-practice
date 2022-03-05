package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productsFieldNames          = builder.RawFieldNames(&Products{})
	productsRows                = strings.Join(productsFieldNames, ",")
	productsRowsExpectAutoSet   = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	productsRowsWithPlaceHolder = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProductsIdPrefix = "cache:products:id:"
)

type (
	ProductsModel interface {
		Insert(data *Products) (sql.Result, error)
		FindOne(id int64) (*Products, error)
		Update(data *Products) error
		Delete(id int64) error
	}

	defaultProductsModel struct {
		sqlc.CachedConn
		table string
	}

	Products struct {
		Id   int64  `db:"id"`   // id
		Name string `db:"name"` // name
	}
)

func NewProductsModel(conn sqlx.SqlConn, c cache.CacheConf) ProductsModel {
	return &defaultProductsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`products`",
	}
}

func (m *defaultProductsModel) Insert(data *Products) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, productsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name)

	return ret, err
}

func (m *defaultProductsModel) FindOne(id int64) (*Products, error) {
	productsIdKey := fmt.Sprintf("%s%v", cacheProductsIdPrefix, id)
	var resp Products
	err := m.QueryRow(&resp, productsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductsModel) Update(data *Products) error {
	productsIdKey := fmt.Sprintf("%s%v", cacheProductsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productsRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Id)
	}, productsIdKey)
	return err
}

func (m *defaultProductsModel) Delete(id int64) error {

	productsIdKey := fmt.Sprintf("%s%v", cacheProductsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, productsIdKey)
	return err
}

func (m *defaultProductsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProductsIdPrefix, primary)
}

func (m *defaultProductsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
