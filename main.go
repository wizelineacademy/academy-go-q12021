package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	writeTimeout = 15 * time.Second
	readTimeout  = 15 * time.Second
	idleTimeout  = 60 * time.Second
)

func main() {
	// Environment variables
	mongoDB := os.Getenv("MONGO_DB")
	mongoURI := os.Getenv("MONGO_STRING")
	csvFile := os.Getenv("CSV_FILE")
	addr := os.Getenv("ADDR")

	// Setup logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	// Create app
	convoy := New(
		AppConfig{
			MongoDB:  mongoDB,
			MongoURI: mongoURI,
			CSVFile:  csvFile,
		}, logger,
	)

	// Get router
	router := convoy.Router()

	// Create a Server instance with the router
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      router,
	}

	// Start the server
	logger.Fatal(srv.ListenAndServe())
}
