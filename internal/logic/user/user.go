// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package user

import (
	"context"
	"letga/internal/biz"
	"letga/internal/consts"
	"letga/internal/dao"
	"letga/internal/model"
	"letga/internal/model/do"
	"letga/internal/service"
	"letga/utility/hashid"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 创建用户
// - 服务层保证至少一种登录方式
// - 账号|手机号|邮箱其中一个必填
// - 当存在Account账号，则密码必填
func (s *sUser) CreateUser(ctx context.Context, in *model.UserCreateInput) (*model.User, error) {
	var (
		data      *do.User
		available bool
		err       error
	)
	// 使 账号|手机号|邮箱 其中一个必填
	if len(in.Account) == 0 && len(in.Email) == 0 && len(in.Mobile) == 0 {
		return nil, gerror.NewCode(biz.UserPassportInvalid)
	}
	// 账号防重，如果有
	if len(in.Account) > 0 {
		if available, err = dao.User.IsAccountAvailable(ctx, in.Account); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.WithCodef(biz.UserAccountExists, in.Account))
		}
		if len(in.Password) == 0 {
			return nil, gerror.NewCode(biz.UserPasswordEmpty)
		}
	} else {
		// 随机8位英文，重复由数据库抛出异常
		in.Account = grand.Letters(8)
	}
	// 手机号防重，如果有
	if len(in.Mobile) > 0 {
		if available, err = dao.User.IsMobileAvailable(ctx, in.Mobile); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.UserMobileExists)
		}
	}
	// Email防重，如果有
	if len(in.Email) > 0 {
		if available, err = dao.User.IsEmailAvailable(ctx, in.Email); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.UserEmailExists)
		}
	}
	// 支持无密码创建
	if len(in.Password) == 0 {
		in.Password = grand.Letters(6)
	}
	salt := grand.Letters(4)
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	data.Uuid = guid.S()
	data.Salt = salt
	data.Password = s.mustEncryptPasword(in.Password, salt)
	return dao.User.Insert(ctx, data)
}

// 获取用户
func (s *sUser) GetUser(ctx context.Context, key string) (*model.User, error) {
	var (
		user  *model.User
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_USER_SALT); err != nil {
		return nil, err
	}
	if user, err = dao.User.Get(ctx, keyId); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(biz.UserNotExists)
	}
	// if user.Status != "normal" {
	// 	return nil, gerror.NewCode(biz.UserStatusInvalid)
	// }
	return dao.User.Get(ctx, keyId)
}

// 修改用户
func (s *sUser) UpdateUser(ctx context.Context, in *model.UserUpdateInput) (*model.User, error) {
	var (
		data      *do.User
		user      *model.User
		err       error
		available bool
	)
	// 扫描数据
	if user, err = s.GetUser(ctx, in.Key); err != nil {
		return nil, err
	}
	// 特殊过滤 (超级管理员只能自己修改)
	if user.Id == uint(consts.RootAdminId) {
		if biz.Ctx().Get(ctx).User.Uuid != user.Uuid {
			return nil, gerror.NewCode(biz.AuthNotPermission)
		}
	}
	// 账户防重
	if len(in.Account) > 0 {
		if available, err = dao.User.IsAccountAvailable(ctx, in.Account, []uint{user.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.WithCodef(biz.UserAccountExists, in.Account))
		}
	}
	// 手机号防重
	if len(in.Mobile) > 0 {
		if available, err = dao.User.IsMobileAvailable(ctx, in.Mobile, []uint{user.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.UserMobileExists)
		}
	}
	// 邮箱防重
	if len(in.Email) > 0 {
		if available, err = dao.User.IsEmailAvailable(ctx, in.Email, []uint{user.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.NewCode(biz.UserEmailExists)
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 密码为空时不更新
	if len(in.Password) > 0 {
		salt := grand.Letters(4)
		data.Salt = salt
		data.Password = s.mustEncryptPasword(in.Password, salt)
	} else {
		data.Password = nil
	}
	data.Id = user.Id
	return dao.User.Update(ctx, data)
}

// 设置用户状态
func (s *sUser) SetUserStatus(ctx context.Context, in *model.StatusSetInput) (*model.User, error) {
	var (
		user *model.User
		err  error
	)

	if user, err = s.GetUser(ctx, in.Key); err != nil {
		return nil, err
	}
	// 限制禁止超级管理员
	if user.Id == uint(consts.RootAdminId) {
		return nil, gerror.NewCode(biz.AuthNotPermission)
	}
	return dao.User.SetStatus(ctx, user.Id, in.Value)
}

// 删除用户
func (s *sUser) DeleteUsers(ctx context.Context, keys []string) error {
	var (
		keyIds []uint
		err    error
	)
	// Key集解密
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_USER_SALT); err != nil {
		return err
	}
	// 限制删除超级管理员
	for _, v := range keyIds {
		if v == uint(consts.RootAdminId) {
			return gerror.NewCode(biz.AuthNotPermission)
		}
	}
	// 检测ID集
	if err = dao.User.CheckIds(ctx, keyIds); err != nil {
		return err
	}
	// 嵌套事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除用户
		if err := tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.User.Delete(ctx, keyIds)
		}); err != nil {
			return err
		}
		// 删除授权
		if err := tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthAccess.DeleteByUserIds(ctx, keyIds)
		}); err != nil {
			return err
		}
		// 更新授权
		if err := tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return service.Auth().RefreshPolicy(ctx)
		}); err != nil {
			return err
		}
		return nil
	})
}

