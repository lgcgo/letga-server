package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAuthRole struct{}

var AuthRole = cAuthRole{}

// 添加角色
func (c *cAuthRole) Create(ctx context.Context, req *v1.AuthRoleCreateReq) (res *v1.AuthRoleCreateRes, err error) {
	var (
		in  *model.AuthRoleCreateInput
		out *model.AuthRole
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().CreateRole(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 获取角色
func (c *cAuthRole) Get(ctx context.Context, req *v1.AuthRoleGetReq) (res *v1.AuthRoleGetRes, err error) {
	var (
		out *model.AuthRole
	)
	if out, err = service.Auth().GetRole(ctx, req.Key); err != nil {
		return nil, err
	}
	err = gconv.Struct(out, &res)
	return
}

// 修改角色
func (c *cAuthRole) Update(ctx context.Context, req *v1.AuthRoleUpdateReq) (res *v1.AuthRoleUpdateRes, err error) {
	var (
		in  *model.AuthRoleUpdateInput
		out *model.AuthRole
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().UpdateRole(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 设置角色状态
func (c *cAuthRole) SetStatus(ctx context.Context, req *v1.AuthRoleSetStatusReq) (res *v1.AuthRoleSetStatusRes, err error) {
	var (
		out *model.AuthRole
	)
	if out, err = service.Auth().SetRoleStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 删除角色
func (c *cAuthRole) Delete(ctx context.Context, req *v1.AuthRoleDeleteReq) (res *v1.AuthRoleDeleteRes, err error) {
	err = service.Auth().DeleteRoles(ctx, req.Keys)
	return
}

// 获取角色树
func (c *cAuthRole) Tree(ctx context.Context, req *v1.AuthRoleGetTreeReq) (res *v1.AuthRoleGetTreeRes, err error) {
	var (
		out *model.TreeOutput
	)
	if out, err = service.Auth().GetRoleTree(ctx, &model.AuthRoleTreeInput{
		Keys:       req.Keys,
		Search:     req.Search,
		CtxScene:   req.CtxScene,
		CtxRoleKey: req.CtxRoleKey,
		CtxUserKey: req.CtxUserKey,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 路由统计
func (c *cAuthRole) GetFliterTree(ctx context.Context, req *v1.AuthRoleGetFliterTreeReq) (res *v1.AuthRoleGetFliterTreeRes, err error) {
	var (
		out *model.FliterTree
	)
	if out, err = service.Auth().GetRoleFliterTree(ctx, &model.AuthRoleFliterTreeInput{
		CtxScene: req.CtxScene,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}
