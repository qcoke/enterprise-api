package model

import (
	"github.com/gogf/gf/net/ghttp"
)

const (
	ContextKey = "requestContent"
)

type Context struct {
	Session *ghttp.Session // 当前的SESSION管理对象
	User    *ContextUser   // 上下文用户信息
}

type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	NickName string // 用户昵称
}
