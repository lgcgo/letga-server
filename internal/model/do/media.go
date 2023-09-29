// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure of table media for DAO operations like Where/Data.
type Media struct {
	g.Meta    `orm:"table:media, do:true"`
	Id        interface{} // ID
	UserId    interface{} // 用户ID
	Name      interface{} // 文件名
	Path      interface{} // 路径
	Size      interface{} // 大小
	FileType  interface{} // 文件类型
	MimeType  interface{} // MIME类型
	Hash      interface{} // 哈希值
	Extparam  interface{} // 透传数据
	Storage   interface{} // 储存库
	Status    interface{} // 状态
	CreateAt  *gtime.Time // 创建日期
	UpdateAt  *gtime.Time // 更新日期
	DeletedAt *gtime.Time // 删除日期
}
