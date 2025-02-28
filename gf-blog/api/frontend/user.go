package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserRegisterReq 用户注册请求
type UserRegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"用户" summary:"用户注册"`
	Username string `json:"username" v:"required|length:4,50#用户名不能为空|用户名长度应在4到50之间"`
	Password string `json:"password" v:"required|length:6,30#密码不能为空|密码长度应在6到30之间"`
	Email    string `json:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
}

// UserRegisterRes 用户注册响应
type UserRegisterRes struct {
	Id uint `json:"id"`
}

// UserUpdateProfileReq 更新个人信息请求
type UserUpdateProfileReq struct {
	g.Meta   `path:"/user/profile/update" method:"put" tags:"用户" summary:"更新个人信息"`
	Email    string `json:"email" v:"email#邮箱格式不正确"`
	Password string `json:"password" v:"length:6,30#密码长度应在6到30之间"`
}

// UserUpdateProfileRes 更新个人信息响应
type UserUpdateProfileRes struct{}
