package main

import (
	_ "letga/boot"

	_ "letga/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"letga/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
