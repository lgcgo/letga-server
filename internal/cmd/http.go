package cmd

import (
	"context"
	"letga/internal/controller/admin"
	"letga/internal/service"

	"letga/internal/controller/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Http = gcmd.Command{
		Name:  "http",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			glog.Infof(ctx, "Start http server")

			s := g.Server()
			// 获取静态目录
			staticRoot := g.Cfg().MustGet(ctx, "server.staticRoot").String()
			if staticRoot == "" {
				g.Log().Fatal(ctx, "静态服务根目录配置不能为空")
			}
			uploadDir := g.Cfg().MustGet(ctx, "upload.baseDir").String()
			if uploadDir == "" {
				g.Log().Fatal(ctx, "上传路径配置不能为空")
			}
			staticPath := staticRoot + uploadDir
			s.AddStaticPath(uploadDir, staticPath)
			s.Group(uploadDir, func(group *ghttp.RouterGroup) {
				// 路由控制
			})
			// 后台服务模块
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
				// 这里添加无需权限验证的路由
				)
				group.Middleware(service.Middleware().Authentication)
				group.Bind(
					admin.AuthRole,
					admin.AuthRoleAccess,
					admin.AuthRoute,
					admin.AuthAccess,
					admin.Media,
					admin.Menu,
					admin.User,
				)
			})

			// API服务模块
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					api.AccountSign,
				)
				group.Middleware(service.Middleware().Authentication)
				group.Bind(
					api.Account,
				)
			})
			s.Run()
			// sApi.Start()
			// g.Wait()
			return nil
		},
	}
)
