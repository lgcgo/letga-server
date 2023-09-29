package v1

import (
	"letga/api/admin/v1/model"

	"github.com/gogf/gf/v2/frame/g"
)

/**
* 权限授权
**/
type AuthAccessSetupReq struct {
	g.Meta         `path:"/auth/access/setup" tags:"Admin@AuthService" method:"post" summary:"Setup access"`
	UserKey        string   `json:"userKey" v:"required"`
	AppendRoleKeys []string `json:"appendRoleKeys" v:"required-without:removeRoleKeys"`
	RemoveRoleKeys []string `json:"removeRoleKeys" v:"required-without:appendRoleKeys"`
}
type AuthAccessSetupRes struct {
}

// 设置授权状态
type AuthAccessSetStatusReq struct {
	g.Meta `path:"/auth/access/status" method:"put" tags:"Admin@AuthService" summary:"Set access status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type AuthAccessSetStatusRes struct {
	model.User
}

// 删除授权
type AuthAccessDeleteReq struct {
	g.Meta `path:"/auth/access" method:"delete" tags:"Admin@AuthService" summary:"Delete access"`
	Keys   []string `json:"keys" v:"required"`
}
type AuthAccessDeleteRes struct{}

// 授权列表
type AuthAccessGetPageReq struct {
	g.Meta `path:"/auth/accesses" method:"get" tags:"Admin@AuthService" summary:"Get access page"`
	model.PageParams
	Key     string `json:"key"`     // 索引
	UserKey string `json:"userKey"` // 用户索引
	RoleKey string `json:"roleKey"` // 角色索引

	// 上下文场景
	CtxScene string `json:"ctxScene" v:"in:mainTable"`

	// 关联Role索引集
	Roks []string `json:"roks,omitempty"`
}
type AuthAccessGetPageRes struct {
	Data []*model.AuthAccess `json:"data,omitempty"`
	model.Page
}

/**
*  权限角色
**/

// 获取角色
type AuthRoleGetReq struct {
	g.Meta `path:"/auth/role" tags:"Admin@AuthService" method:"get" summary:"Get role"`
	Key    string `json:"key" v:"required"`
}
type AuthRoleGetRes struct {
	model.AuthRole
}

// 创建角色
type AuthRoleCreateReq struct {
	g.Meta    `path:"/auth/role" tags:"Admin@AuthService" method:"post" summary:"Create role"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title" v:"required|length:2,12"`
	Name      string `json:"name" v:"required|length:2,32"`
	// 角色授权
	AppendAccess []model.AuthRoleAccess `json:"appendAccess"`
}
type AuthRoleCreateRes struct {
	model.AuthRole
}

// 更新角色
type AuthRoleUpdateReq struct {
	g.Meta    `path:"/auth/role" tags:"Admin@AuthService" method:"put" summary:"Update role"`
	Key       string `json:"key" v:"required"`
	ParentKey string `json:"parentKey"`
	Title     string `json:"title" v:"required|length:2,32"`
	Name      string `json:"name" v:"required|length:2,32"`
	// 新增的授权路由
	AppendRouteKeys []string `json:"appendRouteKeys"`
	// 移除的授权路由
	RemoveRouteKeys []string `json:"removeRouteKeys"`
}
type AuthRoleUpdateRes struct {
	model.AuthRole
}

