// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id               uint        `json:"id"               orm:"id"                 description:""`             //
	Username         string      `json:"username"         orm:"username"           description:"用户名，唯一"`       // 用户名，唯一
	Password         string      `json:"password"         orm:"password"           description:"用户密码"`         // 用户密码
	Email            string      `json:"email"            orm:"email"              description:"用户邮箱，唯一"`      // 用户邮箱，唯一
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"用户创建时间"`       // 用户创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"用户更新时间"`       // 用户更新时间
	Role             string      `json:"role"             orm:"role"               description:"用户角色，默认是普通用户"` // 用户角色，默认是普通用户
	ResetToken       string      `json:"resetToken"       orm:"reset_token"        description:"密码重置令牌"`       // 密码重置令牌
	ResetTokenExpiry *gtime.Time `json:"resetTokenExpiry" orm:"reset_token_expiry" description:"密码重置令牌过期时间"`   // 密码重置令牌过期时间
	VerificationCode string      `json:"verificationCode" orm:"verification_code"  description:"邮箱验证码"`        // 邮箱验证码
}
