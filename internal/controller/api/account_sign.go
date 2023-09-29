package api

import (
	"context"
	v1 "letga/api/api/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAccountSign struct{}

var AccountSign = &cAccountSign{}

// 用户注册并授权
func (c *cAccountSign) SignUp(ctx context.Context, req *v1.AccountSignUpReq) (res *v1.AccountSignUpRes, err error) {
	var (
		ser  = service.User()
		in   *model.UserCreateInput
		out  *model.AuthToken
		user *model.User
		role *model.AuthRole
	)
	// 校验验证码
	// if err = service.Sms().Verify(ctx, req.Captcha, "register"); err != nil {
	// 	return
	// }
	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	// 创建实体
	if user, err = ser.CreateUser(ctx, in); err != nil {
		return
	}
	// 获取默认角色名称
	if role, err = service.Auth().GetDefaultRole(ctx); err != nil {
		return
	}
	// 角色登录
	if out, err = ser.SigninDrect(ctx, user, role); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}

// 账户|手机号|邮箱 + 密码 登录
func (c *cAccountSign) Signin(ctx context.Context, req *v1.AccountSigninReq) (res *v1.AccountSigninRes, err error) {
	var (
		in  *model.UserSigninInput
		out *model.AuthToken
	)
	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	// 账号登录
	if out, err = service.User().Signin(ctx, in); err != nil {
		return
	}
	// 转换响应
	err = gconv.Struct(out, &res)
	return
}

// 账户登出
func (c *cAccount) Signout(ctx context.Context, req *v1.AccountSignoutReq) (res *v1.AccountSignoutRes, err error) {
	// 这里可以添加一些Token拉黑操作
	return
}
