package config

import "github.com/spf13/viper"

type appConfig struct {
	Name    string
	RunMode string
	Addr    string
}

func newAppConfig() *appConfig {
	viper.SetDefault("app.name", "parser")
	viper.SetDefault("app.run_mode", "release")
	viper.SetDefault("app.addr", ":8080")

	return &appConfig{
		Name:    viper.GetString("app.name"),
		RunMode: viper.GetString("app.run_mode"),
		Addr:    viper.GetString("app.addr"),
	}
}
