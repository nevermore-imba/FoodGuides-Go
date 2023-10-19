package logic

import (
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

	userFoods, err := l.svcCtx.UserFoodModel.FindManyByUserid(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	var foodIds []string
	for _, food := range userFoods {
		foodIds = append(foodIds, strconv.FormatInt(food.Foodid.Int64, 10))
	}

	foods, err := l.svcCtx.FoodModel.FindMany(l.ctx, foodIds)
	if err != nil {
		return nil, err
	}

	var foodReplays []types.FoodReply
	for _, food := range foods {
		foodReply := types.FoodReply{
			Id:           strconv.FormatInt(food.Id, 10),
			Name:         food.Name,
			Protein:      food.Protein,
			Fat:          food.Fat,
			Carbohydrate: food.Carbohydrate,
			Calorie:      food.Calorie,
			Minerals:     food.Minerals,
			Calcium:      food.Calcium,
			Phosphorus:   food.Phosphorus,
			Iron:         food.Iron,
			Purine:       food.Purine,
		}
		foodReplays = append(foodReplays, foodReply)
	}

	return &types.FoodListResponse{List: foodReplays}, nil
}
