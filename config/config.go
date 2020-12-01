package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

//Config structure
type config struct {
	Server struct {
		Address string
		Port int
		Timeout  time.Duration
	}
	CsvPath struct{
		Path string
	}
}

var C config

// Ready config file yml
func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src",  "golang-bootcamp-2020", "config"))
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