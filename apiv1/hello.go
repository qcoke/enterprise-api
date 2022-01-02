package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}

type DetailReq struct {
	g.Meta `path:"/detail" tags:"Detail" method:"get" summary:"我是详情"`
}

type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