// 获取用户分页
func (s *sUser) GetUserPage(ctx context.Context, in *model.UserPageInput) (*model.UserPageOutput, error) {
	var (
		data            []*model.User
		count           int
		ctxDisabledKeys []string
		err             error
	)
	if data, count, err = dao.User.PageData(ctx, &in.PageParams, &in.UserSearch); err != nil {
		return nil, err
	}
	// 根据上下文组装禁用项
	for _, v := range data {
		switch in.CtxScene {
		case "mainTable":
		case "accessSetupTable":
			if v.Id == consts.RootAdminId {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		}
	}
	return &model.UserPageOutput{
		Data: data,
		Page: model.Page{
			Total:           count,
			CtxDisabledKeys: ctxDisabledKeys,
		},
	}, nil
}

// 获取当前用户(用于前台)
func (s *sUser) GetCurrentUser(ctx context.Context) (*model.User, error) {
	return dao.User.GetByUuid(ctx, biz.Ctx().Get(ctx).User.Uuid)
}

// 修改用户账户(用于前台)
func (s *sUser) UpdateCurrentAccount(ctx context.Context, account string) (*model.User, error) {
	var (
		user      *model.User
		err       error
		available bool
	)
	// 扫描数据
	if user, err = s.GetCurrentUser(ctx); err != nil {
		return nil, err
	}
	// 检测防重
	if available, err = dao.User.IsAccountAvailable(ctx, account, []uint{user.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.WithCodef(biz.UserAccountExists, account))
	}
	data := &do.User{
		Id:      user.Id,
		Account: account,
	}
	return dao.User.Update(ctx, data)
}

// 修改用户手机号(用于前台)
func (s *sUser) UpdateCurrentMobile(ctx context.Context, mobile string) (*model.User, error) {
	var (
		user      *model.User
		err       error
		available bool
	)
	// 扫描数据
	if user, err = s.GetCurrentUser(ctx); err != nil {
		return nil, err
	}
	// 检测防重
	if available, err = dao.User.IsMobileAvailable(ctx, mobile, []uint{user.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.UserMobileExists)
	}
	data := &do.User{
		Id:     user.Id,
		Mobile: mobile,
	}
	return dao.User.Update(ctx, data)
}

// 修改用户邮箱(用于前台)
func (s *sUser) UpdateCurrentEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		mod       *model.User
		err       error
		available bool
	)
	// 扫描数据
	if mod, err = s.GetCurrentUser(ctx); err != nil {
		return nil, err
	}
	// 检测防重
	if available, err = dao.User.IsEmailAvailable(ctx, email, []uint{mod.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.UserEmailExists)
	}
	data := &do.User{
		Id:    mod.Id,
		Email: email,
	}
	return dao.User.Update(ctx, data)
}

// 修改用户密码(用于前台)
func (s *sUser) UpdateCurrentPassword(ctx context.Context, password string) (*model.User, error) {
	var (
		mod *model.User
		err error
	)
	// 扫描数据
	if mod, err = s.GetCurrentUser(ctx); err != nil {
		return nil, err
	}
	data := &do.User{
		Id:       mod.Id,
		Password: s.mustEncryptPasword(password, grand.Letters(4)),
	}
	return dao.User.Update(ctx, data)
}

