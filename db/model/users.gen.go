// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID       int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string  `gorm:"column:name;not null" json:"name"`
	LastName *string `gorm:"column:last_name" json:"last_name"`
	Username string  `gorm:"column:username;not null" json:"username"`
	Password string  `gorm:"column:password;not null" json:"password"`
	Email    string  `gorm:"column:email;not null" json:"email"`
	Phone    *string `gorm:"column:phone" json:"phone"`
	Role     string  `gorm:"column:role;not null" json:"role"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
