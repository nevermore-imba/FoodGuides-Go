package logic

import (
	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"
	"FoodGuides/service/usermanage/rpc/userclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	token := types.JwtToken{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		RefreshAfter: res.RefreshAfter,
	}

	response := types.UserReply{
		Id:       res.Id,
		Username: res.Username,
		Email:    res.Email,
		JwtToken: token,
	}

	return &types.LoginResponse{UserReply: response}, nil
}
