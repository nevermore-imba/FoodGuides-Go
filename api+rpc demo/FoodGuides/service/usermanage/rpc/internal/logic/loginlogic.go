package logic

import (
	"FoodGuides/common/cryptx"
	"FoodGuides/common/jwtx"
	"FoodGuides/service/usermanage/model"
	"context"
	"errors"
	"time"

	"FoodGuides/service/usermanage/rpc/internal/svc"
	"FoodGuides/service/usermanage/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.Response, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, errors.New("密码错误")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.AccessExpire
	jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &user.Response{
		Id:           res.Id,
		Email:        res.Email,
		Username:     res.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
