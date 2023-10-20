package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/foodmanage/api/internal/logic"
	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddFoodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFoodRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(Ctx(r.Context()), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewAddFoodLogic(r.Context(), svcCtx)
		resp, err := l.AddFood(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "新增成功"))
		}
	}
}
