// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure for table posts.
type Posts struct {
	Id         uint        `json:"id"         orm:"id"          description:""`     //
	Title      string      `json:"title"      orm:"title"       description:"文章标题"` // 文章标题
	Content    string      `json:"content"    orm:"content"     description:"文章内容"` // 文章内容
	UserId     uint        `json:"userId"     orm:"user_id"     description:"作者ID"` // 作者ID
	CategoryId uint        `json:"categoryId" orm:"category_id" description:"分类ID"` // 分类ID
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"` // 更新时间
	Status     string      `json:"status"     orm:"status"      description:"文章状态"` // 文章状态
}
