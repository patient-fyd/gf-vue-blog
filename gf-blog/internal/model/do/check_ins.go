// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckIns is the golang structure of table check_ins for DAO operations like Where/Data.
type CheckIns struct {
	g.Meta      `orm:"table:check_ins, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户ID
	CheckInDate *gtime.Time // 签到日期
	CreatedAt   *gtime.Time // 创建时间
}
