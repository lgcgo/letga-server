// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} // ID
	Uuid          interface{} // 唯一ID
	Account       interface{} // 账号
	Mobile        interface{} // 手机号
	Email         interface{} // 电子邮箱
	Password      interface{} // 密码
	Salt          interface{} // 密码盐
	Nickname      interface{} // 昵称
	Avatar        interface{} // 头像
	Signature     interface{} // 个性签名
	SigninRole    interface{} // 登录角色
	SigninFailure interface{} // 失败次数
	SigninIp      interface{} // 登录IP
	SigninAt      *gtime.Time // 登录日期
	Status        interface{} // 状态
	CreateAt      *gtime.Time // 创建日期
	UpdateAt      *gtime.Time // 更新日期
	DeletedAt     *gtime.Time // 删除日期
}
