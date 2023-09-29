// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRouteDao is the data access object for table auth_route.
type AuthRouteDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AuthRouteColumns // columns contains all the column names of Table for convenient usage.
}

// AuthRouteColumns defines and stores column names for table auth_route.
type AuthRouteColumns struct {
	Id       string // ID
	MenuId   string // 菜单ID
	Title    string // 标题
	Path     string // 路由地址
	Method   string // 请求方法
	Remark   string // 备注
	Status   string // 状态
	Weight   string // 权重
	CreateAt string // 创建日期
	UpdateAt string // 更新日期
}

// authRouteColumns holds the columns for table auth_route.
var authRouteColumns = AuthRouteColumns{
	Id:       "id",
	MenuId:   "menu_id",
	Title:    "title",
	Path:     "path",
	Method:   "method",
	Remark:   "remark",
	Status:   "status",
	Weight:   "weight",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewAuthRouteDao creates and returns a new DAO object for table data access.
func NewAuthRouteDao() *AuthRouteDao {
	return &AuthRouteDao{
		group:   "default",
		table:   "auth_route",
		columns: authRouteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthRouteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthRouteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthRouteDao) Columns() AuthRouteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthRouteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthRouteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthRouteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
