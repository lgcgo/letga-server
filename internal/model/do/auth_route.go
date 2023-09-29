// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoute is the golang structure of table auth_route for DAO operations like Where/Data.
type AuthRoute struct {
	g.Meta   `orm:"table:auth_route, do:true"`
	Id       interface{} // ID
	MenuId   interface{} // 菜单ID
	Title    interface{} // 标题
	Path     interface{} // 路由地址
	Method   interface{} // 请求方法
	Remark   interface{} // 备注
	Status   interface{} // 状态
	Weight   interface{} // 权重
	CreateAt *gtime.Time // 创建日期
	UpdateAt *gtime.Time // 更新日期
}
