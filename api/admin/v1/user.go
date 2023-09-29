package v1

import (
	"letga/api/admin/v1/model"

	"github.com/gogf/gf/v2/frame/g"
)

/**
* 用户管理
**/

// 获取用户
type UserGetReq struct {
	g.Meta `path:"/user" method:"get" tags:"Admin@UserService" summary:"Get user"`
	Key    string `json:"key" v:"required"`
}
type UserGetRes struct {
	model.User
}

// 创建用户
type UserCreateReq struct {
	g.Meta    `path:"/user" method:"post" tags:"Admin@UserService" summary:"Create user"`
	Account   string `json:"account" v:"required|length:4,20|regex:^[a-zA-Z][a-zA-Z0-9_]+$"`      // 账号
	Password  string `json:"password" v:"required|password"`                                      // 密码
	Nickname  string `json:"nickname" v:"required|length:2,12|regex:^[\u4E00-\u9FA5A-Za-z0-9]+$"` // 昵称
	Mobile    string `json:"mobile" v:"phone"`                                                    // 手机号
	Email     string `json:"email" v:"email"`                                                     // 电子邮箱
	Avatar    string `json:"avatar"`                                                              // 头像
	Signature string `json:"signature"`                                                           // 个性签名
}
type UserCreateRes struct {
	model.User
}

// 修改用户
type UserUpdateReq struct {
	g.Meta    `path:"/user" method:"put" tags:"Admin@UserService" summary:"Update user"`
	Key       string `json:"key" v:"required"`                                                    // 表索引
	Account   string `json:"account" v:"required|length:4,20|regex:^[a-zA-Z][a-zA-Z0-9_]+$"`      // 账号
	Password  string `json:"password" v:"password"`                                               // 密码
	Nickname  string `json:"nickname" v:"required|length:2,12|regex:^[\u4E00-\u9FA5A-Za-z0-9]+$"` // 昵称
	Mobile    string `json:"mobile" v:"phone"`                                                    // 手机号
	Email     string `json:"email" v:"email"`                                                     // 电子邮箱
	Avatar    string `json:"avatar"`                                                              // 头像
	Signature string `json:"signature"`                                                           // 个性签名
}
type UserUpdateRes struct {
	model.User
}

// 设置用户状态
type UserSetStatusReq struct {
	g.Meta `path:"/user/status" method:"put" tags:"Admin@UserService" summary:"Set user status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type UserSetStatusRes struct {
	model.User
}

// 删除用户
type UserDeleteReq struct {
	g.Meta `path:"/user" method:"delete" tags:"Admin@UserService" summary:"Delete user"`
	Keys   []string `json:"keys" v:"required"`
}
type UserDeleteRes struct{}

// 用户分页
type UserGetPageReq struct {
	g.Meta `path:"/users" method:"get" tags:"Admin@UserService" summary:"Get user page"`
	model.PageParams
	Key        string `json:"key"`        // 搜索字段,索引
	Uuid       string `json:"uuid"`       // 搜索字段,唯一ID
	Account    string `json:"account"`    // 搜索字段,账号
	Nickname   string `json:"nickname"`   // 搜索字段,昵称
	Mobile     string `json:"mobile"`     // 搜索字段,手机号
	Email      string `json:"email"`      // 搜索字段,电子邮箱
	SigninRole string `json:"signinRole"` // 搜索字段,登录角色
	SigninIp   string `json:"signinIp"`   // 搜索字段,登录IP
	Status     string `json:"status"`     // 搜索字段,状态

	// 上下文场景
	CtxScene string `json:"ctxScene" v:"in:mainTable,accessSetupTable"`
}
type UserGetPageRes struct {
	Data []*model.User `json:"data,omitempty"`
	model.Page
}
