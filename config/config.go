package config

import (
	"encoding/json"
	"log"
	"os"
)

type DataSources struct {
	CSV  string `json:"csv"`
	JSON string `json:"json"`
}

type Config struct {
	DataSources DataSources `json:"dataSource"`
	Env         string      `json:"env"`
	Port        int         `json:"port"`
}

func New(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config, nil
}
