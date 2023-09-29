// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package biz

import (
	"context"
	"letga/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	localBizCtx IBizCtx
)

type (
	IBizCtx interface {
		Init(r *ghttp.Request, customCtx *Context)
		Get(ctx context.Context) *Context
		SetUser(ctx context.Context, ctxUser *ContextUser)
		SetData(ctx context.Context, data g.Map)
	}
)

type sBizCtx struct{}

type Context struct {
	User *ContextUser // User in context.
	Data g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

type ContextUser struct {
	Uuid string // 唯一ID
	Role string // 角色集
}

// 注册业务上下文
func RegisterCtx() {
	localBizCtx = &sBizCtx{}
}

// 获取业务上下文单例
func Ctx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sBizCtx) Get(ctx context.Context) *Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *ContextUser) {
	s.Get(ctx).User = ctxUser
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
