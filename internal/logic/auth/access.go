// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package auth

import (
	"context"
	"letga/internal/biz"
	"letga/internal/consts"
	"letga/internal/dao"
	"letga/internal/model"
	"letga/internal/model/do"
	"letga/internal/service"
	"letga/utility/hashid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 批量创建用户权限
func (s *sAuth) SetupAccess(ctx context.Context, in *model.AuthAccessSetupInput) error {
	var (
		user *model.User
		err  error
	)
	// 获取用户
	if user, err = service.User().GetUser(ctx, in.UserKey); err != nil {
		return err
	}

	var (
		appendRoleIds []uint
		removeRoleIds []uint
	)
	// 检测新增角色Key集
	if appendRoleIds, err = hashid.BatchDecode(in.AppendRoleKeys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return err
	}
	if err = dao.AuthRoute.CheckIds(ctx, appendRoleIds); err != nil {
		return err
	}

	// 检测移除角色Key集
	if removeRoleIds, err = hashid.BatchDecode(in.RemoveRoleKeys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return err
	}
	if err = dao.AuthRoute.CheckIds(ctx, appendRoleIds); err != nil {
		return err
	}

	// 组装新增数据
	var (
		insertData []*do.AuthAccess
		isExist    bool
		role       *model.AuthRole
	)
	for _, v := range appendRoleIds {
		data := &do.AuthAccess{
			UserId: user.Id,
			RoleId: v,
		}
		if isExist, err = dao.AuthAccess.IsExist(ctx, data); err != nil {
			return err
		}
		if isExist {
			if role, err = dao.AuthRole.Get(ctx, v); err != nil {
				return err
			}
			return gerror.NewCode(biz.WithCodef(biz.AuthAccessRoleInvalid, role.Name))
		}
		insertData = append(insertData, &do.AuthAccess{
			UserId: user.Id,
			RoleId: v,
		})
	}
	// 检测是否已经存在授权

	// 嵌套事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 移除数据
		if len(removeRoleIds) > 0 {
			if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				return dao.AuthAccess.DeleteByUserIdAndRoleIds(ctx, user.Id, removeRoleIds)
			}); err != nil {
				return err
			}
		}
		// 新增数据
		if len(insertData) > 0 {
			if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = dao.AuthAccess.BatchInsert(ctx, insertData)
				return err
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

// 更改用户状态
func (s *sAuth) SetAccessStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthAccess, error) {
	var (
		access *model.AuthAccess
		err    error
	)
	if access, err = s.GetAccess(ctx, in.Key); err != nil {
		return nil, err
	}
	// 限制禁止条件
	// if {
	// }
	return dao.AuthAccess.SetStatus(ctx, access.Id, in.Value)
}

// 获取用户授权
func (s *sAuth) GetAccess(ctx context.Context, key string) (*model.AuthAccess, error) {
	var (
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_AUTH_ACCESS_SALT); err != nil {
		return nil, err
	}
	return dao.AuthAccess.Get(ctx, keyId)
}

// 获取授权分页
func (s *sAuth) GetAccessPage(ctx context.Context, in *model.AuthAccessPageInput) (*model.AuthAccessPageOutput, error) {
	var (
		data            []*model.AuthAccess
		count           int
		ctxDisabledKeys []string
		err             error
		roleIDs         []uint
	)

	// 组装Sider角色筛选项
	if len(in.Roks) > 0 {
		if roleIDs, err = hashid.BatchDecode(in.Roks, consts.TABLE_AUTH_ROLE_SALT); err != nil {
			return nil, err
		}
	}

	// 搜索处理
	// 注意：底层的搜索处理器handler并未适配关联表非直属字段的搜索，这里是临时解决办法
	var (
		searchUserIds []uint
		searchRoleIds []uint
	)
	if len(in.Search) > 0 {
		// 用户昵称搜索
		if searchUserIds, err = dao.User.SearchIdsByNickName(ctx, in.Search); err != nil {
			return nil, err
		}
		// 角色标题搜索
		if searchRoleIds, err = dao.AuthRole.SearchIdsByTitle(ctx, in.Search); err != nil {
			return nil, err
		}
	}

	// 条件处理器
	var handler = func(m *gdb.Model) *gdb.Model {
		if len(roleIDs) > 0 {
			m = m.WhereOrIn(dao.AuthAccess.Columns().RoleId, roleIDs)
		}
		if len(searchUserIds) > 0 {
			m = m.WhereOrIn(dao.AuthAccess.Columns().UserId, searchUserIds)
		}
		if len(searchRoleIds) > 0 {
			m = m.WhereOrIn(dao.AuthAccess.Columns().RoleId, searchRoleIds)
		}
		return m
	}
	if data, count, err = dao.AuthAccess.PageData(ctx, &in.PageParams, &in.AuthAccessSearch, handler); err != nil {
		return nil, err
	}

	// 禁用项
	for _, v := range data {
		switch in.CtxScene {
		case "mainTable":
			if v.RoleId == consts.RootRoleId {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		}
	}
	return &model.AuthAccessPageOutput{
		Data: data,
		Page: model.Page{
			Total:           count,
			CtxDisabledKeys: ctxDisabledKeys,
		},
	}, nil
}

// 删除用户授权
func (s *sAuth) DeleteAccesses(ctx context.Context, keys []string) error {
	var (
		err    error
		keyIds []uint
	)
	// Key集解密
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_AUTH_ACCESS_SALT); err != nil {
		return err
	}
	// 检测ID集
	if err = dao.AuthAccess.CheckIds(ctx, keyIds); err != nil {
		return err
	}
	return dao.AuthAccess.Delete(ctx, keyIds)
}

// 使用角色ID集删除
func (s *sAuth) DeleteAccessByRoleIds(ctx context.Context, roleIds []uint) error {
	return dao.AuthAccess.DeleteByRoleIds(ctx, roleIds)
}
