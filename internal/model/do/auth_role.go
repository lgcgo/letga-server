// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRole is the golang structure of table auth_role for DAO operations like Where/Data.
type AuthRole struct {
	g.Meta   `orm:"table:auth_role, do:true"`
	Id       interface{} // ID
	ParentId interface{} // 父ID
	Title    interface{} // 标题
	Name     interface{} // 名称
	Status   interface{} // 状态
	Weight   interface{} // 权重
	CreateAt *gtime.Time // 创建日期
	UpdateAt *gtime.Time // 修改日期
}
