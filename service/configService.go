package service

import "github.com/spf13/viper"

// ConfigService intrface of service
type ConfigService interface {
	GetConfig() ConfigImpl
}

// ConfigImpl is the implementation for ConfigServer interface
type ConfigImpl struct {
	CSV    Csv
	Server Server
	Log    Log
	External External
}

// Server holds server config vars
type Server struct {
	Port int
}
// Log holds logs config vars
type Log struct {
	Level string
}

// External holds external config vars
type External struct {
	ApiUrl string
}

// Csv holds data for csv config
type Csv struct {
	FileName string
}

var config ConfigImpl

// GetConfig return a reference to the ConfigImpl configuration struct
func (c ConfigImpl) GetConfig() ConfigImpl {
	if c == (ConfigImpl{}) {
		c = loadConfig()
	}
	return c
}

func loadConfig() ConfigImpl {
	config := ConfigImpl{}
	viper.Unmarshal(&config)
	return config
}
