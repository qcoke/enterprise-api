package service

import (
	"context"
	"enterprise-api/app/model"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文管理服务
var Context = contextService{}

type contextService struct{}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// Get 获取上下文变量，如果没有设置，那么返回 nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 设置用户数据
func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