// 用户登录
func (s *sUser) Signin(ctx context.Context, in *model.UserSigninInput) (*model.AuthToken, error) {
	var (
		user *model.User
		role *model.AuthRole
		err  error
	)
	switch in.Type {
	case "passport":
		user, err = s.signinWithPassport(ctx, &model.UserSigninPassportInput{
			Passport: in.Passport,
			Password: in.Password,
			Captcha:  in.Captcha,
		})
	case "mobile":
		user, err = s.signinWithMobile(ctx, &model.UserSigninMobileInput{
			Mobile:  in.Passport,
			Captcha: in.Captcha,
		})
	case "email":
		user, err = s.signinWithEmail(ctx, &model.UserSigninEmailInput{
			Email:   in.Passport,
			Captcha: in.Captcha,
		})
	default:
		return nil, gerror.NewCode(biz.UserSigninTypeIcorrect)
	}
	// 登录异常
	if err != nil {
		return nil, err
	}
	// 获取登录角色
	if len(in.Role) > 0 {
		if role, err = dao.AuthRole.GetByName(ctx, in.Role); err != nil {
			return nil, err
		}
	} else {
		if role, err = service.Auth().GetDefaultRole(ctx); err != nil {
			return nil, err
		}
	}
	// 角色不存在
	if role == nil {
		return nil, gerror.NewCode(biz.AuthRoleNotExists)
	}

	// 角色权限受限
	var isExist = true
	if user.Id != consts.RootAdminId && role.Id != consts.RootRoleId && role.Id != consts.DefaultRoleId {
		if isExist, err = dao.AuthAccess.IsExist(ctx, &do.AuthAccess{
			UserId: user.Id,
			RoleId: role.Id,
		}); err != nil {
			return nil, err
		}
	}
	if !isExist {
		return nil, gerror.NewCode(biz.AuthNotPermission)
	}
	// // 角色登录
	// if err = s.SigninDrect(ctx, user, role); err != nil {
	// 	return nil, err
	// }
	return s.SigninDrect(ctx, user, role)
}

// 账户|手机号|邮箱 + 密码 登录
func (s *sUser) signinWithPassport(ctx context.Context, in *model.UserSigninPassportInput) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	// 优先尝试找手机号
	if user, err = dao.User.GetByMobile(ctx, in.Passport); err != nil {
		return nil, err
	}
	// 手机号找不到，尝试找邮箱
	if user == nil {
		if user, err = dao.User.GetByEmail(ctx, in.Passport); err != nil {
			return nil, err
		}
	}
	// 邮箱还是找不到，最后找账号
	if user == nil {
		if user, err = dao.User.GetByAccount(ctx, in.Passport); err != nil {
			return nil, err
		}
	}
	// 没有找到账户
	if user == nil {
		return nil, gerror.NewCode(biz.UserSignIncorrect)
	}
	// 状态异常
	if user.Status != "normal" {
		return nil, gerror.NewCode(biz.UserStatusInvalid)
	}
	// 登录次数受限
	if user.SigninFailure >= consts.MaxSigninFailure {
		return nil, gerror.NewCode(biz.UserSigninFailureInvalid)
	}
	// 账户密码不匹配
	if s.mustEncryptPasword(in.Password, user.Salt) != user.Password {
		// 追加登录错误次数
		if err = dao.User.IncrementSigninFailure(ctx, user.Id); err != nil {
			return nil, err
		}
		return nil, gerror.NewCode(biz.UserSignIncorrect)
	}
	return user, nil
}

// 手机号 + 验证码 登录
func (s *sUser) signinWithMobile(ctx context.Context, in *model.UserSigninMobileInput) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	if user, err = dao.User.GetByMobile(ctx, in.Mobile); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(biz.UserAccountNotExists)
	}
	// 状态异常
	if user.Status != "normal" {
		return nil, gerror.NewCode(biz.UserStatusInvalid)
	}
	// 校验邮箱验证码
	// if err = service.Sms().Verify(ctx, in.Captcha, "mobile", "signin"); err != nil {
	// 	// 追加登录错误次数
	// 	if err = dao.User.IncrementSigninFailure(ctx, user.Id); err != nil {
	// 		return nil, err
	// 	}
	// }
	return user, nil
}

// 邮箱 + 验证码 登录
func (s *sUser) signinWithEmail(ctx context.Context, in *model.UserSigninEmailInput) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	if user, err = dao.User.GetByEmail(ctx, in.Email); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(biz.UserAccountNotExists)
	}
	// 状态异常
	if user.Status != "normal" {
		return nil, gerror.NewCode(biz.UserStatusInvalid)
	}
	// 校验邮箱验证码
	// if err = service.Sms().Verify(ctx, in.Captcha, "email", "signin"); err != nil {
	// 	// 追加登录错误次数
	// 	if err = dao.User.IncrementSigninFailure(ctx, user.Id); err != nil {
	// 		return nil, err
	// 	}
	// }
	return user, nil
}

// 角色登录
func (s *sUser) SigninDrect(ctx context.Context, user *model.User, role *model.AuthRole) (*model.AuthToken, error) {
	var (
		// access  []*model.AuthAccess
		token *model.AuthToken
		err   error
	)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新数据
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			user, err = dao.User.Update(ctx, &do.User{
				Id:            user.Id,
				SigninRole:    role.Name,   // 更新登录角色
				SigninFailure: 0,           // 重置登录错误次数
				SigninAt:      gtime.Now(), // 更新最后登录时间
			})
			return err
		}); err != nil {
			return err
		}
		// 获取Token
		return tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			token, err = service.Auth().Authorization(ctx, &model.AuthAuthorizationInput{
				Subject: user.Uuid,
				Role:    role.Name,
			})
			return err
		})
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// 密码盐加密
func (s *sUser) mustEncryptPasword(password string, salt string) string {
	return gmd5.MustEncryptString(password + salt)
}
