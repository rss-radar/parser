package user

import "github.com/rss-radar/parser/database"

func (u *User) Create() (err error) {
	err = database.DB.Create(&u).Error

	return
}
