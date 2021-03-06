// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MigrationsForDao is the golang structure of table migrations for DAO operations like Where/Data.
type Migrations struct {
	g.Meta             `orm:"dto:true"`
	IdMigration        interface{} // surrogate key
	Name               interface{} // migration name, unique
	CreatedAt          *gtime.Time // date migrated or rolled back
	Statements         interface{} // SQL statements for this migration
	RollbackStatements interface{} // SQL statment for rolling back migration
	Status             interface{} // update indicates it is a normal migration while rollback means this migration is rolled back
}
