// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleDao is the data access object for table article.
type ArticleDao struct {
	Table   string         // Table is the underlying table name of the DAO.
	Group   string         // Group is the database configuration group name of current DAO.
	Columns ArticleColumns // Columns contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for table article.
type ArticleColumns struct {
	Id        string // ID
	Local     string // zh 中文 en 英文
	Title     string // 标题
	Body      string // 文章内容
	TypeId    string // 文章分类ID
	Author    string // 作者
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

//  articleColumns holds the columns for table article.
var articleColumns = ArticleColumns{
	Id:        "id",
	Local:     "local",
	Title:     "title",
	Body:      "body",
	TypeId:    "type_id",
	Author:    "author",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		Group:   "default",
		Table:   "article",
		Columns: articleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}