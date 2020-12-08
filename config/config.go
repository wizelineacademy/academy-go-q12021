/**
Get config environment settings
*/
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// config structure
type config struct {
	Server struct {
		Address string
		Port    int
		Timeout time.Duration
	}
	CsvPath struct {
		Path string
	}
}

// C global var type config
var C config

// ReadConfig read  yml file
func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "golang-bootcamp-2020", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetServerAddr obtain the full server address in a string
func (c config) GetServerAddr() string {
	return C.Server.Address + ":" + strconv.Itoa(C.Server.Port)
}
