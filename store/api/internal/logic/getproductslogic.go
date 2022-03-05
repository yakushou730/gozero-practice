package logic

import (
	"context"

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
	// temporary hardcoded for development
	return &types.ProductListResponse{
		Products: []types.ProductResponse{
			{
				Id:   1,
				Name: "Product 1",
			},
			{
				Id:   2,
				Name: "Product 2",
			},
			{
				Id:   3,
				Name: "Product 3",
			},
			{
				Id:   4,
				Name: "Product 4",
			},
			{
				Id:   5,
				Name: "Product 5",
			},
		},
	}, nil
}
