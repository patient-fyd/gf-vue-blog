// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogsDao is the data access object for the table logs.
type LogsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns LogsColumns // columns contains all the column names of Table for convenient usage.
}

// LogsColumns defines and stores column names for the table logs.
type LogsColumns struct {
	Id          string //
	UserId      string // 用户ID
	Action      string // 操作类型
	Description string // 操作描述
	CreatedAt   string // 操作时间
}

// logsColumns holds the columns for the table logs.
var logsColumns = LogsColumns{
	Id:          "id",
	UserId:      "user_id",
	Action:      "action",
	Description: "description",
	CreatedAt:   "created_at",
}

// NewLogsDao creates and returns a new DAO object for table data access.
func NewLogsDao() *LogsDao {
	return &LogsDao{
		group:   "default",
		table:   "logs",
		columns: logsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LogsDao) Columns() LogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *LogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
