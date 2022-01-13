package service

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/template-single/internal/service/internal/dao"
	"golang.org/x/net/context"
)

var User = userService{}

type userService struct{}

// List 查询列表
func (u *userService) List(ctx context.Context, pageSize uint, pageNo uint) (gdb.Result, error) {
	list, err := dao.User.Ctx(ctx).All()
	if err == nil {
		return list, nil
	}
	return nil, err
}
