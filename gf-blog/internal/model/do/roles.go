// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure of table roles for DAO operations like Where/Data.
type Roles struct {
	g.Meta    `orm:"table:roles, do:true"`
	Id        interface{} //
	RoleName  interface{} // 角色名称
	CreatedAt *gtime.Time // 创建时间
}
