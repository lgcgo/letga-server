package model

// 用户返回的数据项
type User struct {
	Key           string `json:"key"`           //用户Id
	Uuid          string `json:"uuid"`          // 唯一ID
	Account       string `json:"account"`       // 账号
	Nickname      string `json:"nickname"`      // 昵称
	Avatar        string `json:"avatar"`        // 头像
	Mobile        string `json:"mobile"`        // 手机号
	Email         string `json:"email"`         // 电子邮箱
	Signature     string `json:"signature"`     // 个性签名
	SigninFailure uint   `json:"signinFailure"` // 登录失败次数
	SigninRole    string `json:"signinRole"`    // 登录角色
	SigninIp      string `json:"signinIp"`      // 登录IP
	SigninAt      string `json:"signinAt"`      // 登录日期
	Status        string `json:"status"`        // 状态
	CreateAt      string `json:"createAt"`      // 创建日期
	UpdateAt      string `json:"updateAt"`      // 更新日期
}

// 角色授权
// type UserAccess struct {
// 	Key       string      `json:"key"`       // 索引
// 	UserKey   string      `json:"userKey"`   // 用户索引
// 	User      *User       `json:"user"`      // 用户模型
// 	RoleKey   string      `json:"roleKey"`   // 角色模型
// 	Role      *AuthRole   `json:"role"`      // 角色对象
// 	Status    string      `json:"status"`    // 状态
// 	CreateAt  string      `json:"createAt"`  // 创建日期
// }
