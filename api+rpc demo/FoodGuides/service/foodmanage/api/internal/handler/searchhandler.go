package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/foodmanage/api/internal/logic"
	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(Ctx(r.Context()), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "搜索成功"))
		}
	}
}
