package user

import "github.com/rss-radar/parser/database"

func FindByUsername(username string) (*User, error) {
	user := &User{}

	d := database.DB.Where("username = ?", username).First(&user)
	return user, d.Error
}
