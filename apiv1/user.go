package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这个文件统一用来定义服务器的返回值

type UserListReq struct {
	g.Meta `path:"/list" tags:"User" method:"get" summary:"获取用户列表"`
}

type UserListRes struct {
	result bool
}
