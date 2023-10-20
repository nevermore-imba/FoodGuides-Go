package svc

import (
	"FoodGuides/service/usermanage/api/internal/config"
	"FoodGuides/service/usermanage/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.MySql.DataSource)),
	}
}
