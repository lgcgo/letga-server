// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthAccess is the golang structure of table auth_access for DAO operations like Where/Data.
type AuthAccess struct {
	g.Meta   `orm:"table:auth_access, do:true"`
	Id       interface{} // ID
	RoleId   interface{} // 角色ID
	UserId   interface{} // 用户ID
	Status   interface{} // 状态
	CreateAt *gtime.Time // 创建日期
}
