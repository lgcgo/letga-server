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
	"letga/utility/tree"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 获取角色
func (s *sAuth) GetRole(ctx context.Context, key string) (*model.AuthRole, error) {
	var (
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return nil, err
	}
	return dao.AuthRole.Get(ctx, keyId)
}

// 获取角色
func (s *sAuth) GetDefaultRole(ctx context.Context) (*model.AuthRole, error) {
	return dao.AuthRole.Get(ctx, consts.DefaultRoleId)
}

// 创建角色
func (s *sAuth) CreateRole(ctx context.Context, in *model.AuthRoleCreateInput) (*model.AuthRole, error) {
	var (
		parent    *model.AuthRole
		role      *model.AuthRole
		data      *do.AuthRole
		err       error
		available bool
	)
	// 检测父级
	if len(in.ParentKey) > 0 {
		if parent, err = s.GetRole(ctx, in.ParentKey); err != nil {
			return nil, gerror.NewCode(biz.AuthRoleParentNotExists)
		}
		// 默认角色不能作为父级
		if parent.Id == consts.DefaultRoleId {
			return nil, gerror.NewCode(biz.AuthRoleParentNotExists)
		}
	}
	// 名称防重
	if available, err = dao.AuthRole.IsNameAvailable(ctx, in.Name); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.AuthRoleNameExists)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if parent != nil {
		data.ParentId = parent.Id
	}
	var (
		appendRouteIds []uint
	)
	// 检测新增路由Key集
	if appendRouteIds, err = hashid.BatchDecode(in.AppendRouteKeys, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
		return nil, err
	}
	if err = dao.AuthRoute.CheckIds(ctx, appendRouteIds); err != nil {
		return nil, err
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if role, err = dao.AuthRole.Insert(ctx, data); err != nil {
			return err
		}
		if len(in.AppendRouteKeys) > 0 {
			if _, err = s.AppendRoleAccess(ctx, &model.AuthRoleAccessAppendInput{
				RoleKey:   role.Key,
				RouteKeys: in.AppendRouteKeys,
			}); err != nil {
				return err
			}
		}
		// 更新授权政策
		return s.RefreshPolicy(ctx)
	}); err != nil {
		return nil, err
	}
	return role, nil
}

// 修改角色
func (s *sAuth) UpdateRole(ctx context.Context, in *model.AuthRoleUpdateInput) (*model.AuthRole, error) {
	var (
		handler   *tree.Tree
		list      []*model.AuthRole
		data      *do.AuthRole
		parent    *model.AuthRole
		role      *model.AuthRole
		err       error
		available bool
	)
	// 扫描数据
	if role, err = s.GetRole(ctx, in.Key); err != nil {
		return nil, err
	}

	if role.Id == consts.RootRoleId {
		// 超级管理员，不能拥有父级
		if len(in.ParentKey) > 0 {
			return nil, gerror.NewCode(biz.AuthRoleParentInvalid)
		}
		// 超级管理员，不能拥有路由授权
		if len(in.AppendRouteKeys) > 0 || len(in.RemoveRouteKeys) > 0 {
			return nil, gerror.NewCode(biz.AuthRoleAccessInvalid)
		}
	} else {
		// 非超级管理员必须拥有父级
		if len(in.ParentKey) == 0 {
			return nil, gerror.NewCode(biz.AuthRoleParentInvalid)
		}
		// 获取父级
		if parent, err = s.GetRole(ctx, in.ParentKey); err != nil {
			return nil, err
		}
		// 默认角色不能作为父级
		if parent.Id == consts.DefaultRoleId {
			return nil, gerror.NewCode(biz.AuthRoleParentInvalid)
		}
		// 默认角色的父级必须是root(即不能更改)
		if role.Id == consts.DefaultRoleId && (parent == nil || role.Id == consts.RootRoleId) {
			return nil, gerror.NewCode(biz.AuthRoleParentInvalid)
		}

		// 获取所有数据
		if list, err = dao.AuthRole.GetAll(ctx); err != nil {
			return nil, err
		}
		// 获取树句柄
		if handler, err = tree.NewWithData(list); err != nil {
			return nil, err
		}
		// 父级不能是自身以及子级
		var tempKeys []string
		if tempKeys, err = handler.GetChildKeys(in.Key); err != nil {
			return nil, err
		}
		tempKeys = append(tempKeys, role.Key)
		for _, v := range tempKeys {
			if in.ParentKey == v {
				return nil, gerror.NewCode(biz.AuthRoleParentInvalid)
			}
		}
	}

	// 名称防重
	if available, err = dao.AuthRole.IsNameAvailable(ctx, in.Name, []uint{role.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.AuthRoleNameExists)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if parent != nil {
		data.ParentId = parent.Id
	}
	data.Id = role.Id

	// 嵌套事务
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新实体
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			role, err = dao.AuthRole.Update(ctx, data)
			return err
		}); err != nil {
			return err
		}
		// 创建授权
		if len(in.AppendRouteKeys) > 0 {
			if _, err = s.AppendRoleAccess(ctx, &model.AuthRoleAccessAppendInput{
				RoleKey:   role.Key,
				RouteKeys: in.AppendRouteKeys,
			}); err != nil {
				return err
			}
		}
		// 移除授权
		if len(in.RemoveRouteKeys) > 0 {
			if _, err = s.RemoveRoleAccess(ctx, &model.AuthRoleAccessRemoveInput{
				RoleKey:   role.Key,
				RouteKeys: in.RemoveRouteKeys,
			}); err != nil {
				return err
			}
		}
		// 更新授权政策
		return tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return s.RefreshPolicy(ctx)
		})
	}); err != nil {
		return nil, err
	}
	return role, nil
}

