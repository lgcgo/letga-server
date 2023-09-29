// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRole is the golang structure for table auth_role.
type AuthRole struct {
	Id       uint        `json:"id"       ` // ID
	ParentId uint        `json:"parentId" ` // 父ID
	Title    string      `json:"title"    ` // 标题
	Name     string      `json:"name"     ` // 名称
	Status   string      `json:"status"   ` // 状态
	Weight   int         `json:"weight"   ` // 权重
	CreateAt *gtime.Time `json:"createAt" ` // 创建日期
	UpdateAt *gtime.Time `json:"updateAt" ` // 修改日期
}
