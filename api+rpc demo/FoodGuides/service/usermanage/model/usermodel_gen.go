// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByName(ctx context.Context, name string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`       // 用户Id
		Name       string    `db:"name"`     // 用户名称
		Password   string    `db:"password"` // 用户密码
		Email      string    `db:"email"`    // 用户邮箱
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) withSession(session sqlx.Session) *defaultUserModel {
	return &defaultUserModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Password, data.Email)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Email, newData.Id)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
