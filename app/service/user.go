package service

import (
	"enterprise-api/app/dao"
	"enterprise-api/app/model"

	"github.com/gogf/gf/database/gdb"
)

// User 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) List(pageNo uint, pageSize uint) gdb.Result {
	results, err := dao.User.FindAll()
	if err != nil {
		results = nil
	}
	return results
}

func (s *userService) Detail(userId string) gdb.Record {
	results, err := dao.User.Where("Id = ?", userId).One()
	if err != nil {
		return results
	}
	return nil
}

func (s *userService) AddOrUpdate(u *model.User) error {
	uId := string(u.Id)
	if uId == "" {
		// Add
		if _, err := dao.User.Save(u); err != nil {
			return err
		}
	} else {
		// Update
		if _, err := dao.User.Update(u); err != nil {
			return err
		}
	}
	return nil
}
