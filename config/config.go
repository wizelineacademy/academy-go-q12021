package config

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Config returns the information required by the application to run
type Config struct {
	InfoLog         *log.Logger
	ErrorLog        *log.Logger
	Addr            *string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// Init the application's configuration
func Init(infoLog, errorLog *log.Logger) (*Config, error) {

	const (
		readTimeout     = 5 * time.Second   //Seconds
		writeTimeout    = 10 * time.Second  //Seconds
		idleTimeout     = 120 * time.Second //Seconds
		shutdownTimeout = 30 * time.Second  //Seconds
	)

	// Read config file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Get the server address/port
	addr := viper.GetString("addr")

	// Retrun the config
	return &Config{infoLog, errorLog, &addr, readTimeout, writeTimeout, idleTimeout, shutdownTimeout}, nil
}