// 设置角色状态
type AuthRoleSetStatusReq struct {
	g.Meta `path:"/auth/role/status" method:"put" tags:"Admin@AuthService" summary:"Set role status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type AuthRoleSetStatusRes struct {
	model.AuthRole
}

// 删除角色
type AuthRoleDeleteReq struct {
	g.Meta `path:"/auth/role" tags:"Admin@AuthService" method:"delete" summary:"Delete role"`
	Keys   []string `json:"keys" v:"required"`
}
type AuthRoleDeleteRes struct {
}

// 获取角色树
type AuthRoleGetTreeReq struct {
	g.Meta `path:"/auth/role/tree" tags:"Admin@AuthService" method:"get" summary:"Get role tree"`
	Keys   []string `json:"keys"`
	Search string   `json:"search"`

	// 业务上下文
	CtxScene     string `json:"ctxScene" v:"in:mainTable,roleAccessTransfer,roleForm"`
	CtxRoleKey   string `json:"ctxRoleKey"`
	CtxUserKey   string `json:"ctxUserKey"`
	CtxFormName  string `json:"ctxFormName"`
	CtxFormField string `json:"ctxFormField"`
}
type AuthRoleGetTreeRes struct {
	Data            []*model.AuthRoleTreeData `json:"data"`
	CtxSelectedKeys []string                  `json:"ctxSelectedKeys"`
	CtxDisabledKeys []string                  `json:"ctxDisabledKeys"`
}

// 角色过滤树
type AuthRoleGetFliterTreeReq struct {
	g.Meta `path:"/auth/role/flitertree" method:"get" tags:"Admin@AuthService" summary:"Get role flitertree"`
	// 上下文场景
	CtxScene string `json:"ctxScene" v:"in:accessSidebar"`
}
type AuthRoleGetFliterTreeRes struct {
	Data []*model.FliterTree `json:"data"`
}

/**
*  权限路由
**/

// 获取权限路由
type AuthRouteGetReq struct {
	g.Meta `path:"/auth/route" tags:"Admin@AuthService" method:"get" summary:"Get route"`
	Key    string `json:"key" v:"required"`
}
type AuthRouteGetRes struct {
	model.AuthRoute
}

// 创建权限路由
type AuthRouteCreateReq struct {
	g.Meta  `path:"/auth/route" tags:"Admin@AuthService" method:"post" summary:"Create route"`
	MenuKey string `json:"menuKey" v:"max-length:12"`
	Title   string `json:"title" v:"required|length:4,12"`
	Path    string `json:"path" v:"required|length:2,64"`
	Method  string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Remark  string `json:"remark" v:"max-length:256"`
	Weigh   uint   `json:"weigh" v:"integer|between:0,9999"`
}
type AuthRouteCreateRes struct {
	model.AuthRoute
}

// 更新权限路由
type AuthRouteUpdateReq struct {
	g.Meta  `path:"/auth/route" tags:"Admin@AuthService" method:"put" summary:"Update route"`
	Key     string `json:"key" v:"required"`
	MenuKey string `json:"menuKey" v:"max-length:12"`
	Title   string `json:"title" v:"required|length:4,12"`
	Path    string `json:"path" v:"required|length:2,64"`
	Method  string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Remark  string `json:"remark" v:"max-length:256"`
	Weigh   uint   `json:"weigh" v:"integer|between:0,9999"`
}
type AuthRouteUpdateRes struct {
	model.AuthRoute
}

// 设置路由状态
type AuthRouteSetStatusReq struct {
	g.Meta `path:"/auth/route/status" method:"put" tags:"Admin@AuthService" summary:"Set route status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type AuthRouteSetStatusRes struct {
	model.AuthRoute
}

// 删除权限路由
type AuthRouteDeleteReq struct {
	g.Meta `path:"/auth/route" tags:"Admin@AuthService" method:"delete" summary:"Delete route"`
	Keys   []string `json:"keys" v:"required"`
}
type AuthRouteDeleteRes struct {
}

// 路由分页
type AuthRouteGetPageReq struct {
	g.Meta `path:"/auth/routes" method:"get" tags:"Admin@AuthService" summary:"Get route page"`
	model.PageParams
	Key     string `json:"key"`                             // 索引
	MenuKey string `json:"menuKey"`                         // 菜单索引
	Title   string `json:"title"`                           // 标题
	Path    string `json:"path" v:"regex:^/[A-Za-z0-9/]+$"` // 路由地址
	Method  string `json:"method"`                          // 请求方法
	Remark  string `json:"remark"`                          // 备注
	Status  string `json:"status"`                          // 状态

	// 上下文场景
	CtxScene   string `json:"ctxScene" v:"in:mainTable,roleAccessTable,roleAccessModel"`
	CtxRoleKey string `json:"ctxRoleKey"`

	// 关联Menu索引集
	Mnks []string `json:"mnks,omitempty"`
}
type AuthRouteGetPageRes struct {
	Data []*model.AuthRoute `json:"data,omitempty"`
	model.Page
	CtxDirectKeys []string `json:"ctxDirectKeys"`
}
