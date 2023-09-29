package model

import (
	"letga/utility/tree"
)

// 分页参数
type PageParams struct {
	Current  int
	PageSize int
	Sort     string
	Search   string
	Apks     []string // appendKeys，指定追加的项
	Rmks     []string // removeKeys，指定排除的项
}
type Page struct {
	Total           int
	Apks            []string `json:"apks"`            // appendKeys，指定追加的项
	Rmks            []string `json:"rmks"`            // removeKeys，指定排除的项
	CtxSelectedKeys []string `json:"ctxSelectedKeys"` // 上下文中的选中项
	CtxDisabledKeys []string `json:"ctxDisabledKeys"` // 上下文中的禁用项
}

// 树状数据
type TreeOutput struct {
	Data            []*tree.TreeNode `json:"data"`
	CtxSelectedKeys []string         `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string         `json:"ctxDisabledKeys"`
}

// 统计树列表
type FliterTreeData struct {
	Key       string `json:"key"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title"`
	Count     int    `json:"count"`
	Weight    int    `json:"weight"`
}
type FliterTreeItem struct {
	Name  string           `json:"name"`
	Data  []*tree.TreeNode `json:"data"`
	Total int              `json:"total"`
}
type FliterTree struct {
	Data            []*FliterTreeItem
	CtxSelectedKeys []string `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string `json:"ctxDisabledKeys"`
}

// 状态设置
type StatusSetInput struct {
	Key   string
	Value string
}
