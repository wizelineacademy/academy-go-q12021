package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api-booking-time/config"
	"api-booking-time/infrastructure/datastore"
)

func main() {
	config.ReadConfig()
	//db := datastore.OpenDb()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api booking-time")
	})
	api.HandleFunc("/centres", getMenuView).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":" + config.Settings.Server.Port, r))
}
