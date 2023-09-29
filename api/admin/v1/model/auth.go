package model

// 签发授权
type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

// 角色用户
type AuthAccess struct {
	Key      string    `json:"key"`      // 索引
	UserKey  string    `json:"userKey"`  // 用户索引
	User     *User     `json:"user"`     // 用户模型
	RoleKey  string    `json:"roleKey"`  // 角色模型
	Role     *AuthRole `json:"role"`     // 角色对象
	Status   string    `json:"status"`   // 状态
	CreateAt string    `json:"createAt"` // 创建日期
}

// 权限角色
type AuthRole struct {
	Key       string `json:"key"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
	Status    string `json:"status"`
}
type AuthRoleTreeData struct {
	Key       string              `json:"key"`
	ParentKey string              `json:"parentKey"`
	Title     string              `json:"title"`
	Weight    int                 `json:"weight"`
	Source    *AuthRole           `json:"source"`
	Children  []*AuthRoleTreeData `json:"children"`
}

// 角色授权
type AuthRoleAccess struct {
	RoleKey  string
	RouteKey string
}

// 权限路由返回数据
type AuthRoute struct {
	Key      string `json:"key"`      // 索引
	MenuKey  string `json:"menuKey"`  // 菜单索引
	Menu     *Menu  `json:"menu"`     // 关联菜单
	Title    string `json:"title"`    // 标题
	Path     string `json:"path"`     // 路由地址
	Method   string `json:"method"`   // 请求方法
	Remark   string `json:"remark"`   // 备注
	Status   string `json:"status"`   // 状态
	Weight   int    `json:"weight"`   // 权重
	CreateAt string `json:"createAt"` // 创建日期
	UpdateAt string `json:"updateAt"` // 更新日期
}
