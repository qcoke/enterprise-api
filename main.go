package main

import (
	_ "enterprise-api/boot"
	_ "enterprise-api/router"
	"github.com/gogf/swagger"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.SetPort(8080)
	s.Run()
}
