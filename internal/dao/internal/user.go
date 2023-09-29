// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id            string // ID
	Uuid          string // 唯一ID
	Account       string // 账号
	Mobile        string // 手机号
	Email         string // 电子邮箱
	Password      string // 密码
	Salt          string // 密码盐
	Nickname      string // 昵称
	Avatar        string // 头像
	Signature     string // 个性签名
	SigninRole    string // 登录角色
	SigninFailure string // 失败次数
	SigninIp      string // 登录IP
	SigninAt      string // 登录日期
	Status        string // 状态
	CreateAt      string // 创建日期
	UpdateAt      string // 更新日期
	DeletedAt     string // 删除日期
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:            "id",
	Uuid:          "uuid",
	Account:       "account",
	Mobile:        "mobile",
	Email:         "email",
	Password:      "password",
	Salt:          "salt",
	Nickname:      "nickname",
	Avatar:        "avatar",
	Signature:     "signature",
	SigninRole:    "signin_role",
	SigninFailure: "signin_failure",
	SigninIp:      "signin_ip",
	SigninAt:      "signin_at",
	Status:        "status",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
	DeletedAt:     "deleted_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
