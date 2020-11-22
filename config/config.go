package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

// config structure allow me to store Database information as it was using MySQL for storage of the API
type config struct {
	Server struct {
		Address string
	}
	Sqlitedb struct {
		DBPath string
	}
	Sources struct {
		DigimonAPI string
	}
	Dest struct {
		DigimonCSV string
	}
}

// C an instance of the config structure
var C config

// ReadConfig loads configuration from config.yml
func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "MiguelAGrover", "golang-bootcamp-2020", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
