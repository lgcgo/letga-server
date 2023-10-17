// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package casbin

import (
	"bufio"
	"errors"
	"os"
	"strings"

	casbinpkg "github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

const ROLE_NAME_PERFIX = "role::"
const ROOT_ROLE_NAME = "root"

// 字符串模型
const modelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act || r.sub == "` + ROLE_NAME_PERFIX + ROOT_ROLE_NAME + `"
`

// 配置项
type Config struct {
	PolicyFilePath string
}

var (
	config   *Config
	adapter  persist.Adapter
	enforcer *casbinpkg.Enforcer
	model    casbinmodel.Model
	err      error
)

func Init(cfg *Config) {
	// 使用字符串获取 Casbin模型
	if model, err = casbinmodel.NewModelFromString(modelText); err != nil {
		panic(err.Error())
	}
	// 获取 Casbin执行器
	adapter = fileadapter.NewAdapter(cfg.PolicyFilePath)
	if enforcer, err = casbinpkg.NewEnforcer(model, adapter); err != nil {
		panic(err.Error())
	}
	config = cfg
}

// 路由政策
type RoutePolicy struct {
	Role   string // 用户角色
	Path   string // 资源路径
	Method string // 请求方法
}

// 格式化资源政策
func (r *RoutePolicy) PolicyFormat() string {
	var (
		strArr []string
	)

	strArr = append(strArr, "p")
	strArr = append(strArr, ROLE_NAME_PERFIX+r.Role)
	strArr = append(strArr, r.Path)
	strArr = append(strArr, r.Method)

	return strings.Join(strArr, ", ")
}

// 角色政策
type RolePolicy struct {
	ParentRole string // 父级角色名称
	Role       string // 角色名称
}

// 格式化角色政策
func (r *RolePolicy) PolicyFormat() string {
	var strArr []string

	strArr = append(strArr, "g")
	if len(r.ParentRole) == 0 {
		// 默认挂载的超级管理员
		strArr = append(strArr, ROLE_NAME_PERFIX+ROOT_ROLE_NAME)
	} else {
		strArr = append(strArr, ROLE_NAME_PERFIX+r.ParentRole)
	}
	strArr = append(strArr, ROLE_NAME_PERFIX+r.Role)
	strArr = append(strArr, "*")

	return strings.Join(strArr, ", ")
}

// 用户政策
// type UserPolicy struct {
// 	Subject string // 用户主题
// 	Role    string // 角色名称
// }

// 格式化用户政策
// func (r *UserPolicy) PolicyFormat() string {
// 	var strArr []string

// 	strArr = append(strArr, "g")
// 	strArr = append(strArr, r.Subject)
// 	strArr = append(strArr, ROLE_NAME_PERFIX+r.Role)

// 	return strings.Join(strArr, ", ")
// }

type VerifyPayload struct {
	Subject string // 授权主题
	Role    string // 授权角色
	Path    string // 请求路径
	Method  string // 请求方法
}

type CsvPlicyPaylod struct {
	RoutePolicys []*RoutePolicy
	RolePolicys  []*RolePolicy
	// UserPolicys  []*UserPolicy
}

// 检测角色的路由权限
func Verify(p *VerifyPayload) (bool, error) {
	var (
		ok  bool
		err error
	)

	// 验证角色权限
	if ok, err = enforcer.Enforce(ROLE_NAME_PERFIX+p.Role, p.Path, p.Method); err != nil {
		return false, err
	}
	// if !ok {
	// 	// 验证用户权限
	// 	if ok, err = c.Enforcer.Enforce(p.Subject p.Path, p.Method); err != nil {
	// 		return false, err
	// 	}
	// }

	return ok, nil
}

// 刷新授权
func Refresh() error {
	if config == nil {
		return errors.New("missing config")
	}
	Init(config)
	return nil
}

// 更新Policy.csv文件
func SavePolicyCsv(payload *CsvPlicyPaylod) error {
	var (
		file   *os.File
		writer *bufio.Writer
		err    error
	)

	// 获取文件句柄
	file, err = os.OpenFile(config.PolicyFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件
	writer = bufio.NewWriter(file)
	for _, v := range payload.RoutePolicys {
		writer.WriteString(v.PolicyFormat())
		writer.WriteString("\n")
	}
	for _, v := range payload.RolePolicys {
		writer.WriteString(v.PolicyFormat())
		writer.WriteString("\n")
	}
	// for _, v := range payload.UserPolicys {
	// 	writer.WriteString(v.PolicyFormat())
	// 	writer.WriteString("\n")
	// }
	writer.Flush()
	return Refresh()
}
