package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
    CSV_SOURCE      string `mapstructure:"CSV_SOURCE"`
}

var PokedexConfig = LoadConfig()

/* Reads configuration from pokemon.env file */
func LoadConfig() (config Config) {
    viper.AddConfigPath("./config/")
    viper.SetConfigName("pokemon")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}
