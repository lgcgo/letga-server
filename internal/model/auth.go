package model

import (
	"letga/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

// 授权Token
type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    uint   `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

type AuthAuthorizationInput struct {
	Subject string
	Role    string
}

type AuthVerifyInput struct {
	Subject string
	Path    string
	Method  string
	Role    string
}

/**
* 用户授权
 */
type AuthAccess struct {
	gmeta.Meta `orm:"table:user_access"`
	Key        string    `json:"key"`                        // 索引
	UserKey    string    `json:"userKey"`                    // 索引
	User       *User     `json:"user" orm:"with:id=user_id"` // 用户对象
	RoleKey    string    `json:"roleKey"`                    // 角色索引
	Role       *AuthRole `json:"role" orm:"with:id=role_id"` // 角色对象
	entity.AuthAccess
}

type AuthAccessSetupInput struct {
	UserKey        string
	AppendRoleKeys []string
	RemoveRoleKeys []string
}

type AuthAccessSearch struct {
	// gmeta.Meta `search:""`
	Key     string `f:"hashFc"`              // 索引
	UserKey string `f:"hashFc:user,id"`      // 用户索引
	RoleKey string `f:"hashFc:auth_role,id"` // 角色索引
	Status  string `f:"equal"`               // 状态
}

type AuthAccessPageInput struct {
	PageParams
	AuthAccessSearch
	CtxScene string
	Roks     []string
}
type AuthAccessPageOutput struct {
	Data []*AuthAccess
	Page
}

/**
* 角色管理
**/
type AuthRole struct {
	gmeta.Meta `orm:"table:auth_role"`
	Key        string `json:"key"`       // 索引
	ParentKey  string `json:"parentKey"` // 父索引
	entity.AuthRole
}

type AuthRoleCreateInput struct {
	ParentKey       string
	Title           string
	Name            string
	AppendRouteKeys []string
}

type AuthRoleUpdateInput struct {
	Key             string
	ParentKey       string
	Title           string
	Name            string
	AppendRouteKeys []string
	RemoveRouteKeys []string
}

type AuthRoleTreeInput struct {
	Keys         []string
	Search       string
	CtxScene     string
	CtxRoleKey   string
	CtxUserKey   string
	CtxFormName  string
	CtxFormField string
}

type AuthRoleFliterTreeInput struct {
	CtxScene string
}

/**
* 角色授权
 */
type AuthRoleAccess struct {
	gmeta.Meta `orm:"table:auth_role"`
	Key        string `json:"key"`      // 角色索引
	RoleKey    string `json:"roleKey"`  // 角色索引
	RouteKey   string `json:"routeKey"` // 权限索引
	entity.AuthRoleAccess
}

// 新增授权
type AuthRoleAccessAppendInput struct {
	RoleKey   string
	RouteKeys []string `json:"routeKeys"` // 权限索引
}
type AuthRoleAccessAppendOutput struct {
}

// 移除授权
type AuthRoleAccessRemoveInput struct {
	RoleKey   string
	RouteKeys []string `json:"routeKeys"` // 权限索引
}
type AuthRoleAccessRemoveOutput struct {
}

// 角色授权过滤字段
type AuthRoleAccessSearch struct {
	gmeta.Meta `search:"title,path"`
	Key        string `f:"hashFc"`               // Id
	RoleKey    string `f:"hashFc:auth_role,id"`  // 角色索引
	RouteKey   string `f:"hashFc:auth_route,id"` // 路由索引
	Status     string `f:"equal"`                // 状态
}

// 角色授权分页输入
type AuthRoleAccessPageInput struct {
	PageParams
	AuthRoleAccessSearch
	CtxScene   string
	CtxRoleKey string
	Apks       []string
	Rmks       []string
}

// 角色授权分页输出
type AuthRoleAccessPageOutput struct {
	Data []*AuthRoleAccess
	Page
	// 上下文索引集
	// CtxCurrentKeys []string
	// 新增索引集
	AppendKeys []string
	// 移除索引集
	RemoveKeys []string
}

/**
* 路由管理
**/
type AuthRoute struct {
	gmeta.Meta `orm:"table:auth_route"`
	Key        string `json:"key"`                        // 索引
	MenuKey    string `json:"menuKey"`                    // 菜单索引
	Menu       *Menu  `json:"menu" orm:"with:id=menu_id"` // 关联菜单
	entity.AuthRoute
}
type AuthRouteCreateInput struct {
	MenuKey string
	Title   string
	Path    string
	Method  string
	Remark  string
	Weight  int
}
type AuthRouteUpdateInput struct {
	Key     string
	MenuKey string
	Title   string
	Path    string
	Method  string
	Remark  string
	Weight  int
}

// 路由过滤字段
type AuthRouteSearch struct {
	gmeta.Meta `search:"title,path"`
	Key        string `f:"hashFc"`         // Id
	MenuKey    string `f:"hashFc:menu,id"` // 菜单索引
	Title      string `f:"like"`           // 标题
	Path       string `f:"like"`           // 路由地址
	Method     string `f:"equal"`          // 请求方法
	Status     string `f:"equal"`          // 状态
}

// 路由过滤索引
type AuthRoutePageInput struct {
	PageParams
	AuthRouteSearch
	CtxScene   string
	CtxRoleKey string

	Mnks []string `json:"mnks"` // 关联Menu索引
}

type AuthRoutePageOutput struct {
	Data []*AuthRoute
	Page
	// 上下文中来自子级的集合
	CtxDirectKeys []string
}

// 路由选择库
// type AuthRouteSelectLibraryPageInput struct {
// 	PageParams
// 	AuthRouteSearch
// 	CtxName string
// 	CtxKey  string
// }
// type AuthRouteSelectLibraryPageOutput struct {
// 	Data []*AuthRoute
// 	Page
// 	CtxCurrentKeys []string
// }
