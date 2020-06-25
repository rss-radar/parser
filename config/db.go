package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Connection string
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string

	URL   string
	Debug bool
}

func newDBConfig() *dbConfig {
	viper.SetDefault("db.connection", "postgres")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.database", viper.GetString("app.name"))
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "")

	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")

	database := viper.GetString("db.database")
	database = database + "_" + AppConfig.RunMode

	url := createDatabaseURL(username, password, host, port, database)

	return &dbConfig{
		Connection: viper.GetString("db.connection"),
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   username,
		Password:   password,
		URL:        url,
	}
}

func createDatabaseURL(username string, password string, host string, port int, database string) string {
	return fmt.Sprintf("sslmode=disable host=%s port=%d dbname=%s user=%s password=%s",
		host,
		port,
		database,
		username,
		password,
	)
}
