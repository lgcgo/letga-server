// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id       uint        `json:"id"       ` // ID
	ParentId uint        `json:"parentId" ` // 父ID
	Title    string      `json:"title"    ` // 标题
	Icon     string      `json:"icon"     ` // 图标
	CoverUrl string      `json:"coverUrl" ` // 封面图片
	Remark   string      `json:"remark"   ` // 描述
	Status   string      `json:"status"   ` // 状态
	Weight   int         `json:"weight"   ` // 权重
	CreateAt *gtime.Time `json:"createAt" ` // 创建日期
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新日期
}
