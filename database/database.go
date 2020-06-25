package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rss-radar/parser/config"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database %v", err))
	}

	DB = db

	return db
}
