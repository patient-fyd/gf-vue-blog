// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostsDao is the data access object for the table posts.
type PostsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns PostsColumns // columns contains all the column names of Table for convenient usage.
}

// PostsColumns defines and stores column names for the table posts.
type PostsColumns struct {
	Id         string //
	Title      string // 文章标题
	Content    string // 文章内容
	UserId     string // 作者ID
	CategoryId string // 分类ID
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Status     string // 文章状态
}

// postsColumns holds the columns for the table posts.
var postsColumns = PostsColumns{
	Id:         "id",
	Title:      "title",
	Content:    "content",
	UserId:     "user_id",
	CategoryId: "category_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Status:     "status",
}

// NewPostsDao creates and returns a new DAO object for table data access.
func NewPostsDao() *PostsDao {
	return &PostsDao{
		group:   "default",
		table:   "posts",
		columns: postsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PostsDao) Columns() PostsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PostsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
