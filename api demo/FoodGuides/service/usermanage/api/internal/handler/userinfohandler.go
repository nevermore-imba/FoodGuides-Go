package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/usermanage/api/internal/logic"
	"FoodGuides/service/usermanage/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "获取成功"))
		}
	}
}
