// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure of table logs for DAO operations like Where/Data.
type Logs struct {
	g.Meta      `orm:"table:logs, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户ID
	Action      interface{} // 操作类型
	Description interface{} // 操作描述
	CreatedAt   *gtime.Time // 操作时间
}
