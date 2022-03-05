package logic

import (
	"context"

	"github.com/yakushou730/gozero-practice/product/rpc/product"

	"github.com/yakushou730/gozero-practice/store/api/internal/svc"
	"github.com/yakushou730/gozero-practice/store/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductsLogic {
	return GetProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsLogic) GetProducts(req types.ProductListRequest) (*types.ProductListResponse, error) {
	productListResponse, err := l.svcCtx.ProductRpc.ProductList(l.ctx, &product.ProductListRequest{})
	if err != nil {
		return nil, err
	}

	var productList []types.ProductResponse
	for _, p := range productListResponse.ProductItems {
		productList = append(productList, types.ProductResponse{
			Id:   p.GetId(),
			Name: p.GetName(),
		})
	}

	return &types.ProductListResponse{
		Products: productList,
	}, nil
}
