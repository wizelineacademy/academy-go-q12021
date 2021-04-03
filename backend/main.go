package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"main/config"
	"main/controller"
	"main/router"
	"main/service"
	"main/usecase"

	"github.com/unrolled/render"
)

// Abnormal exit constants
const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCSVFile
)

func main() {
	log.Println("Will start")

	var configFile string
	flag.StringVar(
		&configFile,
		"public-config-file",
		"config.yml",
		"Path to public config file",
	)
	flag.Parse()

	log.Println("Read config file")
	// Read config file
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}
	
	log.Println("Will open database file")
	rf, err := os.Open(cfg.DB)
	if err != nil {
		log.Println("Error")
		log.Fatal(err.Error())
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}

	wf, err := os.OpenFile(cfg.DB, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}
	defer rf.Close()
	defer wf.Close()

	csvw := csv.NewWriter(wf)

	s, _ := service.New(rf, csvw)
	u := usecase.New(s)
	c := controller.New(u, render.New())
	r := router.New(c)

	// Start server
	fmt.Println("Starting server at port [8080].")
	log.Fatal(http.ListenAndServe(":8080", r))
}
