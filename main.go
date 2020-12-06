package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
	"github.com/wizelineacademy/golang-bootcamp-2020/config"
)

func main() {

	const (
		readTimeout     = 5 * time.Second   //Seconds
		writeTimeout    = 10 * time.Second  //Seconds
		idleTimeout     = 120 * time.Second //Seconds
		shutdownTimeout = 30 * time.Second  //Seconds
	)

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	//Init Config
	config, err := config.Init(infoLog, errorLog)
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			errorLog.Fatalf("message: fatal error config file not found: %s", err)

		} else {
			// Config file was found but another error was produced
			errorLog.Fatalf("message: fatal error initializing config: %s", err)
		}
	}
	defer config.DB.Close()

	sm := config.InitRoutes()
	// HTTP Server configuration
	s := &http.Server{
		Addr:         *config.Addr, //set the bind address
		Handler:      sm,           //set the default handler
		ReadTimeout:  readTimeout,  //max time to read request from the client
		WriteTimeout: writeTimeout, //max time to write request from the client
		IdleTimeout:  idleTimeout,  //max time for connections using TCP Keep-Alive
		ErrorLog:     errorLog,     //set the logger for the server
	}

	//Console message
	log.Printf("Starting server on %s", *config.Addr)
	log.Printf("Go to server http://localhost%s", *config.Addr)
	// Anon function to start the HTTP Server.
	go func() {
		// It returns an error in case of any failure.
		err := s.ListenAndServe()
		if err != nil {
			errorLog.Fatalf("msg: error initializing HTTP Server, err: %v\n", err)
		}
	}()

	// Listen for os signals to terminate the server.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// If the server was interrupted, print the interruption signal that stopped it.
	sig := <-sigChan
	infoLog.Println("Received terminate, graceful shutdown", sig)

	// Wait until the server has finished processing pending jobs, then terminate. Or, if the time threshold is met, terminate.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	shutdownErr := s.Shutdown(ctx)
	if err != nil {
		errorLog.Fatalf("msg: error shutting down the HTTP Server, err: %v\n", shutdownErr)
	}
}
