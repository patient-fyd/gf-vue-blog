// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure of table permissions for DAO operations like Where/Data.
type Permissions struct {
	g.Meta         `orm:"table:permissions, do:true"`
	Id             interface{} //
	PermissionName interface{} // 权限名称
	CreatedAt      *gtime.Time // 创建时间
}
