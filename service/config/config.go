package config

import (
	"log"

	"github.com/etyberick/golang-bootcamp-2020/entity"

	"github.com/spf13/viper"
)

// Runtime contains all the global runtime configuration
type runtime struct {
	Info entity.Config
}

// Load configuration from the configuration file
func (r *runtime) load() {
	// Set default values
	viper.SetDefault("database", "./quotes.csv")
	viper.SetDefault("api_port", "8080")
	// Set config location values
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Open the file
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		log.Println("warning: loading default config values")
	}
	// Parse values
	r.Info.CSVFilepath = viper.GetString("database")
	r.Info.Port = viper.GetString("api_port")
	log.Println(r.Info)
}

// New entity.Config instance with initialized values
func New() entity.Config {
	r := &runtime{}
	r.load()
	return r.Info
}
