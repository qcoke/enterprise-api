package handler

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/template-single/apiv1"
)

var Hello = handlerHello{}

type handlerHello struct{}

func (a *handlerHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello")
	return
}

func (a *handlerHello) Detail(ctx context.Context, req *apiv1.DetailReq) (res *apiv1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Detail")
	return
}
