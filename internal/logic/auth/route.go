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

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 创建路由
func (s *sAuth) CreateRoute(ctx context.Context, in *model.AuthRouteCreateInput) (*model.AuthRoute, error) {
	var (
		data      *do.AuthRoute
		route     *model.AuthRoute
		available bool
		err       error
		menu      *model.Menu
	)
	// 检测菜单
	if len(in.MenuKey) > 0 {
		if menu, err = service.Menu().GetMenu(ctx, in.MenuKey); err != nil {
			return nil, err
		}
	}
	// 路径防重
	if available, err = dao.AuthRoute.IsPathMethodAvailable(ctx, in.Path, in.Method); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.WithCodef(biz.AuthRouteExists, in.Method+":"+in.Path))
	}
	// 写入数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if menu != nil {
		data.MenuId = menu.Id
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		route, err = dao.AuthRoute.Insert(ctx, data)
		return err
	}); err != nil {
		return nil, err
	}
	return route, nil
}

// 获取路由
func (s *sAuth) GetRoute(ctx context.Context, key string) (*model.AuthRoute, error) {
	var (
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
		return nil, err
	}
	return dao.AuthRoute.Get(ctx, keyId)
}

// 修改路由
func (s *sAuth) UpdateRoute(ctx context.Context, in *model.AuthRouteUpdateInput) (*model.AuthRoute, error) {
	var (
		data      *do.AuthRoute
		route     *model.AuthRoute
		menu      *model.Menu
		err       error
		available bool
	)
	// 扫描数据
	if route, err = s.GetRoute(ctx, in.Key); err != nil {
		return nil, err
	}
	if route == nil {
		return nil, gerror.NewCode(biz.AuthRouteNotExists)
	}
	// 检测菜单
	if len(in.MenuKey) > 0 {
		if menu, err = service.Menu().GetMenu(ctx, in.MenuKey); err != nil {
			return nil, err
		}
	}
	// 路径防重
	if available, err = dao.AuthRoute.IsPathMethodAvailable(ctx, in.Path, in.Method, []uint{route.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.WithCodef(biz.AuthRouteExists, in.Path+":"+in.Method))
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	data.Id = route.Id
	if menu != nil {
		data.MenuId = menu.Id
	}

	// 嵌套事务
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新实体
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			route, err = dao.AuthRoute.Update(ctx, data)
			return err
		}); err != nil {
			return err
		}
		// 更新授权政策
		return tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return s.RefreshPolicy(ctx)
		})
	}); err != nil {
		return nil, err
	}
	return route, nil
}

// 更改域状态
func (s *sAuth) SetRouteStatus(ctx context.Context, in *model.StatusSetInput) (*model.AuthRoute, error) {
	var (
		route *model.AuthRoute
		err   error
	)
	if route, err = s.GetRoute(ctx, in.Key); err != nil {
		return nil, err
	}
	if route == nil {
		return nil, gerror.NewCode(biz.AuthRouteNotExists)
	}
	// 限制禁止
	// if  {
	// }
	return dao.AuthRoute.SetStatus(ctx, route.Id, in.Value)
}

// 删除路由
func (s *sAuth) DeleteRoutes(ctx context.Context, keys []string) error {
	var (
		routeIds []uint
		err      error
	)
	// Key集解密
	if routeIds, err = hashid.BatchDecode(keys, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
		return err
	}
	// 检测ID集
	if err = dao.AuthRoute.CheckIds(ctx, routeIds); err != nil {
		return err
	}
	// 嵌套事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除实体
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthRoute.Delete(ctx, routeIds)
		}); err != nil {
			return err
		}
		// 删除关联角色权限数据
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthRoleAccess.DeleteByRouteIds(ctx, routeIds)
		}); err != nil {
			return err
		}
		// 更新授权政策
		return tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return s.RefreshPolicy(ctx)
		})
	})
}

