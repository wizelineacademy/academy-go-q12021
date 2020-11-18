package config

import "github.com/fallentemplar/goenv"

type AppConfig struct {
	PORT string
	DB   DBConfig
}

type DBConfig struct {
	NAME string
	HOST string
	PORT string
}

//C is the app global config object
var C AppConfig

//ReadConfig reads the env variables and saves its values
func (config *AppConfig) ReadConfig() {
	config.PORT = goenv.GetString("PORT", "8080")
	config.DB.HOST = goenv.GetString("DBHOST", "localhost")
	config.DB.PORT = goenv.GetString("DBPORT", "3306")
	config.DB.NAME = goenv.GetString("DBNAME", "srms")
}
