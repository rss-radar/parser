package bootstrap

import (
	"github.com/jinzhu/gorm"
	userModel "github.com/rss-radar/parser/app/models/user"
	"github.com/rss-radar/parser/database"
)

func SetupDB() *gorm.DB {
	db := database.InitDB()

	db.AutoMigrate(&userModel.User{})

	return db
}
