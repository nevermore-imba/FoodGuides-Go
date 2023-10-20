package svc

import (
	"FoodGuides/service/foodmanage/api/internal/config"
	"FoodGuides/service/foodmanage/rpc/foodclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FoodRpc foodclient.Food
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FoodRpc: foodclient.NewFood(zrpc.MustNewClient(c.FoodRpc)),
	}
}
