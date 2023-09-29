// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta   `orm:"table:menu, do:true"`
	Id       interface{} // ID
	ParentId interface{} // 父ID
	Title    interface{} // 标题
	Icon     interface{} // 图标
	CoverUrl interface{} // 封面图片
	Remark   interface{} // 描述
	Status   interface{} // 状态
	Weight   interface{} // 权重
	CreateAt *gtime.Time // 创建日期
	UpdateAt *gtime.Time // 更新日期
}
