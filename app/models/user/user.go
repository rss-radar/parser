package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"column:email;type:varchar(255);unique"`
}
