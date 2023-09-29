package boot

import (
	_ "letga/internal/logic"
	_ "letga/internal/packed"
	"letga/utility/casbin"
	"letga/utility/hashid"
	"letga/utility/token"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.New()

// 初始公共组件
func init() {
	initCasbin()
	initToken()
	initHashid()
}

// 初始化Casbin组件
func initCasbin() {
	casbin.Init(&casbin.Config{
		PolicyFilePath: g.Cfg().MustGet(ctx, "casbin.policyFilePath").String(),
	})
}

// 初始化Token组件
func initToken() {
	var (
		tokenAudienceString      = g.Cfg().MustGet(ctx, "token.audience").String()
		tokenAudienceStringArray = strings.Split(tokenAudienceString, ",")
	)

	token.Init(&token.Config{
		Issuer:   g.Cfg().MustGet(ctx, "token.issuer").String(),
		Audience: tokenAudienceStringArray,
		SignKey:  g.Cfg().MustGet(ctx, "token.signKey").Bytes(),
	})
}

// 初始Hashid组件
func initHashid() {
	hashid.Init(&hashid.Config{
		Salt:      g.Cfg().MustGet(ctx, "hashid.key.salt").String(),
		MinLength: g.Cfg().MustGet(ctx, "hashid.key.minLength").Int(),
	})
}
