package main

import (
	"github.com/etyberick/golang-bootcamp-2020/entity"
	"github.com/etyberick/golang-bootcamp-2020/integration/router"
)

func main() {
	//Load Config
	config := entity.Config{
		Port:        "8080",
		CSVFilepath: "quotes.csv",
	}

	//Initialize API
	router.Initialize(config)
}
