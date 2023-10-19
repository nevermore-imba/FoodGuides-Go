package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ FoodModel = (*customFoodModel)(nil)

type (
	// FoodModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFoodModel.
	FoodModel interface {
		foodModel
		FindMany(ctx context.Context, ids []string) ([]*Food, error)
	}

	customFoodModel struct {
		*defaultFoodModel
	}
)

func (m *customFoodModel) FindMany(ctx context.Context, ids []string) ([]*Food, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (%s)", foodRows, m.table, strings.Join(ids, ","))
	var resp []*Food
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewFoodModel returns a model for the database table.
func NewFoodModel(conn sqlx.SqlConn) FoodModel {
	return &customFoodModel{
		defaultFoodModel: newFoodModel(conn),
	}
}
