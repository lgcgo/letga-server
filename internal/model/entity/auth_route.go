// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoute is the golang structure for table auth_route.
type AuthRoute struct {
	Id       uint        `json:"id"       ` // ID
	MenuId   uint        `json:"menuId"   ` // 菜单ID
	Title    string      `json:"title"    ` // 标题
	Path     string      `json:"path"     ` // 路由地址
	Method   string      `json:"method"   ` // 请求方法
	Remark   string      `json:"remark"   ` // 备注
	Status   string      `json:"status"   ` // 状态
	Weight   int         `json:"weight"   ` // 权重
	CreateAt *gtime.Time `json:"createAt" ` // 创建日期
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新日期
}
