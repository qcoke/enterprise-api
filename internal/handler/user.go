package handler

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/template-single/apiv1"
	"github.com/gogf/template-single/internal/service"
)

// User 对外暴露的对象（大写字母开头）
var User = handlerUser{}

type handlerUser struct{}

func (u *handlerUser) List(ctx context.Context, req *apiv1.UserListReq) (res *apiv1.UtilRes, err error) {
	userList, err := service.User.List(ctx, 10, 1)
	g.RequestFromCtx(ctx).Response.WriteJsonExit(userList)
	return
}
