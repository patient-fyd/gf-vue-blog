// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for the table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for the table users.
type UsersColumns struct {
	Id               string //
	Username         string // 用户名，唯一
	Password         string // 用户密码
	Email            string // 用户邮箱，唯一
	CreatedAt        string // 用户创建时间
	UpdatedAt        string // 用户更新时间
	Role             string // 用户角色，默认是普通用户
	ResetToken       string // 密码重置令牌
	ResetTokenExpiry string // 密码重置令牌过期时间
	VerificationCode string // 邮箱验证码
}

// usersColumns holds the columns for the table users.
var usersColumns = UsersColumns{
	Id:               "id",
	Username:         "username",
	Password:         "password",
	Email:            "email",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	Role:             "role",
	ResetToken:       "reset_token",
	ResetTokenExpiry: "reset_token_expiry",
	VerificationCode: "verification_code",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
