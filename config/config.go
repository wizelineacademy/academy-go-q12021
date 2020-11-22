package config

import (
	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		URL  string
		Port string
	}
}

// Config - variable containing server configuration
var C config

// ReadConfig - reads server configuration
func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
