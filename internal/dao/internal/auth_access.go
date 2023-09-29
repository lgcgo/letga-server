// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthAccessDao is the data access object for table auth_access.
type AuthAccessDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AuthAccessColumns // columns contains all the column names of Table for convenient usage.
}

// AuthAccessColumns defines and stores column names for table auth_access.
type AuthAccessColumns struct {
	Id       string // ID
	RoleId   string // 角色ID
	UserId   string // 用户ID
	Status   string // 状态
	CreateAt string // 创建日期
}

// authAccessColumns holds the columns for table auth_access.
var authAccessColumns = AuthAccessColumns{
	Id:       "id",
	RoleId:   "role_id",
	UserId:   "user_id",
	Status:   "status",
	CreateAt: "create_at",
}

// NewAuthAccessDao creates and returns a new DAO object for table data access.
func NewAuthAccessDao() *AuthAccessDao {
	return &AuthAccessDao{
		group:   "default",
		table:   "auth_access",
		columns: authAccessColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthAccessDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthAccessDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthAccessDao) Columns() AuthAccessColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthAccessDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthAccessDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthAccessDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
