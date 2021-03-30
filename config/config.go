package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"

	"github.com/spf13/viper"
)

// config struct for config settings
type config struct {
	Database struct {
		Repository string
	}

	Server struct {
		Address string
	}
}

// C is exported variable for the entire project/app
var C config

// ReadConfig reads configuration settings from yaml file.
func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		log.Fatalln(1)
	}

	spew.Dump(C)
}