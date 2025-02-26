// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	Id        uint        `json:"id"        orm:"id"         description:""`     //
	RoleName  string      `json:"roleName"  orm:"role_name"  description:"角色名称"` // 角色名称
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
}
