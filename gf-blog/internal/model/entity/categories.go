// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Categories is the golang structure for table categories.
type Categories struct {
	Id          uint        `json:"id"          orm:"id"          description:""`     //
	Name        string      `json:"name"        orm:"name"        description:"分类名称"` // 分类名称
	Description string      `json:"description" orm:"description" description:"分类描述"` // 分类描述
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"` // 更新时间
}
