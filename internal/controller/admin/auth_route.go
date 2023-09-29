package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAuthRoute struct{}

var AuthRoute = cAuthRoute{}

// 添加路由
func (c *cAuthRoute) Create(ctx context.Context, req *v1.AuthRouteCreateReq) (res *v1.AuthRouteCreateRes, err error) {
	var (
		in  *model.AuthRouteCreateInput
		out *model.AuthRoute
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().CreateRoute(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 获取权限路由
func (c *cAuthRoute) Get(ctx context.Context, req *v1.AuthRouteGetReq) (res *v1.AuthRouteGetRes, err error) {
	var (
		out *model.AuthRoute
	)
	if out, err = service.Auth().GetRoute(ctx, req.Key); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 修改权限路由
func (c *cAuthRoute) Update(ctx context.Context, req *v1.AuthRouteUpdateReq) (res *v1.AuthRouteUpdateRes, err error) {
	var (
		in  *model.AuthRouteUpdateInput
		out *model.AuthRoute
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().UpdateRoute(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 设置路由状态
func (c *cAuthRoute) SetStatus(ctx context.Context, req *v1.AuthRouteSetStatusReq) (res *v1.AuthRouteSetStatusRes, err error) {
	var (
		out *model.AuthRoute
	)
	if out, err = service.Auth().SetRouteStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 删除权限路由
func (c *cAuthRoute) Delete(ctx context.Context, req *v1.AuthRouteDeleteReq) (res *v1.AuthRouteDeleteRes, err error) {
	err = service.Auth().DeleteRoutes(ctx, req.Keys)
	return
}

// 获取路由分页
func (c *cAuthRoute) GetPage(ctx context.Context, req *v1.AuthRouteGetPageReq) (res *v1.AuthRouteGetPageRes, err error) {
	var (
		in  *model.AuthRoutePageInput
		out *model.AuthRoutePageOutput
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().GetRoutePage(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}
