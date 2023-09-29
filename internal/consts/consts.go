package consts

const (
	ContextKey             = "ContextKey" // 上下文变量存储键名，前后端系统共享
	RootAdminId            = 1            // 超级管理员ID
	RootRoleId             = 1            // 超级管理员角色ID
	DefaultRoleId          = 2            // 注册时默认拥有的角色ID
	MaxSigninFailure       = 5            // 最大的登录错误次数
	MediaMaxUploadInMinute = 10           // 同一用户1分钟之内最大上传数量
)
