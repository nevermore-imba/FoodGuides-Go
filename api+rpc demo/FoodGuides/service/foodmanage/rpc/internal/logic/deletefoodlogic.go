package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"FoodGuides/service/foodmanage/rpc/food"
	"FoodGuides/service/foodmanage/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoodLogic {
	return &DeleteFoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFoodLogic) DeleteFood(in *food.DeleteFoodRequest) (*food.StatusResponse, error) {
	uid, _ := strconv.ParseInt(in.Userid, 10, 64)
	foodId, _ := strconv.ParseInt(in.FoodId, 10, 64)
	userFood, _ := l.svcCtx.UserFoodModel.FindOneByUserid(l.ctx, sql.NullInt64{
		Int64: uid,
		Valid: true,
	})

	if userFood == nil {
		return nil, errors.New(fmt.Sprintf("该用户名下没有关联的食物，用户 ID：%d", uid))
	}

	if userFood.Foodid.Int64 != foodId {
		return nil, errors.New(fmt.Sprintf("该用户名下没有此关联的食物，用户 ID：%d， 食物 ID: %d", uid, foodId))
	}

	err := l.svcCtx.UserFoodModel.Delete(l.ctx, userFood.Id)

	if err != nil {
		return nil, err
	}

	return &food.StatusResponse{
		Success: 1,
	}, nil
}
