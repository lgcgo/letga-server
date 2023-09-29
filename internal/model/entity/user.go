// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            uint        `json:"id"            ` // ID
	Uuid          string      `json:"uuid"          ` // 唯一ID
	Account       string      `json:"account"       ` // 账号
	Mobile        string      `json:"mobile"        ` // 手机号
	Email         string      `json:"email"         ` // 电子邮箱
	Password      string      `json:"password"      ` // 密码
	Salt          string      `json:"salt"          ` // 密码盐
	Nickname      string      `json:"nickname"      ` // 昵称
	Avatar        string      `json:"avatar"        ` // 头像
	Signature     string      `json:"signature"     ` // 个性签名
	SigninRole    string      `json:"signinRole"    ` // 登录角色
	SigninFailure uint        `json:"signinFailure" ` // 失败次数
	SigninIp      string      `json:"signinIp"      ` // 登录IP
	SigninAt      *gtime.Time `json:"signinAt"      ` // 登录日期
	Status        string      `json:"status"        ` // 状态
	CreateAt      *gtime.Time `json:"createAt"      ` // 创建日期
	UpdateAt      *gtime.Time `json:"updateAt"      ` // 更新日期
	DeletedAt     *gtime.Time `json:"deletedAt"     ` // 删除日期
}
