package logic

import (
	"context"
	"strconv"

	"FoodGuides/service/foodmanage/rpc/food"
	"FoodGuides/service/foodmanage/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *food.SearchRequest) (*food.FoodInfoResponse, error) {
	res, err := l.svcCtx.FoodModel.FindOneByName(l.ctx, in.SearchKey)
	if err != nil {
		return nil, err
	}

	return &food.FoodInfoResponse{
		Id:           strconv.FormatInt(res.Id, 10),
		Name:         res.Name,
		Protein:      res.Protein,
		Fat:          res.Fat,
		Carbohydrate: res.Carbohydrate,
		Calorie:      res.Calorie,
		Minerals:     res.Minerals,
		Calcium:      res.Calcium,
		Phosphorus:   res.Phosphorus,
		Iron:         res.Iron,
		Purine:       res.Purine,
	}, nil
}
