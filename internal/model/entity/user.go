// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        ` // Id
	UserName  string      `json:"userName"  ` // 用户名
	PassWord  string      `json:"passWord"  ` // 密码
	Email     string      `json:"email"     ` // 邮箱
	CreatedAt *gtime.Time `json:"createdAt" ` // 注册时间
}
