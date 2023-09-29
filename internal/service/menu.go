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
	IMenu interface {
		// 添加菜单
		CreateMenu(ctx context.Context, in *model.MenuCreateInput) (*model.Menu, error)
		// 获取菜单
		GetMenu(ctx context.Context, key string) (*model.Menu, error)
		// 修改菜单
		UpdateMenu(ctx context.Context, in *model.MenuUpdateInput) (*model.Menu, error)
		// 更改菜单状态
		SetRoleStatus(ctx context.Context, in *model.StatusSetInput) (*model.Menu, error)
		// 删除菜单
		DeleteMenus(ctx context.Context, keys []string) error
		// 获取菜单树
		GetMenuTree(ctx context.Context, in *model.MenuTreeInput) (*model.TreeOutput, error)
		// 路由统计
		GetFliterTree(ctx context.Context, in *model.MenuFliterTreeInput) (*model.FliterTree, error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
