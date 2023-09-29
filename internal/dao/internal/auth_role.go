// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRoleDao is the data access object for table auth_role.
type AuthRoleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AuthRoleColumns // columns contains all the column names of Table for convenient usage.
}

// AuthRoleColumns defines and stores column names for table auth_role.
type AuthRoleColumns struct {
	Id       string // ID
	ParentId string // 父ID
	Title    string // 标题
	Name     string // 名称
	Status   string // 状态
	Weight   string // 权重
	CreateAt string // 创建日期
	UpdateAt string // 修改日期
}

// authRoleColumns holds the columns for table auth_role.
var authRoleColumns = AuthRoleColumns{
	Id:       "id",
	ParentId: "parent_id",
	Title:    "title",
	Name:     "name",
	Status:   "status",
	Weight:   "weight",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewAuthRoleDao creates and returns a new DAO object for table data access.
func NewAuthRoleDao() *AuthRoleDao {
	return &AuthRoleDao{
		group:   "default",
		table:   "auth_role",
		columns: authRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthRoleDao) Columns() AuthRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
