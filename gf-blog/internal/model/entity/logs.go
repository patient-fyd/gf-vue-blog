// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure for table logs.
type Logs struct {
	Id          uint        `json:"id"          orm:"id"          description:""`     //
	UserId      uint        `json:"userId"      orm:"user_id"     description:"用户ID"` // 用户ID
	Action      string      `json:"action"      orm:"action"      description:"操作类型"` // 操作类型
	Description string      `json:"description" orm:"description" description:"操作描述"` // 操作描述
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"操作时间"` // 操作时间
}
