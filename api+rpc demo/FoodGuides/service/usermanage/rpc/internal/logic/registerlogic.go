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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	// 判断邮箱是否已经被注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return nil, errors.New("该邮箱已注册")
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	newUser := model.User{
		Name:     in.Username,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Email:    in.Email,
	}

	// 插入一条新的用户数据
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, err
	}

	newUser.Id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.AccessExpire

	var jwtToken string
	jwtToken, err = jwtx.GetJwtToken(
		l.svcCtx.Config.AccessSecret,
		now,
		accessExpire,
		newUser.Id,
	)
	if err != nil {
		return nil, err
	}

	return &user.Response{
		Id:           newUser.Id,
		Email:        newUser.Email,
		Username:     newUser.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
