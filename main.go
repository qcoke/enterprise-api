package main

import (
	_ "github.com/gogf/template-single/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/template-single/internal/cmd"
)

func main() {
	var (
		ctx = gctx.New()
	)
	if err := cmd.Main.Run(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
