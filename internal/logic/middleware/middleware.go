// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package middleware

import (
	"letga/internal/biz"
	"letga/internal/model"
	"letga/internal/service"
	"letga/utility/response"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &biz.Context{
		User: &biz.ContextUser{},
		Data: make(g.Map),
	}
	// 注册业务上下文
	biz.RegisterCtx()
	// 初始化业务上下文
	biz.Ctx().Init(r, customCtx)

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 响应数据
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err     = r.GetError()
		res     = r.GetHandlerResponse()
		code    = gerror.Code(err)
		bizCode biz.BizCode
		ok      bool
	)
	// 转换自定义错误码
	if err != nil {
		switch code {
		case gcode.CodeValidationFailed:
			code = biz.WithCodef(biz.CodeValidationFailed, err.Error())
		case gcode.CodeNotAuthorized:
			code = biz.CodeNotAuthorized
		default:
			// do nothing
		}
	} else {
		code = biz.CodeOk
	}
	if bizCode, ok = code.(biz.BizCode); !ok {
		bizCode = biz.CodeInternal.(biz.BizCode)
	}
	// 设置状态码
	r.Response.Writer.WriteHeader(bizCode.Code())

	if code == biz.CodeOk {
		response.Success(r, bizCode.BizDetail().ErrorCode, bizCode.Message(), res)
	} else {
		response.ErrorExit(r, bizCode.BizDetail().ErrorCode, bizCode.Message(), res)
	}

}

// 权限认证
func (s *sMiddleware) Authentication(r *ghttp.Request) {
	var (
		ctx         = r.Context()
		code        = biz.CodeOk
		tokenTicket string
	)
	// Header传值 Authorization: Bearer <token>
	if r.Header.Get("Authorization") == "" {
		code = biz.AuthHeaderInvalid
	} else {
		strArr := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		// 支持Bearer方案
		if strArr[0] != "Bearer" || len(strArr[1]) == 0 {
			code = biz.AuthHeaderInvalid
		} else {
			tokenTicket = strArr[1]
		}
	}
	var (
		claims map[string]interface{}
		role   string
		uuid   string
		err    error
	)
	// 验证授权
	if code == biz.CodeOk {
		if claims, err = service.Auth().VerifyToken(ctx, tokenTicket); err == nil {
			// 从签名中获取用户角色
			role = claims["isr"].(string)
			uuid = claims["sub"].(string)
		} else {
			code = gerror.Code(err)
		}
	}
	// 验证路由
	if code == biz.CodeOk {
		if err = service.Auth().Verify(ctx, &model.AuthVerifyInput{
			Subject: uuid,
			Path:    r.URL.Path,
			Method:  r.Method,
			Role:    role,
		}); err == nil {
			// 设置上下文用户
			biz.Ctx().SetUser(r.Context(), &biz.ContextUser{
				Uuid: uuid,
				Role: role,
			})
		} else {
			code = gerror.Code(err)
		}
	}

	// 错误拦截
	if code != biz.CodeOk {
		var (
			bizCode biz.BizCode
			ok      bool
		)
		if bizCode, ok = code.(biz.BizCode); !ok {
			bizCode = biz.CodeInternal.(biz.BizCode)
		}
		// 设置状态码
		r.Response.Writer.WriteHeader(bizCode.Code())
		response.ErrorExit(r, bizCode.BizDetail().ErrorCode, bizCode.Message(), nil)
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
