package model

import (
	"letga/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

type Media struct {
	gmeta.Meta `orm:"table:media"`
	Key        string `json:"key"`     // 表索引
	UserKey    string `json:"userKey"` // 表索引
	entity.Media
}

// 用户搜索
type MediaSearch struct {
	gmeta.Meta `search:"name"`
	Key        string `f:"hashFc"`         // 索引
	UserKey    string `f:"hashFc:user,id"` // 用户索引
	Name       string `f:"like"`           // 文件名
	Hash       string `f:"like"`           // 哈希值
	FileType   string `f:"equal"`          // 文件类型
	Storage    string `f:"equal"`          // 储存库
	Status     string `f:"equal"`          // 状态
}
type MediaPageInput struct {
	PageParams
	MediaSearch
}
type MediaPageOutput struct {
	Data []*Media
	Page
}
