package logic

import (
	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"
	"FoodGuides/service/usermanage/rpc/userclient"
	"context"
	"encoding/json"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResponse, error) {
	// 通过 l.ctx.Value("uid") 可获取 jwt 载体中 `uid` 信息
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	user, err := l.svcCtx.UserRpc.Userinfo(l.ctx, &userclient.UserinfoRequest{
		Userid: strconv.FormatInt(uid, 10),
	})
	if err != nil {
		return nil, err
	}

	response := types.UserReply{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return &types.UserInfoResponse{UserReply: response}, nil
}