// 更改角色状态
func (s *sAuth) SetRoleStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthRole, error) {
	var (
		handler *tree.Tree
		list    []*model.AuthRole
		role    *model.AuthRole
		err     error
	)
	if role, err = s.GetRole(ctx, in.Key); err != nil {
		return nil, err
	}
	// 限制更改超级管理员与默认角色
	if role.Id == consts.RootRoleId || role.Id == consts.DefaultRoleId {
		return nil, gerror.NewCode(biz.AuthNotPermission)
	}
	// 获取所有数据
	if list, err = dao.AuthRole.GetAll(ctx); err != nil {
		return nil, err
	}
	// 获取树句柄
	if handler, err = tree.NewWithData(list); err != nil {
		return nil, err
	}
	// 获取子Key集
	var (
		keyIds      []uint
		childKeys   []string
		childKeyIds []uint
	)
	if childKeys, err = handler.GetChildKeys(role.Key); err != nil {
		return nil, err
	}
	// Key集解密
	if childKeyIds, err = hashid.BatchDecode(childKeys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return nil, err
	}
	keyIds = append(childKeyIds, role.Id)
	for _, v := range keyIds {
		if _, err = dao.AuthRole.SetStatus(ctx, v, in.Value); err != nil {
			return nil, err
		}
	}
	return s.GetRole(ctx, in.Key)
}

// 删除角色
func (s *sAuth) DeleteRoles(ctx context.Context, keys []string) error {
	var (
		handler *tree.Tree
		list    []*model.AuthRole
		keyIds  []uint
		err     error
	)
	// Key集解密
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return err
	}
	// 检测ID集
	if err = dao.AuthRole.CheckIds(ctx, keyIds); err != nil {
		return err
	}
	// 限制删除超级管理员与默认角色
	for _, v := range keyIds {
		if v == consts.RootRoleId || v == consts.DefaultRoleId {
			return gerror.NewCode(biz.AuthRoleInvalid)
		}
	}
	// 获取所有数据
	if list, err = dao.AuthRole.GetAll(ctx); err != nil {
		return err
	}
	// 获取树句柄
	if handler, err = tree.NewWithData(list); err != nil {
		return err
	}
	// 合并子Key集
	var temKeys = keys
	for _, v := range keys {
		var tempChildKeys []string
		if tempChildKeys, err = handler.GetChildKeys(v); err != nil {
			return err
		}
		temKeys = append(temKeys, tempChildKeys...)
	}
	keys = append(keys, temKeys...)
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
		return err
	}
	// 嵌套事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除实体
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthRole.Delete(ctx, keyIds)
		}); err != nil {
			return err
		}
		// 删除用户授权数据
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return s.DeleteAccessByRoleIds(ctx, keyIds)
		}); err != nil {
			return err
		}
		// 更新授权政策
		return tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return s.RefreshPolicy(ctx)
		})
	})
}

