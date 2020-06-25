package bootstrap

import (
	"github.com/jinzhu/gorm"
	"github.com/rss-radar/parser/database"
)

func SetupDB() *gorm.DB {
	db := database.InitDB()

	return db
}
