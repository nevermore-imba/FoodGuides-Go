package svc

import (
	"FoodGuides/service/foodmanage/api/internal/config"
	"FoodGuides/service/foodmanage/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	FoodModel     model.FoodModel
	UserFoodModel model.UserFoodModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FoodModel:     model.NewFoodModel(sqlx.NewMysql(c.Mysql.DataSource)),
		UserFoodModel: model.NewUserFoodModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
