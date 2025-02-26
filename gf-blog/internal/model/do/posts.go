// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure of table posts for DAO operations like Where/Data.
type Posts struct {
	g.Meta     `orm:"table:posts, do:true"`
	Id         interface{} //
	Title      interface{} // 文章标题
	Content    interface{} // 文章内容
	UserId     interface{} // 作者ID
	CategoryId interface{} // 分类ID
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Status     interface{} // 文章状态
}
