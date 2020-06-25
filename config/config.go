package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	RunModeDebug   = "debug"
	RunModeRelease = "release"
	RunModeTest    = "test"

	configFilePath = "./application.yaml"
	configFileType = "yaml"
)

var (
	AppConfig *appConfig
	DBConfig  *dbConfig
)

func InitConfig() {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read config file %v", err))
	}

	AppConfig = newAppConfig()
	DBConfig = newDBConfig()
}
