// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure for table comments.
type Comments struct {
	Id        uint        `json:"id"        orm:"id"         description:""`     //
	PostId    uint        `json:"postId"    orm:"post_id"    description:"文章ID"` // 文章ID
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户ID"` // 用户ID
	Content   string      `json:"content"   orm:"content"    description:"评论内容"` // 评论内容
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
}
