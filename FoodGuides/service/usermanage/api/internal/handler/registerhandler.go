package handler

import (
	"FoodGuides/common/responsex"
	"net/http"

	"FoodGuides/service/usermanage/api/internal/logic"
	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(Ctx(r.Context()), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJson(w, responsex.FailureResponse(nil, err.Error(), 1000))
		} else {
			httpx.OkJson(w, responsex.SuccessResponse(resp, "注册成功"))
		}
	}
}
