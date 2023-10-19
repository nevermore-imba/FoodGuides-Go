package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/foodmanage/api/internal/logic"
	"FoodGuides/service/foodmanage/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FoodListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFoodListLogic(r.Context(), svcCtx)
		resp, err := l.FoodList()
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "请求成功"))
		}
	}
}
