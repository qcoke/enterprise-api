package api

import (
	"enterprise-api/app/service"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// List
// @summary 用户列表接口
// @tags    用户服务
// @router  /user/list [GET]
// @success 200 {object} string "执行结果"
func (a *userApi) List(r *ghttp.Request) {
	res := service.User.Detail("1")
	r.Response.Write("hello", res)
}

// SignIn
// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *userApi) SignIn(r *ghttp.Request) {
//	var (
//		data *model.UserApiSignInReq
//	)
//	if err := r.Parse(&data); err != nil {
//		response.JsonExit(r, 1, err.Error())
//	}
//	if err := service.User.SignIn(r.Context(), data.Passport, data.Password); err != nil {
//		response.JsonExit(r, 1, err.Error())
//	} else {
//		response.JsonExit(r, 0, "ok")
//	}
//}
