// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta           `orm:"table:users, do:true"`
	Id               interface{} //
	Username         interface{} // 用户名，唯一
	Password         interface{} // 用户密码
	Email            interface{} // 用户邮箱，唯一
	CreatedAt        *gtime.Time // 用户创建时间
	UpdatedAt        *gtime.Time // 用户更新时间
	Role             interface{} // 用户角色，默认是普通用户
	ResetToken       interface{} // 密码重置令牌
	ResetTokenExpiry *gtime.Time // 密码重置令牌过期时间
	VerificationCode interface{} // 邮箱验证码
}
