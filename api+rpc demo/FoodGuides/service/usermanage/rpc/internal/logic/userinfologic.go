package logic

import (
	"FoodGuides/service/usermanage/model"
	"context"
	"errors"
	"strconv"

	"FoodGuides/service/usermanage/rpc/internal/svc"
	"FoodGuides/service/usermanage/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoRequest) (*user.Response, error) {
	uid, _ := strconv.ParseInt(in.Userid, 10, 64)

	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return &user.Response{
		Id:       res.Id,
		Email:    res.Email,
		Username: res.Name,
	}, nil
}
