package handler

import (
	"net/http"

	"github.com/yakushou730/gozero-practice/store/api/internal/logic"
	"github.com/yakushou730/gozero-practice/store/api/internal/svc"
	"github.com/yakushou730/gozero-practice/store/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProductsLogic(r.Context(), svcCtx)
		resp, err := l.GetProducts(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
