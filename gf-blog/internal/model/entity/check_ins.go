// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckIns is the golang structure for table check_ins.
type CheckIns struct {
	Id          uint        `json:"id"          orm:"id"            description:""`     //
	UserId      uint        `json:"userId"      orm:"user_id"       description:"用户ID"` // 用户ID
	CheckInDate *gtime.Time `json:"checkInDate" orm:"check_in_date" description:"签到日期"` // 签到日期
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"` // 创建时间
}
