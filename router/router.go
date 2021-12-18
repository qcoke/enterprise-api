package router

import (
	"enterprise-api/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/v1", func(group *ghttp.RouterGroup) {
		group.ALL("/user", api.User)
	})
}
