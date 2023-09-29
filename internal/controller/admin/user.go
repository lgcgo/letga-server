package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cUser struct{}

var User = cUser{}

// 创建用户
func (c *cUser) Create(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	var (
		in  *model.UserCreateInput
		out *model.User
	)
	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	// 创建实体
	if out, err = service.User().CreateUser(ctx, in); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}

// 获取用户
func (c *cUser) Get(ctx context.Context, req *v1.UserGetReq) (res *v1.UserGetRes, err error) {
	var (
		out *model.User
	)
	// 获取数据
	if out, err = service.User().GetUser(ctx, req.Key); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}

// 修改用户
func (c *cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	var (
		in  *model.UserUpdateInput
		out *model.User
	)
	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	// 更新数据
	if out, err = service.User().UpdateUser(ctx, in); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}

// 删除用户
func (c *cUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	// 删除数据
	err = service.User().DeleteUsers(ctx, req.Keys)
	return
}

// 设置用户状态
func (c *cUser) SetStatus(ctx context.Context, req *v1.UserSetStatusReq) (res *v1.UserSetStatusRes, err error) {
	var (
		user *model.User
	)
	// 更新数据
	if user, err = service.User().SetUserStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	// 相应转换
	err = gconv.Struct(user, &res)
	return
}

// 获取用户分页
func (c *cUser) GetPage(ctx context.Context, req *v1.UserGetPageReq) (res *v1.UserGetPageRes, err error) {
	var (
		in  *model.UserPageInput
		out *model.UserPageOutput
	)
	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	// 获取数据
	if out, err = service.User().GetUserPage(ctx, in); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}
