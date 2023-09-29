package model

import (
	"letga/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

// 创建用户
type UserCreateInput struct {
	Account  string
	Password string
	Nickname string
	Avatar   string
	Mobile   string
	Email    string
}

// 获取用户
type User struct {
	gmeta.Meta `orm:"table:user"`
	Key        string `json:"key"` // 索引
	entity.User
}

// 用户搜索
type UserSearch struct {
	gmeta.Meta `search:"account,mobile,email,nickname"`
	Key        string `f:"hashFc"` // 索引
	Uuid       string `f:"equal"`  // 唯一ID
	Account    string `f:"like"`   // 账号
	Nickname   string `f:"like"`   // 昵称
	Mobile     string `f:"like"`   // 手机号
	Email      string `f:"like"`   // 电子邮箱
	SigninRole string `f:"equal"`  // 登录角色
	SigninIp   string `f:"like"`   // 登录IP
	Status     string `f:"equal"`  // 状态
}
type UserPageInput struct {
	PageParams
	UserSearch
	CtxScene string
}
type UserPageOutput struct {
	Data []*User
	Page
}

// 修改用户
type UserUpdateInput struct {
	Key       string
	Account   string
	Password  string
	Nickname  string
	Avatar    string
	Mobile    string
	Email     string
	Signature string
}

// 用户登录
type UserSigninInput struct {
	Passport string
	Password string
	Captcha  string
	Type     string
	Role     string
}

// 密码登录
type UserSigninPassportInput struct {
	Passport string
	Password string
	Captcha  string
	Role     string
}

// 手机号登录
type UserSigninMobileInput struct {
	Mobile  string
	Captcha string
}

// 邮箱登录
type UserSigninEmailInput struct {
	Email   string
	Captcha string
}
