package logic

import (
	"context"
	"strconv"

	"FoodGuides/service/foodmanage/rpc/food"
	"FoodGuides/service/foodmanage/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FoodListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFoodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FoodListLogic {
	return &FoodListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FoodListLogic) FoodList(in *food.FoodListRequest) (*food.FoodListResponse, error) {
	uid, _ := strconv.ParseInt(in.Userid, 10, 64)
	userFoods, err := l.svcCtx.UserFoodModel.FindManyByUserid(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	var foodIds []string
	for _, f := range userFoods {
		foodIds = append(foodIds, strconv.FormatInt(f.Foodid.Int64, 10))
	}

	foods, err := l.svcCtx.FoodModel.FindMany(l.ctx, foodIds)
	if err != nil {
		return nil, err
	}

	var foodList []*food.FoodInfoResponse
	for _, f := range foods {
		item := food.FoodInfoResponse{
			Id:           strconv.FormatInt(f.Id, 10),
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
		foodList = append(foodList, &item)
	}

	return &food.FoodListResponse{
		Data: foodList,
	}, nil
}
