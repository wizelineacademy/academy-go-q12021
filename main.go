package main

import (
	"encoding/csv"
	"net/http"
	"os"

	"github.com/cesararredondow/academy-go-q12021/handlers"
	"github.com/cesararredondow/academy-go-q12021/routes"
	"github.com/cesararredondow/academy-go-q12021/services"
	"github.com/cesararredondow/academy-go-q12021/usecases"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCsvFile
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	filePath := goDotEnvVariable("filePath")
	PORT := goDotEnvVariable("PORT")
	LogLevel := goDotEnvVariable("lOGLEVEL")
	pokemonAPI := goDotEnvVariable("pokemonAPI")

	level, err := log.ParseLevel(LogLevel)
	if err != nil {
		log.Fatal("Failed creating logger: %w", err)
	}

	l := log.New()
	l.SetLevel(level)
	l.Out = os.Stdout
	l.Formatter = &log.JSONFormatter{}
	render := render.New()
	l.Info("starting the app")

	router := mux.NewRouter()
	rf, err := os.Open(filePath)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	wf, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingCsvFile)
	}
	defer rf.Close()
	defer wf.Close()

	csvw := csv.NewWriter(wf)

	s1, _ := services.New(rf, csvw, pokemonAPI, filePath)
	u1 := usecases.New(s1)
	h1 := handlers.New(u1, l, render)
	routes.New(h1, router)

	log.Info(http.ListenAndServe(":"+PORT, router))
}
