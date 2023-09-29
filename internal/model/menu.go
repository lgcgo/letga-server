package model

import (
	"letga/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

/**
* 菜单管理
**/
type Menu struct {
	gmeta.Meta `orm:"table:menu"`
	Key        string `json:"key"`       // 索引
	ParentKey  string `json:"parentKey"` // 父索引
	entity.Menu
}

type MenuCreateInput struct {
	ParentKey string
	Title     string
	Icon      string
	CoverUrl  string
	Remark    string
	Weight    int
}

type MenuUpdateInput struct {
	Key       string
	ParentKey string
	Title     string
	Icon      string
	CoverUrl  string
	Remark    string
	Weight    int
}

type MenuTreeInput struct {
	Keys        []string
	Search      string
	CtxScene    string
	CtxMenuKey  string
	CtxRouteKey string
}

type MenuFliterTreeInput struct {
	CtxScene string
}
