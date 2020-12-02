package main

import (
	"github.com/etyberick/golang-bootcamp-2020/integration/router"
	"github.com/etyberick/golang-bootcamp-2020/service/config"
)

func main() {
	//Load Config
	c := config.New()
	//Initialize API
	router.Initialize(c)
}
