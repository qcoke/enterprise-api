package service

import (
	"enterprise-api/app/dao"

	"github.com/gogf/gf/database/gdb"
)

// User 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) Detail(userId string) gdb.Record {
	results, err := dao.User.Where("Id = ?", userId).One()
	if err != nil {
		return results
	}
	return nil
}
