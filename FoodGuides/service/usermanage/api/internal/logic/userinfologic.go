package logic

import (
	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"
	"FoodGuides/service/usermanage/model"
	"context"
	"encoding/json"
	"errors"

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

	// 查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	response := types.UserReply{
		Id:       user.Id,
		Username: user.Name,
		Email:    user.Email,
	}

	return &types.UserInfoResponse{UserReply: response}, nil
}
