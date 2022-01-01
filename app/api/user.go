package api

import (
	"enterprise-api/app/service"
	"enterprise-api/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// Add
// @summary 添加用户
// @tags    用户服务
// @router  /v1/user/add [POST]
// @success 200 {object} string "执行结果"
func (a *userApi) Add(r *ghttp.Request) {
	res := service.User.List(1, 10)
	response.JsonExit(r, 200, "查询成功", res)
}

// Update
// @summary 更新用户
// @tags    用户服务
// @router  /v1/user/list [POST]
// @success 200 {object} string "执行结果"
func (a *userApi) Update(r *ghttp.Request) {
	res := service.User.List(1, 10)
	response.JsonExit(r, 200, "查询成功", res)
}

// Delete
// @summary 删除用户
// @tags    用户服务
// @router  /v1/user/delete [POST]
// @success 200 {object} string "执行结果"
func (a *userApi) Delete(r *ghttp.Request) {
	res := service.User.List(1, 10)
	response.JsonExit(r, 200, "查询成功", res)
}

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

// Login
// @summary 用户登录
// @tags    用户服务
// @router  /v1/user/login [POST]
// @success 200 {object} string "执行结果"
func (a *userApi) Login(r *ghttp.Request) {
	response.JsonExit(r, 200, "登录成功", nil)
}
