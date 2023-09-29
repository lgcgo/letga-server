// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for table menu.
type MenuDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns MenuColumns // columns contains all the column names of Table for convenient usage.
}

// MenuColumns defines and stores column names for table menu.
type MenuColumns struct {
	Id       string // ID
	ParentId string // 父ID
	Title    string // 标题
	Icon     string // 图标
	CoverUrl string // 封面图片
	Remark   string // 描述
	Status   string // 状态
	Weight   string // 权重
	CreateAt string // 创建日期
	UpdateAt string // 更新日期
}

// menuColumns holds the columns for table menu.
var menuColumns = MenuColumns{
	Id:       "id",
	ParentId: "parent_id",
	Title:    "title",
	Icon:     "icon",
	CoverUrl: "cover_url",
	Remark:   "remark",
	Status:   "status",
	Weight:   "weight",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "menu",
		columns: menuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
