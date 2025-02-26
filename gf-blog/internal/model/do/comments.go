// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure of table comments for DAO operations like Where/Data.
type Comments struct {
	g.Meta    `orm:"table:comments, do:true"`
	Id        interface{} //
	PostId    interface{} // 文章ID
	UserId    interface{} // 用户ID
	Content   interface{} // 评论内容
	CreatedAt *gtime.Time // 创建时间
}
