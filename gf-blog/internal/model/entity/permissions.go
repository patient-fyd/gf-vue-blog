// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure for table permissions.
type Permissions struct {
	Id             uint        `json:"id"             orm:"id"              description:""`     //
	PermissionName string      `json:"permissionName" orm:"permission_name" description:"权限名称"` // 权限名称
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"` // 创建时间
}
