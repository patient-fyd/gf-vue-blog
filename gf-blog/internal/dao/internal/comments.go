// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentsDao is the data access object for the table comments.
type CommentsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CommentsColumns // columns contains all the column names of Table for convenient usage.
}

// CommentsColumns defines and stores column names for the table comments.
type CommentsColumns struct {
	Id        string //
	PostId    string // 文章ID
	UserId    string // 用户ID
	Content   string // 评论内容
	CreatedAt string // 创建时间
}

// commentsColumns holds the columns for the table comments.
var commentsColumns = CommentsColumns{
	Id:        "id",
	PostId:    "post_id",
	UserId:    "user_id",
	Content:   "content",
	CreatedAt: "created_at",
}

// NewCommentsDao creates and returns a new DAO object for table data access.
func NewCommentsDao() *CommentsDao {
	return &CommentsDao{
		group:   "default",
		table:   "comments",
		columns: commentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommentsDao) Columns() CommentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
