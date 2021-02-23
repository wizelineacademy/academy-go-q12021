package utils

import (
	"github.com/spf13/viper"
	"log"
)

func GetEnvVar(key string) string {
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

  value, ok := viper.Get(key).(string)

  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  return value
}