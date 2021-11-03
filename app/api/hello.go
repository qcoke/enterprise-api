package api

import (
	"github.com/gogf/gf"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// UserApiSignInReq 登录请求参数，用于前后端交互参数格式约定
type UserApiSignInReq struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Writeln("GF 版本：", gf.VERSION)
}
