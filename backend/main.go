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
	log.Println("Turning server on ...")

	var configFile string
	flag.StringVar(
		&configFile,
		"public-config-file",
		"config.yml",
		"Path to public config file",
	)
	flag.Parse()

	log.Println("Reading config file...")
	// Read config file
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	log.Println("Generating File Reader ...")
	rf, err := os.Open(cfg.DB)
	if err != nil {
		log.Fatal("Failed open File Reader: %w", err)
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}

	log.Println("Generating File Writter ...")
	wf, err := os.OpenFile(cfg.DB, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("Failed open File Writter: %w", err)
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}
	log.Println("File Reader and Writter generated succesfully")

	defer rf.Close()
	defer wf.Close()

	csvw := csv.NewWriter(wf)

	service, err := service.New(rf, csvw, cfg.DB)
	if err != nil {
		log.Fatal("Failed running service : %w", err)
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}

	useCase := usecase.New(service)
	controller := controller.New(useCase, render.New())
	router := router.New(controller)

	// Start server
	fmt.Printf("Server running at port [%s].", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, router))
}
