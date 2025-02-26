// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PostTags is the golang structure of table post_tags for DAO operations like Where/Data.
type PostTags struct {
	g.Meta `orm:"table:post_tags, do:true"`
	PostId interface{} // 文章ID
	TagId  interface{} // 标签ID
}
