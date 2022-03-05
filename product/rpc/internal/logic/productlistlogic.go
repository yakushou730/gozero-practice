package logic

import (
	"context"

	"github.com/yakushou730/gozero-practice/product/rpc/internal/svc"
	"github.com/yakushou730/gozero-practice/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *product.ProductListRequest) (*product.ProductListResponse, error) {
	// temporary hardcoded for development
	productItems := []*product.ProductItem{
		{
			Id:   1,
			Name: "product A",
		},
		{
			Id:   2,
			Name: "product B",
		},
	}

	return &product.ProductListResponse{
		ProductItems: productItems,
	}, nil
}
