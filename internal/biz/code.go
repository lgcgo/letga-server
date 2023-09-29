// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package biz

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
)

type BizCode struct {
	code    int
	message string
	detail  BizCodeDetail
}
type BizCodeDetail struct {
	ShowType  int
	ErrorCode int
}

var (
	CodeOk                 = NewCode(200, 0, "OK")
	CodeValidationFailed   = NewCode(200, 10001, "%s")
	CodeNotAuthorized      = NewCode(401, 20102, "Not Authorized")
	CodeNotFound           = NewCode(404, 40400, "Not Found")
	CodeInternal           = NewCode(500, 50001, "Internal Error")      // 系统错误
	CodeServiceUnavailable = NewCode(500, 50001, "Service Unavailable") // 维护状态
	// ...
)

func (c BizCode) BizDetail() BizCodeDetail {
	return c.detail
}

func (c BizCode) Code() int {
	return c.code
}

func (c BizCode) Message() string {
	return c.message
}

func (c BizCode) Detail() interface{} {
	return c.detail
}

func NewCode(httpCode int, errorCode int, message string, errorShowType ...int) gcode.Code {
	var showType int = 0
	if len(errorShowType) > 0 {
		showType = errorShowType[0]
	}
	return BizCode{
		code:    httpCode,
		message: message,
		detail: BizCodeDetail{
			ShowType:  showType,
			ErrorCode: errorCode,
		},
	}
}

// 业务错误码
func WithCode(code gcode.Code, errorCode int, message string, errorShowType ...int) gcode.Code {
	var showType int = 0
	if len(errorShowType) > 0 {
		showType = errorShowType[0]
	}
	return BizCode{
		code:    code.Code(),
		message: message,
		detail: BizCodeDetail{
			ShowType:  showType,
			ErrorCode: errorCode,
		},
	}
}

func WithCodef(code gcode.Code, formatInfo ...interface{}) gcode.Code {
	var (
		bizCode gcode.Code
		ok      bool
	)
	if bizCode, ok = code.(BizCode); !ok {
		bizCode = CodeInternal
	}
	var (
		errorCode = bizCode.(BizCode).BizDetail().ErrorCode
		message   = fmt.Sprintf(code.Message(), formatInfo...)
		showType  = bizCode.(BizCode).BizDetail().ShowType
	)
	return WithCode(code, errorCode, message, showType)
}
