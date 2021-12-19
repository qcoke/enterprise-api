package api

import (
	"enterprise-api/app/service"
	"enterprise-api/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// List
// @summary 用户列表接口
// @tags    用户服务
// @router  /v1/user/list [GET]
// @success 200 {object} string "执行结果"
func (a *userApi) List(r *ghttp.Request) {
	res := service.User.List(1, 10)
	response.JsonExit(r, 200, "查询成功", res)
}

// Detail
// @summary 用户详情接口
// @tags    用户服务
// @router  /v1/user/detail [GET]
// @success 200 {object} string "执行结果"
func (a *userApi) Detail(r *ghttp.Request) {
	res := service.User.List(1, 10)
	response.JsonExit(r, 200, "查询成功", res)
}
