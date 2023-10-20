package logic

import (
	"FoodGuides/service/foodmanage/rpc/foodclient"
	"context"
	"encoding/json"
	"strconv"

	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFoodLogic {
	return &AddFoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFoodLogic) AddFood(req *types.AddFoodRequest) (*types.AddFoodResponse, error) {
	// 获取 jwt 载体中 `uid` 信息，
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err := l.svcCtx.FoodRpc.AddFood(l.ctx, &foodclient.AddFoodRequest{
		Userid: strconv.FormatInt(uid, 10),
		FoodId: req.FoodId,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddFoodResponse{}, nil
}
