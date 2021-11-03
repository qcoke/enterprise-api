package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns userColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserColumns defines and stores column names for table user.
type userColumns struct {
	Id       string // 用户ID
	Passport string // 用户账号
	Password string // 用户密码
	Nickname string // 用户昵称
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

var (
	// User is globally public accessible object for table user operations.
	User = UserDao{
		M:     g.DB("default").Model("user").Safe(),
		DB:    g.DB("default"),
		Table: "user",
		Columns: userColumns{
			Id:       "id",
			Passport: "passport",
			Password: "password",
			Nickname: "nickname",
			CreateAt: "create_at",
			UpdateAt: "update_at",
		},
	}
)
