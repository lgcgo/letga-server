package model

// 权限菜单返回数据
type Menu struct {
	Key       string `json:"key"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	CoverUrl  string `json:"coverUrl"`
	Remark    string `json:"remark"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
	Status    string `json:"status"`
}
type MenuTreeData struct {
	Key       string          `json:"key"`
	ParentKey string          `json:"parentKey"`
	Title     string          `json:"title"`
	Weight    int             `json:"weight"`
	Source    *Menu           `json:"source"`
	Children  []*MenuTreeData `json:"children"`
}
