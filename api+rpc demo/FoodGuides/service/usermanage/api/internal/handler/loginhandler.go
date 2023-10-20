package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/usermanage/api/internal/logic"
	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(Ctx(r.Context()), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "登录成功"))
		}
	}
}
