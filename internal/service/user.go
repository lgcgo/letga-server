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
	IUser interface {
		// 创建用户
		// - 服务层保证至少一种登录方式
		// - 账号|手机号|邮箱其中一个必填
		// - 当存在Account账号，则密码必填
		CreateUser(ctx context.Context, in *model.UserCreateInput) (*model.User, error)
		// 获取用户
		GetUser(ctx context.Context, key string) (*model.User, error)
		// 修改用户
		UpdateUser(ctx context.Context, in *model.UserUpdateInput) (*model.User, error)
		// 设置用户状态
		SetUserStatus(ctx context.Context, in *model.StatusSetInput) (*model.User, error)
		// 删除用户
		DeleteUsers(ctx context.Context, keys []string) error
		// 获取用户分页
		GetUserPage(ctx context.Context, in *model.UserPageInput) (*model.UserPageOutput, error)
		// 获取当前用户(用于前台)
		GetCurrentUser(ctx context.Context) (*model.User, error)
		// 修改用户账户(用于前台)
		UpdateCurrentAccount(ctx context.Context, account string) (*model.User, error)
		// 修改用户手机号(用于前台)
		UpdateCurrentMobile(ctx context.Context, mobile string) (*model.User, error)
		// 修改用户邮箱(用于前台)
		UpdateCurrentEmail(ctx context.Context, email string) (*model.User, error)
		// 修改用户密码(用于前台)
		UpdateCurrentPassword(ctx context.Context, password string) (*model.User, error)
		// 用户登录
		Signin(ctx context.Context, in *model.UserSigninInput) (*model.AuthToken, error)
		// 角色登录
		SigninDrect(ctx context.Context, user *model.User, role *model.AuthRole) (*model.AuthToken, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
