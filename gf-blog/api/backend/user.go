package backend

import (
	"gf-blog/api"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	g.Meta   `path:"/user/create" method:"post" tags:"用户管理" summary:"创建用户"`
	Username string `json:"username" v:"required|length:4,50#用户名不能为空|用户名长度应在4到50之间"`
	Password string `json:"password" v:"required|length:6,30#密码不能为空|密码长度应在6到30之间"`
	Email    string `json:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
	Role     string `json:"role" v:"required|in:user,admin#角色不能为空|角色只能是user或admin"`
}

// UserCreateRes 创建用户响应
type UserCreateRes struct {
	Id uint `json:"id"`
}

// UserUpdateReq 更新用户请求
type UserUpdateReq struct {
	g.Meta   `path:"/user/update" method:"put" tags:"用户管理" summary:"更新用户"`
	Id       uint   `json:"id" v:"required#用户ID不能为空"`
	Username string `json:"username" v:"length:4,50#用户名长度应在4到50之间"`
	Password string `json:"password" v:"length:6,30#密码长度应在6到30之间"`
	Email    string `json:"email" v:"email#邮箱格式不正确"`
	Role     string `json:"role" v:"in:user,admin#角色只能是user或admin"`
}

// UserUpdateRes 更新用户响应
type UserUpdateRes struct {
	Id uint `json:"id"`
}

// UserDeleteReq 删除用户请求
type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"用户管理" summary:"删除用户"`
	Id     uint `json:"id" v:"required#用户ID不能为空"`
}

// UserDeleteRes 删除用户响应
type UserDeleteRes struct{}

// UserGetReq 获取用户请求
type UserGetReq struct {
	g.Meta `path:"/user/get" method:"get" tags:"用户管理" summary:"获取用户信息"`
	Id     uint `json:"id" v:"required#用户ID不能为空"`
}

// UserGetRes 获取用户响应
type UserGetRes struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserListReq 获取用户列表请求
type UserListReq struct {
	g.Meta   `path:"/user/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	api.CommonPaginationReq
}

// UserListRes 获取用户列表响应
type UserListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
