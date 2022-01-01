// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `orm:"id,primary" json:"id"`        // Id
	UserName  string      `orm:"user_name"  json:"userName"`  // 用户名
	PassWord  string      `orm:"pass_word"  json:"passWord"`  // 密码
	Email     string      `orm:"email"      json:"email"`     // 邮箱
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 注册时间
}

// FriendLink is the golang structure for table friend_link.
type FriendLink struct {
	Id          int    `orm:"id,primary"  json:"id"`          // Id
	Title       string `orm:"title"       json:"title"`       // 标题
	Description string `orm:"description" json:"description"` // 描述
	Url         string `orm:"url"         json:"url"`         // 链接地址
	Type        string `orm:"type"        json:"type"`        // 跳转类型
}

// Article is the golang structure for table article.
type Article struct {
	Id        int         `orm:"id,primary" json:"id"`        // ID
	Local     string      `orm:"local"      json:"local"`     // zh 中文 en 英文
	Title     string      `orm:"title"      json:"title"`     // 标题
	Body      string      `orm:"body"       json:"body"`      // 文章内容
	TypeId    int         `orm:"type_id"    json:"typeId"`    // 文章分类ID
	Author    string      `orm:"author"     json:"author"`    // 作者
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
}
