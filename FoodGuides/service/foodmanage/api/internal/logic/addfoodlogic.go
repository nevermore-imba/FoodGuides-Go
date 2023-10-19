package logic

import (
	"FoodGuides/service/foodmanage/model"
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

	foodId, _ := strconv.ParseInt(req.FoodId, 10, 64)

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

	return &types.AddFoodResponse{}, nil
}
