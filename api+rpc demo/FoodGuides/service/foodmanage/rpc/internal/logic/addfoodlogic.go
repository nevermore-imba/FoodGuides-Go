package logic

import (
	"FoodGuides/service/foodmanage/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"FoodGuides/service/foodmanage/rpc/food"
	"FoodGuides/service/foodmanage/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFoodLogic {
	return &AddFoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFoodLogic) AddFood(in *food.AddFoodRequest) (*food.StatusResponse, error) {
	uid, _ := strconv.ParseInt(in.Userid, 10, 64)
	foodId, _ := strconv.ParseInt(in.FoodId, 10, 64)

	_, err := l.svcCtx.FoodModel.FindOne(l.ctx, foodId)

	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("不存在 ID 为 %d 的食材", foodId))
		}
		return nil, err
	}

	data := model.UserFood{
		Userid: sql.NullInt64{
			Int64: uid,
			Valid: true,
		},
		Foodid: sql.NullInt64{
			Int64: foodId,
			Valid: true,
		},
	}

	_, err = l.svcCtx.UserFoodModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &food.StatusResponse{
		Success: 1,
	}, nil
}
