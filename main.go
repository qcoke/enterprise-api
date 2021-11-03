package main

import (
	_ "enterprise-api/boot"
	_ "enterprise-api/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
