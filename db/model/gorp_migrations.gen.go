// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGorpMigration = "gorp_migrations"

// GorpMigration mapped from table <gorp_migrations>
type GorpMigration struct {
	ID        string     `gorm:"column:id;primaryKey" json:"id"`
	AppliedAt *time.Time `gorm:"column:applied_at" json:"applied_at"`
}

// TableName GorpMigration's table name
func (*GorpMigration) TableName() string {
	return TableNameGorpMigration
}
