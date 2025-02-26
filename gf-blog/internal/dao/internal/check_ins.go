// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckInsDao is the data access object for the table check_ins.
type CheckInsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CheckInsColumns // columns contains all the column names of Table for convenient usage.
}

// CheckInsColumns defines and stores column names for the table check_ins.
type CheckInsColumns struct {
	Id          string //
	UserId      string // 用户ID
	CheckInDate string // 签到日期
	CreatedAt   string // 创建时间
}

// checkInsColumns holds the columns for the table check_ins.
var checkInsColumns = CheckInsColumns{
	Id:          "id",
	UserId:      "user_id",
	CheckInDate: "check_in_date",
	CreatedAt:   "created_at",
}

// NewCheckInsDao creates and returns a new DAO object for table data access.
func NewCheckInsDao() *CheckInsDao {
	return &CheckInsDao{
		group:   "default",
		table:   "check_ins",
		columns: checkInsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CheckInsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CheckInsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CheckInsDao) Columns() CheckInsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CheckInsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CheckInsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CheckInsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
