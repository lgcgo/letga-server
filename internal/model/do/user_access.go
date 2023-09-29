// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccess is the golang structure of table user_access for DAO operations like Where/Data.
type UserAccess struct {
	g.Meta   `orm:"table:user_access, do:true"`
	Id       interface{} // ID
	UserId   interface{} // 用户ID
	RoleId   interface{} // 角色ID
	Status   interface{} // 状态
	CreateAt *gtime.Time // 创建日期
}
