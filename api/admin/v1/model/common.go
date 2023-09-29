package model

import "letga/utility/tree"

// 分页
type PageParams struct {
	Current  int    `json:"current"`
	PageSize int    `json:"pageSize"`
	Sort     string `json:"sort"`
	Search   string `json:"search"`
	// 指定的追加项
	Apks []string `json:"apks,omitempty"`
	// 指定的排除项
	Rmks []string `json:"rmks,omitempty"`
}
type Page struct {
	Total           int      `json:"total"`
	CtxSelectedKeys []string `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string `json:"ctxDisabledKeys"`
}

// 过滤树
type FliterTreeData struct {
	Key       string `json:"key"`
	ParentKey string `json:"parentKey"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
}
type FliterTree struct {
	Name            string           `json:"name"`
	Data            []*tree.TreeNode `json:"data"`
	Total           int              `json:"total"`
	CtxSelectedKeys []string         `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string         `json:"ctxDisabledKeys"`
}
