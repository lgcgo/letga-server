// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthAccess is the golang structure for table auth_access.
type AuthAccess struct {
	Id       uint        `json:"id"       ` // ID
	RoleId   uint        `json:"roleId"   ` // 角色ID
	UserId   uint        `json:"userId"   ` // 用户ID
	Status   string      `json:"status"   ` // 状态
	CreateAt *gtime.Time `json:"createAt" ` // 创建日期
}
