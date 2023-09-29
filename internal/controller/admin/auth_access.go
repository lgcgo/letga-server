package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAuthAccess struct{}

var AuthAccess = cAuthAccess{}

// 设置授权
func (c *cAuthAccess) Setup(ctx context.Context, req *v1.AuthAccessSetupReq) (res *v1.AuthAccessSetupRes, err error) {
	err = service.Auth().SetupAccess(ctx, &model.AuthAccessSetupInput{
		UserKey:        req.UserKey,
		AppendRoleKeys: req.AppendRoleKeys,
		RemoveRoleKeys: req.RemoveRoleKeys,
	})
	return
}

// 设置授权状态
func (c *cAuthAccess) SetStatus(ctx context.Context, req *v1.AuthAccessSetStatusReq) (res *v1.AuthAccessSetStatusRes, err error) {
	var (
		out *model.AuthAccess
	)
	if out, err = service.Auth().SetAccessStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 获取分页
func (c *cAuthAccess) GetPage(ctx context.Context, req *v1.AuthAccessGetPageReq) (res *v1.AuthAccessGetPageRes, err error) {
	var (
		in  *model.AuthAccessPageInput
		out *model.AuthAccessPageOutput
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Auth().GetAccessPage(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 删除授权
func (c *cAuthAccess) Delete(ctx context.Context, req *v1.AuthAccessDeleteReq) (res *v1.AuthAccessDeleteRes, err error) {
	// 删除数据
	err = service.Auth().DeleteAccesses(ctx, req.Keys)
	return
}
