package main

import (
	"testing"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/infrastructure/services"
)

func TestMain(m *testing.M) {
	config.ReadConfig()
	s := services.NewClient()
	if s!=nil{
		m.Run()

	}
}