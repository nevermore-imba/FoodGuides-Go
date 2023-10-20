package logic

import (
	"FoodGuides/service/foodmanage/rpc/food"
	"context"
	"encoding/json"
	"strconv"

	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FoodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFoodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FoodListLogic {
	return &FoodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FoodListLogic) FoodList() (*types.FoodListResponse, error) {
	// 获取 jwt 载体中 `uid` 信息，
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.FoodRpc.FoodList(l.ctx, &food.FoodListRequest{
		Userid: strconv.FormatInt(uid, 10),
	})
	if err != nil {
		return nil, err
	}

	var foodReplays []types.FoodReply
	for _, f := range res.Data {
		foodReply := types.FoodReply{
			Id:           f.Id,
			Name:         f.Name,
			Protein:      f.Protein,
			Fat:          f.Fat,
			Carbohydrate: f.Carbohydrate,
			Calorie:      f.Calorie,
			Minerals:     f.Minerals,
			Calcium:      f.Calcium,
			Phosphorus:   f.Phosphorus,
			Iron:         f.Iron,
			Purine:       f.Purine,
		}
		foodReplays = append(foodReplays, foodReply)
	}

	return &types.FoodListResponse{List: foodReplays}, nil
}
