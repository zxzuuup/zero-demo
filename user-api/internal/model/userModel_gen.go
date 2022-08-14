// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheZeroDemoUserIdPrefix = "cache:zeroDemo:user:id:"
)

type (
	userModel interface {
		TransInsert(ctx context.Context,session sqlx.Session,data *User) (sql.Result, error)
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id uint64) error
		TransCtx(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id       uint64 `db:"id"`
		Nickname string `db:"nickname"`
		Mobile   string `db:"mobile"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id uint64) error {
	zeroDemoUserIdKey := fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, zeroDemoUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id uint64) (*User, error) {
	zeroDemoUserIdKey := fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, zeroDemoUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
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
	zeroDemoUserIdKey := fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Mobile)
	}, zeroDemoUserIdKey)
	return ret, err
}

func (m *defaultUserModel) TransInsert(ctx context.Context,session sqlx.Session,data *User) (sql.Result, error) {
	zeroDemoUserIdKey := fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Nickname, data.Mobile)
	}, zeroDemoUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	zeroDemoUserIdKey := fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Mobile, data.Id)
	}, zeroDemoUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheZeroDemoUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}

//暴露给logic开启事务
func (m *defaultUserModel) TransCtx(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error{
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session)error {
		return fn(ctx,s)
	})
}
