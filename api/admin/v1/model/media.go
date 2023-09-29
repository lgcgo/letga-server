package model

import "github.com/gogf/gf/v2/os/gtime"

type Media struct {
	Key      string      `json:"key"`      // 表索引
	UserKey  string      `json:"userKey"`  // 用户ID
	Name     string      `json:"name"`     // 文件名
	Path     string      `json:"path"`     // 路径
	Size     uint        `json:"size"`     // 大小
	Hash     string      `json:"hash"`     // 哈希值
	FileType string      `json:"fileType"` // 文件类型
	MimeType string      `json:"mimeType"` // MIME类型
	Extparam string      `json:"extparam"` // 透传数据
	Storage  string      `json:"storage"`  // 储存库
	Status   string      `json:"status"`   // 状态
	CreateAt *gtime.Time `json:"createAt"` // 创建日期
	UpdateAt *gtime.Time `json:"updateAt"` // 更新日期
}
