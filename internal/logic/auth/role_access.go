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
	"letga/utility/hashid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 设置角色授权
func (s *sAuth) AppendRoleAccess(ctx context.Context, in *model.AuthRoleAccessAppendInput) (*model.AuthRoleAccessAppendOutput, error) {
	var (
		out      *model.AuthRoleAccessAppendOutput
		role     *model.AuthRole
		data     []*do.AuthRoleAccess
		routeIds []uint
		isExist  bool
		err      error
	)
	// 获取角色
	if role, err = s.GetRole(ctx, in.RoleKey); err != nil {
		return nil, gerror.NewCode(biz.AuthRoleParentNotExists)
	}
	// 检测路由Key集
	if routeIds, err = hashid.BatchDecode(in.RouteKeys, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
		return nil, err
	}
	if err = dao.AuthRoute.CheckIds(ctx, routeIds); err != nil {
		return nil, err
	}
	// 组装新增授权数据
	for _, v := range routeIds {
		data = append(data, &do.AuthRoleAccess{
			RoleId:  role.Id,
			RouteId: v,
		})
	}
	// 检测授权是否已经存在
	for _, v := range data {
		if isExist, err = dao.AuthRoleAccess.IsRoleRouteExist(ctx, v.RoleId.(uint), v.RouteId.(uint)); err != nil {
			return nil, err
		}
		// 如果存在则返回错误
		if isExist {
			return nil, gerror.NewCode(biz.AuthRoleAccessExists)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 写入数据
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthRoleAccess.BatchInsert(ctx, data)
		}); err != nil {
			return err
		}
		// // 更新授权政策
		// return s.SaveLocalPolicy(ctx)
		return nil
	}); err != nil {
		return nil, err
	}

	return out, nil
}

// 移除角色授权
func (s *sAuth) RemoveRoleAccess(ctx context.Context, in *model.AuthRoleAccessRemoveInput) (*model.AuthRoleAccessRemoveOutput, error) {
	var (
		out      *model.AuthRoleAccessRemoveOutput
		role     *model.AuthRole
		data     []*do.AuthRoleAccess
		routeIds []uint
		isExist  bool
		err      error
	)
	// 获取角色
	if role, err = s.GetRole(ctx, in.RoleKey); err != nil {
		return nil, gerror.NewCode(biz.AuthRoleParentNotExists)
	}
	// 检测路由Key集
	if routeIds, err = hashid.BatchDecode(in.RouteKeys, consts.TABLE_AUTH_ROUTE_SALT); err != nil {
		return nil, err
	}
	if err = dao.AuthRoute.CheckIds(ctx, routeIds); err != nil {
		return nil, err
	}
	// 组装移除授权数据
	for _, v := range routeIds {
		data = append(data, &do.AuthRoleAccess{
			RoleId:  role.Id,
			RouteId: v,
		})
	}
	// 检测授权是否已经存在
	for _, v := range data {
		if isExist, err = dao.AuthRoleAccess.IsRoleRouteExist(ctx, v.RoleId.(uint), v.RouteId.(uint)); err != nil {
			return nil, err
		}
		// 如果不存在则返回错误
		if !isExist {
			return nil, gerror.NewCode(biz.AuthRoleAccessNotExists)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 写入数据
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.AuthRoleAccess.BatchDelete(ctx, data)
		}); err != nil {
			return err
		}
		// // 更新授权政策
		// return s.SaveLocalPolicy(ctx)
		return nil
	}); err != nil {
		return nil, err
	}

	return out, nil
}
