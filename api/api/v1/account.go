package v1

import (
	"letga/api/api/v1/model"

	"github.com/gogf/gf/v2/frame/g"
)

/**
*  用户认证
**/
// 刷新token
type AccountRefreshTokenReq struct {
	g.Meta       `path:"/account/token/refresh" tags:"Api@AccountService" method:"put" summary:"Token refresh"`
	RefreshToken string `json:"refreshToken" v:"required"`
}
type AccountRefreshTokenRes struct {
	model.AuthToken
}

// 注册用户
// - 至少保证一种登录方式原则
// - 账号|手机号|邮箱其中一个必填
// - 当存在账号，则密码必填
// - 当手存在机号|邮箱，则验证码必填
type AccountSignUpReq struct {
	g.Meta   `path:"/account/signup" method:"post" tags:"Api@AccountService" summary:"Sign up"`
	Account  string `json:"account" v:"required-without-all:Mobile,Email|length:4,20"` // 账号
	Password string `json:"password" v:"required-with:Account|length:4,20"`            // 密码
	Nickname string `json:"nickname" v:"required|length:2,12"`                         // 昵称
	Mobile   string `json:"mobile" v:"required-without-all:Account,Email|phone"`       // 手机号
	Captcha  string `json:"captcha" v:"required-with:Mobile,Email|length:4,8"`         // 验证码
	Email    string `json:"email" v:"required-without-all:Account,Mobile|email"`       // 电子邮箱
	Avatar   string `json:"avatar"`                                                    // 头像
}
type AccountSignUpRes struct {
	model.AuthToken
}

// 账户|手机号|邮箱 + 密码 登录
// 登录次数超过后，服务要求 Captcha 验证
type AccountSigninReq struct {
	g.Meta   `path:"/account/signin" tags:"Api@AccountService" method:"post" summary:"Sign in passport"`
	Passport string `json:"passport" v:"required|length:5,18"`          // 账户|手机号|邮箱
	Password string `json:"password" v:"required|password"`             // 密码
	Captcha  string `json:"captcha" v:"length:4,8"`                     // 验证码
	Type     string `json:"type" v:"required|in:passport,mobile,email"` // 登录类型
	Role     string `json:"role"`                                       // 请求登录的角色标识
}
type AccountSigninRes struct {
	model.AuthToken
}

// 手机号 + 短信验证码 登录
// type AccountSignMobileReq struct {
// 	g.Meta  `path:"/account/signin/mobile" tags:"Api@AccountService" method:"post" summary:"Sign in mobile"`
// 	Mobile  string `json:"mobile" v:"required|phone"`       // 手机号
// 	Captcha string `json:"captcha" v:"required:length:4,8"` // 验证码
// }
// type AccountSignMobileRes struct {
// 	model.AuthToken
// }

// 注销登录
type AccountSignoutReq struct {
	g.Meta `path:"/account/signout" method:"put" tags:"Api@AccountService" summary:"Sign out"`
}
type AccountSignoutRes struct{}

// 账户信息
type AccountInfoReq struct {
	g.Meta `path:"/account/info" method:"get" tags:"Api@AccountService" summary:"Get info"`
}
type AccountInfoRes struct {
	model.User
}

// 手机 密码重置
type AccountResetMobileReq struct {
	g.Meta  `path:"/account/password/reset/mobile" tags:"Api@AccountService" method:"post" summary:"Reset password mobile"`
	Email   string `json:"email" v:"required|email"`
	Captcha string `json:"captcha" v:"required|length:4,8"`
}
type AccountResetMobileRes struct {
}

// 邮件 密码重置
type AccountResetEmailReq struct {
	g.Meta  `path:"/account/password/reset/email" tags:"Api@AccountService" method:"post" summary:"Reset password email"`
	Email   string `json:"email" v:"required|email"`
	Captcha string `json:"captcha" v:"required|length:4,8"`
}
type AccountResetEmailRes struct {
}
