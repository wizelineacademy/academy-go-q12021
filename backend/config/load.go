package config

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var defaultTimeout = 3 * time.Second

// Load - pulls the config data from the config file
func Load(configFile string) (*Configuration, error) {
	public := viper.New()
	public.SetConfigFile(configFile)
	if err := public.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Configuration{}
	err := public.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	validator := validator.New()
	if err = validator.Struct(config); err != nil {
		return nil, err
	}

	public.WatchConfig()
	public.OnConfigChange(func(in fsnotify.Event) {
		if err := public.Unmarshal(config); err != nil {
			log.Println("failed to update public config after hot reload", "err", err)
		}
	})

	return config, nil
}
