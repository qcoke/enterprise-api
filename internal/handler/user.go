package handler

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/template-single/apiv1"
)

// User 对外暴露的对象（大写字母开头）
var User = handlerUser{}

type handlerUser struct{}

func (u *handlerUser) List(ctx context.Context, req *apiv1.UserListReq) (res *apiv1.UserListRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteJsonExit(res)
	return
}
