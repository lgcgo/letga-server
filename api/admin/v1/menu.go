package v1

import (
	"letga/api/admin/v1/model"

	"github.com/gogf/gf/v2/frame/g"
)

/**
*  菜单管理
**/

// 获取菜单
type MenuGetReq struct {
	g.Meta `path:"/menu" tags:"Admin@MenuService" method:"get" summary:"Get menu"`
	Key    string `json:"key" v:"required"`
}
type MenuGetRes struct {
	model.Menu
}

// 创建菜单
type MenuCreateReq struct {
	g.Meta    `path:"/menu" tags:"Admin@MenuService" method:"post" summary:"Create menu"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title" v:"required|max-length:12"`
	Icon      string `josn:"icon"`
	CoverUrl  string `json:"coverUrl" v:"max-length:256"`
	Remark    string `json:"remark" v:"max-length:256"`
	Weigh     uint   `json:"weigh" v:"between:0,9999"`
}
type MenuCreateRes struct {
	model.Menu
}

// 更新菜单
type MenuUpdateReq struct {
	g.Meta    `path:"/menu" tags:"Admin@MenuService" method:"put" summary:"Update menu"`
	Key       string `json:"key" v:"required"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title" v:"required|max-length:12"`
	Icon      string `josn:"icon"`
	CoverUrl  string `json:"coverUrl" v:"max-length:256"`
	Remark    string `json:"remark" v:"max-length:255"`
	Weigh     uint   `json:"weigh" v:"between:0,9999"`
}
type MenuUpdateRes struct {
	model.Menu
}

// 设置菜单状态
type MenuSetStatusReq struct {
	g.Meta `path:"/menu/status" method:"put" tags:"Admin@MenuService" summary:"Set menu status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type MenuSetStatusRes struct {
	model.Menu
}

// 删除菜单
type MenuDeleteReq struct {
	g.Meta `path:"/menu" tags:"Admin@MenuService" method:"delete" summary:"Delete menu"`
	Keys   []string `json:"keys" v:"required"`
}
type MenuDeleteRes struct {
}

// 获取菜单树
type MenuGetTreeReq struct {
	g.Meta `path:"/menu/tree" tags:"Admin@MenuService" method:"get" summary:"Get menu tree"`
	Keys   []string `json:"keys"`
	Search string   `json:"search"`

	// 业务上下文
	CtxScene    string `json:"ctxScene" v:"in:mainTable,menuForm"`
	CtxMenuKey  string `json:"ctxMenuKey"`
	CtxRouteKey string `json:"ctxRouteKey"`
}
type MenuGetTreeRes struct {
	Data            []*model.MenuTreeData `json:"data"`
	CtxSelectedKeys []string              `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string              `json:"ctxDisabledKeys"`
}

// 菜单过滤树
type MenuGetFliterTreeReq struct {
	g.Meta `path:"/menu/flitertree" method:"get" tags:"Admin@MenuService" summary:"Get menu flitertree"`
	// 上下文场景
	CtxScene string `json:"ctxScene" v:"in:routeSidebar"`
}
type MenuGetFliterTreeRes struct {
	Data []*model.FliterTree `json:"data"`
}
