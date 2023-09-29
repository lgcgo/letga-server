// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"letga/internal/model"
)

type (
	IAuth interface {
		// 批量创建用户权限
		SetupAccess(ctx context.Context, in *model.AuthAccessSetupInput) error
		// 更改用户状态
		SetAccessStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthAccess, error)
		// 获取用户授权
		GetAccess(ctx context.Context, key string) (*model.AuthAccess, error)
		// 获取授权分页
		GetAccessPage(ctx context.Context, in *model.AuthAccessPageInput) (*model.AuthAccessPageOutput, error)
		// 删除用户授权
		DeleteAccesses(ctx context.Context, keys []string) error
		// 使用角色ID集删除
		DeleteAccessByRoleIds(ctx context.Context, roleIds []uint) error
		// 授权, 支持Bearer签发方案(Header Authorization: Bearer <token>)
		Authorization(ctx context.Context, in *model.AuthAuthorizationInput) (*model.AuthToken, error)
		// 刷新授权
		RefreshAuthorization(ctx context.Context, tokenString string) (*model.AuthToken, error)
		// 验证Token
		VerifyToken(ctx context.Context, tokenString string) (map[string]interface{}, error)
		// 验证路由
		Verify(ctx context.Context, in *model.AuthVerifyInput) error
		// 更新授权政策
		RefreshPolicy(ctx context.Context) error
		// 获取角色
		GetRole(ctx context.Context, key string) (*model.AuthRole, error)
		// 获取角色
		GetDefaultRole(ctx context.Context) (*model.AuthRole, error)
		// 创建角色
		CreateRole(ctx context.Context, in *model.AuthRoleCreateInput) (*model.AuthRole, error)
		// 修改角色
		UpdateRole(ctx context.Context, in *model.AuthRoleUpdateInput) (*model.AuthRole, error)
		// 更改角色状态
		SetRoleStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthRole, error)
		// 删除角色
		DeleteRoles(ctx context.Context, keys []string) error
		// 获取角色树
		GetRoleTree(ctx context.Context, in *model.AuthRoleTreeInput) (*model.TreeOutput, error)
		// 路由统计
		GetRoleFliterTree(ctx context.Context, in *model.AuthRoleFliterTreeInput) (*model.FliterTree, error)
		// 设置角色授权
		AppendRoleAccess(ctx context.Context, in *model.AuthRoleAccessAppendInput) (*model.AuthRoleAccessAppendOutput, error)
		// 移除角色授权
		RemoveRoleAccess(ctx context.Context, in *model.AuthRoleAccessRemoveInput) (*model.AuthRoleAccessRemoveOutput, error)
		// 创建路由
		CreateRoute(ctx context.Context, in *model.AuthRouteCreateInput) (*model.AuthRoute, error)
		// 获取路由
		GetRoute(ctx context.Context, key string) (*model.AuthRoute, error)
		// 修改路由
		UpdateRoute(ctx context.Context, in *model.AuthRouteUpdateInput) (*model.AuthRoute, error)
		// 更改域状态
		SetRouteStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthRoute, error)
		// 删除路由
		DeleteRoutes(ctx context.Context, keys []string) error
		// 获取路由分页
		GetRoutePage(ctx context.Context, in *model.AuthRoutePageInput) (*model.AuthRoutePageOutput, error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
