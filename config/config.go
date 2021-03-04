package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Repository string
	}

	Server struct {
		Address string
	}
}

var C config

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