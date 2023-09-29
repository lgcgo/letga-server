// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"letga/internal/consts"
	"letga/internal/dao/internal"
	"letga/internal/model"
	"letga/internal/model/do"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalAuthRoleAccessDao is internal type for wrapping internal DAO implements.
type internalAuthRoleAccessDao = *internal.AuthRoleAccessDao

// authRoleAccessDao is the data access object for table auth_role_access.
// You can define custom methods on it to extend its functionality as you wish.
type authRoleAccessDao struct {
	internalAuthRoleAccessDao
}

var (
	// AuthRoleAccess is globally public accessible object for table auth_role_access operations.
	AuthRoleAccess = authRoleAccessDao{
		internal.NewAuthRoleAccessDao(),
	}
)

// Fill with you ideas below.

// 获取所有角色授权
func (d authRoleAccessDao) GetAll(ctx context.Context) ([]*model.AuthRoleAccess, error) {
	var (
		roleAccess []*model.AuthRoleAccess
		err        error
	)
	// 扫描数据
	if err = d.Ctx(ctx).Hook(HashSelectHook(ctx)).Scan(&roleAccess); err != nil {
		return nil, err
	}
	return roleAccess, nil
}

// 获取指定角色路由集
func (d authRoleAccessDao) GetRouteIdsByRoleIds(ctx context.Context, roleIds []uint) ([]uint, error) {
	var (
		routeIds []uint
		vars     []*gvar.Var
		err      error
	)
	// 扫描数据
	vars, err = d.Ctx(ctx).WhereIn(d.Columns().RoleId, roleIds).Fields(d.Columns().RouteId).Array()
	if err != nil {
		return nil, err
	}
	for _, v := range vars {
		routeIds = append(routeIds, v.Uint())
	}
	return routeIds, nil
}

// 批量写入角色授权
func (d authRoleAccessDao) BatchInsert(ctx context.Context, data []*do.AuthRoleAccess) error {
	_, err := d.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return err
	}
	return nil
}

// 批量删除角色授权
func (d authRoleAccessDao) BatchDelete(ctx context.Context, data []*do.AuthRoleAccess) error {
	for _, v := range data {
		_, err := d.Ctx(ctx).Where(do.AuthRoleAccess{
			RoleId:  v.RoleId,
			RouteId: v.RouteId,
		}).Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

// 使用角色ID删除角色授权
func (d authRoleAccessDao) DeleteByRoleIds(ctx context.Context, roleIds []uint) error {
	return d.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := d.Ctx(ctx).WhereIn(d.Columns().RoleId, roleIds).Delete()
		return err
	})
}

// 使用路由ID删除角色授权
func (d authRoleAccessDao) DeleteByRouteIds(ctx context.Context, routeIds []uint) error {
	return d.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := d.Ctx(ctx).WhereIn(d.Columns().RouteId, routeIds).Delete()
		return err
	})
}

// 检测ID集
func (d authRoleAccessDao) CheckIds(ctx context.Context, keyIds []uint) error {
	return CheckIds(ctx, d.Table(), keyIds, consts.TABLE_AUTH_ROLE_ACCESS_SALT)
}

// 检测单个角色授权
func (d authRoleAccessDao) IsRoleRouteExist(ctx context.Context, roleId uint, routeId uint) (bool, error) {
	count, err := d.Ctx(ctx).Where(d.Columns().RoleId, roleId).Where(d.Columns().RouteId, routeId).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}