// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEmployee = "employees"

// Employee mapped from table <employees>
type Employee struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	LastName    string     `gorm:"column:last_name;not null" json:"last_name"`
	Phone1      *string    `gorm:"column:phone1" json:"phone1"`
	Phone2      *string    `gorm:"column:phone2" json:"phone2"`
	Email1      *string    `gorm:"column:email1" json:"email1"`
	Email2      *string    `gorm:"column:email2" json:"email2"`
	Role        *string    `gorm:"column:role" json:"role"`
	DateAdded   *time.Time `gorm:"column:date_added" json:"date_added"`
	UserAddedID *int32     `gorm:"column:user_added_id" json:"user_added_id"`
}

// TableName Employee's table name
func (*Employee) TableName() string {
	return TableNameEmployee
}