// 获取路由分页
func (s *sAuth) GetRoutePage(ctx context.Context, in *model.AuthRoutePageInput) (*model.AuthRoutePageOutput, error) {
	var (
		data            []*model.AuthRoute
		count           int
		ctxSelectedKeys []string
		ctxDisabledKeys []string
		ctxDirectKeys   []string
		menuIDs         []uint
		err             error
		ctxRole         *model.AuthRole
		// ctxRoleIds           []uint
		ctxRoleRouteIds      []uint
		ctxRoleChildRouteIds []uint
		ctxRouteIds          []uint
	)

	// 组装Sider菜单筛选OR条件
	if len(in.Mnks) > 0 {
		if menuIDs, err = hashid.BatchDecode(in.Mnks, consts.TABLE_MENU_SALT); err != nil {
			return nil, err
		}
	}

	var (
		rmksIDs []uint
		apksIDs []uint
	)
	// 获取指定移除项ID集
	if len(in.Rmks) > 0 {
		// 解析Key集
		if rmksIDs, err = hashid.BatchDecode(in.Rmks, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
			return nil, err
		}
		// 检测ID集
		if err = dao.AuthRoute.CheckIds(ctx, rmksIDs); err != nil {
			return nil, err
		}
	}
	// 获取指定追加项ID集
	if len(in.Apks) > 0 {
		// 解析Key集
		if apksIDs, err = hashid.BatchDecode(in.Apks, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
			return nil, err
		}
		// 检测ID集
		if err = dao.AuthRoute.CheckIds(ctx, apksIDs); err != nil {
			return nil, err
		}
	}

	// 获取角色授权中的上下文角色
	if (in.CtxScene == "roleAccessTable" || in.CtxScene == "roleAccessModel") && len(in.CtxRoleKey) > 0 {
		// 获取上下文角色
		if ctxRole, err = s.GetRole(ctx, in.CtxRoleKey); err != nil {
			return nil, err
		}
		if ctxRole == nil {
			return nil, gerror.NewCode(biz.AuthRoleInvalid)
		}
		var (
			treeHandler   *tree.Tree
			roles         []*model.AuthRole
			roleChildKeys []string
			roleChildIds  []uint
		)
		// 获取所有数据
		if roles, err = dao.AuthRole.GetAll(ctx); err != nil {
			return nil, err
		}
		// 获取树句柄
		if treeHandler, err = tree.NewWithData(roles); err != nil {
			return nil, err
		}
		if roleChildKeys, err = treeHandler.GetChildKeys(ctxRole.Key); err != nil {
			return nil, err
		}
		if roleChildIds, err = hashid.BatchDecode(roleChildKeys, consts.TABLE_AUTH_ROLE_SALT); err != nil {
			return nil, err
		}
		// 获取当前角色的路由
		ctxRoleRouteIds, err = dao.AuthRoleAccess.GetRouteIdsByRoleIds(ctx, []uint{ctxRole.Id})
		if err != nil {
			return nil, err
		}
		// 获取当前角色子级的路由
		ctxRoleChildRouteIds, err = dao.AuthRoleAccess.GetRouteIdsByRoleIds(ctx, roleChildIds)
		if err != nil {
			return nil, err
		}
		ctxRouteIds = append(ctxRoleRouteIds, ctxRoleChildRouteIds...)
	}

	// 条件处理器
	var handler = func(m *gdb.Model) *gdb.Model {
		// 组装菜单筛选查询条件
		if len(menuIDs) > 0 {
			m = m.WhereOrIn(dao.AuthRoute.Columns().MenuId, menuIDs)
		}
		// 组装角色授权关联查询条件
		if in.CtxScene == "roleAccessTable" {
			m = m.WhereIn(dao.AuthRoute.Columns().Id, ctxRouteIds)
		}
		// 组装追加项查询条件
		if len(apksIDs) > 0 {
			m = m.WhereOrIn(dao.AuthRoute.Columns().Id, apksIDs)
		}
		// 组装移除项查询条件
		if len(rmksIDs) > 0 {
			m = m.WhereNotIn(dao.AuthRoute.Columns().Id, rmksIDs)
		}
		return m
	}

	// 获取分页数据
	if data, count, err = dao.AuthRoute.PageData(ctx, &in.PageParams, &in.AuthRouteSearch, handler); err != nil {
		return nil, err
	}

	// 设置上下文禁用项与选中项
	var (
		ctxRoleRouteIdsSet = gset.NewFrom(ctxRoleRouteIds)
		ctxRouteIdsSet     = gset.NewFrom(ctxRouteIds)
	)
	for _, v := range data {
		switch in.CtxScene {
		case "mainTable":
			// 什么也不做
		case "roleAccessTable":
			// 判断是否是子集的路由
			if ctxRoleRouteIdsSet.Contains(v.Id) {
				ctxDirectKeys = append(ctxDirectKeys, v.Key)
			}
		case "roleAccessModel":
			// 判断是否属于上下文路由
			if ctxRouteIdsSet.Contains(v.Id) {
				ctxSelectedKeys = append(ctxSelectedKeys, v.Key)
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		}
	}

	return &model.AuthRoutePageOutput{
		Data: data,
		Page: model.Page{
			Total:           count,
			CtxSelectedKeys: ctxSelectedKeys,
			CtxDisabledKeys: ctxDisabledKeys,
		},
		CtxDirectKeys: ctxDirectKeys,
	}, nil
}
