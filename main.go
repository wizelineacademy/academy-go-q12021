package main

import (
	"academy/router"
	"academy/services/dataload"
)

func main() {
	dataload.LoadData()
	router.InitServer()
}
