// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure for table media.
type Media struct {
	Id        uint        `json:"id"        ` // ID
	UserId    uint        `json:"userId"    ` // 用户ID
	Name      string      `json:"name"      ` // 文件名
	Path      string      `json:"path"      ` // 路径
	Size      uint        `json:"size"      ` // 大小
	FileType  string      `json:"fileType"  ` // 文件类型
	MimeType  string      `json:"mimeType"  ` // MIME类型
	Hash      string      `json:"hash"      ` // 哈希值
	Extparam  string      `json:"extparam"  ` // 透传数据
	Storage   string      `json:"storage"   ` // 储存库
	Status    string      `json:"status"    ` // 状态
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建日期
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新日期
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除日期
}
