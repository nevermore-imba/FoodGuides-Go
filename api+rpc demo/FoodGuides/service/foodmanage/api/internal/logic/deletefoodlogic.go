package logic

import (
	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"
	"FoodGuides/service/foodmanage/rpc/foodclient"
	"context"
	"encoding/json"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoodLogic {
	return &DeleteFoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoodLogic) DeleteFood(req *types.DeleteFoodRequest) (*types.DeleteFoodResponse, error) {
	// 获取 jwt 载体中 `uid` 信息，
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err := l.svcCtx.FoodRpc.DeleteFood(l.ctx, &foodclient.DeleteFoodRequest{
		Userid: strconv.FormatInt(uid, 10),
		FoodId: req.FoodId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteFoodResponse{}, nil
}
