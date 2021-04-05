package config

import (
	"fmt"
)

// Configuration contains application configuration
type Configuration struct {
	AppName  string `mapstructure:"app_name" validate:"required"`
	HTTPPort string `mapstructure:"http_port" validate:"required"`
	DB       string `mapstructure:"db" validate:"required"`
}

// BindAddress generates address with listening port
func (app *Configuration) BindAddress() string {
	return fmt.Sprintf("0.0.0.0:%s", app.HTTPPort)
}
