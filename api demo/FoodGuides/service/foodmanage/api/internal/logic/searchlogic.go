package logic

import (
	"context"
	"strconv"

	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (*types.SearchResponse, error) {
	searchKey := req.Key
	food, err := l.svcCtx.FoodModel.FindOneByName(l.ctx, searchKey)

	if err != nil {
		return nil, err
	}

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
	return &types.SearchResponse{FoodReply: foodReply}, nil
}
