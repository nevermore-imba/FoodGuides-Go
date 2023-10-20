package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"FoodGuides/service/foodmanage/api/internal/svc"
	"FoodGuides/service/foodmanage/api/internal/types"

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
	foodId, _ := strconv.ParseInt(req.FoodId, 10, 64)
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

	return &types.DeleteFoodResponse{}, nil
}
