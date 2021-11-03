package internal

import "github.com/gogf/gf/os/gtime"

type User struct {
	Id       uint        `orm:"id,primary" json:"id"`
	Passport string      `orm:"passport" json:"passport"`
	Password string      `orm:"password"   json:"password"` // 用户密码
	Nickname string      `orm:"nickname"   json:"nickname"` // 用户昵称
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
}
