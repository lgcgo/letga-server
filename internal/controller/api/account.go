package api

import (
	"context"
	v1 "letga/api/api/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAccount struct{}

var Account = &cAccount{}

// 获取账户信息
func (c *cAccount) Info(ctx context.Context, req *v1.AccountInfoReq) (res *v1.AccountInfoRes, err error) {
	var (
		out *model.User
	)
	if out, err = service.User().GetCurrentUser(ctx); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}
