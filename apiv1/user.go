package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"/list" tags:"User" method:"get" summary:"获取用户列表"`
}

type UserAddOrUpdateReq struct {
	g.Meta `path:"/add_or_update" tags:"User" method:"get" summary:"增加或更新用户"`
}

type UserDeleteReq struct {
	g.Meta `path:"/del" tags:"User" method:"get" summary:"删除用户"`
}
