package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255);unique"`
}