// 获取角色树
func (s *sAuth) GetRoleTree(ctx context.Context, in *model.AuthRoleTreeInput) (*model.TreeOutput, error) {
	var (
		treeHandler     *tree.Tree
		list            []*model.AuthRole
		data            []*tree.TreeNode
		ctxSelectedKeys []string
		ctxDisabledKeys []string
		err             error
	)

	// 获取所有数据
	if list, err = dao.AuthRole.GetAll(ctx); err != nil {
		return nil, err
	}
	// 获取树句柄
	if treeHandler, err = tree.NewWithData(list); err != nil {
		return nil, err
	}
	// 获取树数据
	if data, err = treeHandler.Search(in.Search, in.Keys); err != nil {
		return nil, err
	}

	var (
		role               *model.AuthRole
		roleChildKeys      []string
		roleChildKeysArray *garray.StrArray
	)

	// 获取roleForm中的菜单上下文
	if in.CtxScene == "roleForm" && len(in.CtxRoleKey) > 0 {
		if role, err = s.GetRole(ctx, in.CtxRoleKey); err != nil {
			return nil, err
		}
		if roleChildKeys, err = treeHandler.GetChildKeys(role.Key); err != nil {
			return nil, err
		}
		roleChildKeysArray = garray.NewStrArrayFrom(roleChildKeys)
	}

	var (
		user          *model.User
		access        []*model.AuthAccess
		accessRoleIds = garray.NewIntArray(true)
	)
	// 获取formTransfer中的用户实体
	if in.CtxScene == "roleAccessTransfer" && len(in.CtxUserKey) > 0 {
		if user, err = service.User().GetUser(ctx, in.CtxUserKey); err != nil {
			return nil, err
		}
		if access, err = dao.AuthAccess.GetByUserId(ctx, user.Id); err != nil {
			return nil, err
		}
		for _, v := range access {
			accessRoleIds.Append(int(v.RoleId))
		}
	}

	// 根据上下文组装禁用以及选中项
	for _, v := range list {
		switch in.CtxScene {
		case "mainTable":
			// 禁用默认角色和超级管理员
			if v.Id == consts.RootRoleId || v.Id == consts.DefaultRoleId {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		case "roleAccessTransfer":
			// 禁用默认角色和超级管理员
			if v.Id == consts.RootRoleId || v.Id == consts.DefaultRoleId {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
			// 选中用户已经拥有的角色
			if accessRoleIds.Contains(int(v.Id)) || v.Id == consts.DefaultRoleId {
				ctxSelectedKeys = append(ctxSelectedKeys, v.Key)
			}
		case "roleForm":
			// 禁用自身与下级
			if role != nil && (roleChildKeysArray.Contains(v.Key) || v.Key == role.Key) {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
			// 禁用默认角色
			if v.Id == consts.DefaultRoleId {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		}
	}

	return &model.TreeOutput{
		Data:            data,
		CtxDisabledKeys: ctxDisabledKeys,
		CtxSelectedKeys: ctxSelectedKeys,
	}, nil
}

// 路由统计
func (s *sAuth) GetRoleFliterTree(ctx context.Context, in *model.AuthRoleFliterTreeInput) (*model.FliterTree, error) {
	var (
		data            []*model.FliterTreeItem
		err             error
		count           int
		handler         *tree.Tree
		ctxSelectedKeys []string
		ctxDisabledKeys []string
	)
	// 组装角色组
	var (
		roleList       []*model.AuthRole
		roleTreeSource []*model.FliterTreeData
		roleKeys       []string
		roleIDs        []uint
	)

	if roleList, err = dao.AuthRole.GetAll(ctx); err != nil {
		return nil, err
	}
	if handler, err = tree.NewWithData(roleList); err != nil {
		return nil, gerror.NewCode(biz.TreeInitialFail)
	}
	for _, v := range roleList {
		if roleKeys, err = handler.GetChildKeys(v.Key); err != nil {
			return nil, err
		}
		roleKeys = append(roleKeys, v.Key)
		if roleIDs, err = hashid.BatchDecode(roleKeys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
			return nil, err
		}
		switch in.CtxScene {
		case "accessSidebar":
			if count, err = dao.AuthAccess.Ctx(ctx).WhereIn(dao.AuthAccess.Columns().RoleId, roleIDs).Count(); err != nil {
				return nil, err
			}
			roleTreeSource = append(roleTreeSource, &model.FliterTreeData{
				Key:       v.Key,
				ParentKey: v.ParentKey,
				Title:     v.Title,
				Count:     count,
				Weight:    v.Weight,
			})
			// 可用根据上下文需求设置选中以及禁用项
		}
	}
	if handler, err = tree.NewWithData(roleTreeSource); err != nil {
		return nil, gerror.NewCode(biz.TreeInitialFail)
	}
	data = append(data, &model.FliterTreeItem{
		Name:  "role",
		Data:  handler.GetWholeTree(),
		Total: len(roleTreeSource),
	})

	return &model.FliterTree{
		Data:            data,
		CtxSelectedKeys: ctxSelectedKeys,
		CtxDisabledKeys: ctxDisabledKeys,
	}, nil
}
