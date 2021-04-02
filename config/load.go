package configz

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/go-playground/validator/v10"
	"log"
)


func Load(configFile *string) (*Configuration, error){
	viperConf := viper.New()
	viperConf.SetConfigType("yaml")
	viperConf.AddConfigPath(".")
	viperConf.SetConfigFile(*configFile)
	if err := viperConf.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Configuration{}
	err := viperConf.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if err = validate.Struct(config) ; err != nil {
		return nil, err
	}
	viperConf.WatchConfig()
	viperConf.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file has changed: ", e.Name)
		if err := viperConf.Unmarshal(config); err != nil {
			log.Println("Error updating the config...", "err", err)
		}
	})
	return config, nil
}