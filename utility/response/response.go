// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	SILENT        = 0 // 不提示错误
	WARN_MESSAGE  = 1 // 警告信息提示
	ERROR_MESSAGE = 2 // 错误信息提示
	NOTIFICATION  = 4 // 通知提示
	REDIRECT      = 9 // 页面跳转
)

// JsonRes 数据返回通用JSON数据结构
// 参考 https://pro.ant.design/zh-CN/docs/request#统一规范
type JsonRes struct {
	Success      bool        `json:"success"`      // 是否请求成功
	Data         interface{} `json:"data"`         // 返回数据(业务接口定义具体数据结构)
	ErrorCode    int         `json:"errorCode"`    // 错误码((0:成功, 1:失败, >1:错误码))
	ErrorMessage string      `json:"errorMessage"` // 错误提示信息
	ShowType     int         `json:"showType"`     // 错误提示类型
	Redirect     string      `json:"redirect"`     // 引导客户端跳转到指定路由
}

func Success(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	Json(r, true, errorCode, errorMessage, SILENT, data...)
}

func SuccessExit(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	JsonExit(r, true, errorCode, errorMessage, SILENT, data...)
}

func Warn(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	Json(r, false, errorCode, errorMessage, WARN_MESSAGE, data...)
}

func WarnExit(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	JsonExit(r, false, errorCode, errorMessage, WARN_MESSAGE, data...)
}

func Error(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	Json(r, false, errorCode, errorMessage, ERROR_MESSAGE, data...)
}

func ErrorExit(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	JsonExit(r, false, errorCode, errorMessage, ERROR_MESSAGE, data...)
}

func Notification(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	Json(r, false, errorCode, errorMessage, NOTIFICATION, data...)
}

func NotificationExit(r *ghttp.Request, errorCode int, errorMessage string, data ...interface{}) {
	JsonExit(r, false, errorCode, errorMessage, NOTIFICATION, data...)
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, success bool, errorCode int, errorMessage string, showType int, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Success:      success,
		Data:         responseData,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		ShowType:     showType,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, success bool, errorCode int, errorMessage string, showType int, data ...interface{}) {
	Json(r, success, errorCode, errorMessage, showType, data...)
	r.Exit()
}

// JsonRedirect 返回标准JSON数据引导客户端跳转。
func JsonRedirect(r *ghttp.Request, success bool, errorCode int, errorMessage string, redirect string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonRes{
		Success:      success,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Data:         responseData,
		Redirect:     redirect,
		ShowType:     REDIRECT,
	})
}

// JsonRedirectExit 返回标准JSON数据引导客户端跳转，并退出当前HTTP执行函数。
func JsonRedirectExit(r *ghttp.Request, success bool, errorCode int, errorMessage string, redirect string, data ...interface{}) {
	JsonRedirect(r, success, errorCode, errorMessage, redirect, data...)
	r.Exit()
}
