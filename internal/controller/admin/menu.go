package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cMenu struct{}

var Menu = cMenu{}

// 添加菜单
func (c *cMenu) Create(ctx context.Context, req *v1.MenuCreateReq) (res *v1.MenuCreateRes, err error) {
	var (
		in  *model.MenuCreateInput
		out *model.Menu
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Menu().CreateMenu(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 获取菜单
func (c *cMenu) Get(ctx context.Context, req *v1.MenuGetReq) (res *v1.MenuGetRes, err error) {
	var (
		out *model.Menu
	)
	if out, err = service.Menu().GetMenu(ctx, req.Key); err != nil {
		return nil, err
	}
	err = gconv.Struct(out, &res)
	return
}

// 修改菜单
func (c *cMenu) Update(ctx context.Context, req *v1.MenuUpdateReq) (res *v1.MenuUpdateRes, err error) {
	var (
		in  *model.MenuUpdateInput
		out *model.Menu
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Menu().UpdateMenu(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 设置角色状态
func (c *cMenu) SetStatus(ctx context.Context, req *v1.MenuSetStatusReq) (res *v1.MenuSetStatusRes, err error) {
	var (
		out *model.Menu
	)
	if out, err = service.Menu().SetRoleStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 删除菜单
func (c *cMenu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	err = service.Menu().DeleteMenus(ctx, req.Keys)
	return
}

// 获取菜单树
func (c *cMenu) GetTree(ctx context.Context, req *v1.MenuGetTreeReq) (res *v1.MenuGetTreeRes, err error) {
	var (
		out *model.TreeOutput
	)
	if out, err = service.Menu().GetMenuTree(ctx, &model.MenuTreeInput{
		Keys:        req.Keys,
		Search:      req.Search,
		CtxScene:    req.CtxScene,
		CtxRouteKey: req.CtxRouteKey,
		CtxMenuKey:  req.CtxMenuKey,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 路由统计
func (c *cAuthRoute) GetFliterTree(ctx context.Context, req *v1.MenuGetFliterTreeReq) (res *v1.MenuGetFliterTreeRes, err error) {
	var (
		out *model.FliterTree
	)
	if out, err = service.Menu().GetFliterTree(ctx, &model.MenuFliterTreeInput{
		CtxScene: req.CtxScene,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}
