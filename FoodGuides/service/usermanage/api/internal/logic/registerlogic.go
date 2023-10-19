package logic

import (
	"FoodGuides/common/cryptx"
	"FoodGuides/common/jwtx"
	"FoodGuides/service/usermanage/model"
	"context"
	"errors"
	"time"

	"FoodGuides/service/usermanage/api/internal/svc"
	"FoodGuides/service/usermanage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	// 判断邮箱是否已经被注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err == nil {
		return nil, errors.New("该邮箱已注册")
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	user := model.User{
		Name:     req.Username,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
		Email:    req.Email,
	}

	// 插入一条新的用户数据
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	var jwtToken string
	jwtToken, err = jwtx.GetJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		now,
		accessExpire,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	token := types.JwtToken{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	response := types.UserReply{
		Id:       user.Id,
		Username: user.Name,
		Email:    user.Email,
		JwtToken: token,
	}

	return &types.RegisterResponse{UserReply: response}, nil
}
