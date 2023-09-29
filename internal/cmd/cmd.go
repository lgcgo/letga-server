package cmd

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
)

var (
	// 同步处理器
	wg = sync.WaitGroup{}
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			cmds := []*gcmd.Command{&Http}
			for _, v := range cmds {
				var cmd = v
				wg.Add(1)
				err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
					cmd.Func(ctx, parser)
					wg.Done()
				}, func(ctx context.Context, err error) {
					glog.Infof(ctx, "Command err: %s", err.Error())
				})
				if err != nil {
					glog.Infof(ctx, "Command err: %s", err.Error())
				}
			}
			wg.Wait()
			return nil
		},
	}
)
