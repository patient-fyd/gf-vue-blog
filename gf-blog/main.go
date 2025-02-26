package main

import (
	_ "gf-blog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-blog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
