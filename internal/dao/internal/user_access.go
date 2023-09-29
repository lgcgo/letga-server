// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserAccessDao is the data access object for table user_access.
type UserAccessDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserAccessColumns // columns contains all the column names of Table for convenient usage.
}

// UserAccessColumns defines and stores column names for table user_access.
type UserAccessColumns struct {
	Id       string // ID
	UserId   string // 用户ID
	RoleId   string // 角色ID
	Status   string // 状态
	CreateAt string // 创建日期
}

// userAccessColumns holds the columns for table user_access.
var userAccessColumns = UserAccessColumns{
	Id:       "id",
	UserId:   "user_id",
	RoleId:   "role_id",
	Status:   "status",
	CreateAt: "create_at",
}

// NewUserAccessDao creates and returns a new DAO object for table data access.
func NewUserAccessDao() *UserAccessDao {
	return &UserAccessDao{
		group:   "default",
		table:   "user_access",
		columns: userAccessColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserAccessDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserAccessDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserAccessDao) Columns() UserAccessColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserAccessDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserAccessDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserAccessDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
