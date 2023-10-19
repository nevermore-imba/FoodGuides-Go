package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserFoodModel = (*customUserFoodModel)(nil)

type (
	// UserFoodModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFoodModel.
	UserFoodModel interface {
		userFoodModel
		FindManyByUserid(ctx context.Context, userid int64) ([]*UserFood, error)
	}

	customUserFoodModel struct {
		*defaultUserFoodModel
	}
)

func (m *customUserFoodModel) FindManyByUserid(ctx context.Context, userid int64) ([]*UserFood, error) {
	var resp []*UserFood
	query := fmt.Sprintf("select %s from %s where `userid` = ?", userFoodRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUserFoodModel returns a model for the database table.
func NewUserFoodModel(conn sqlx.SqlConn) UserFoodModel {
	return &customUserFoodModel{
		defaultUserFoodModel: newUserFoodModel(conn),
	}
}
