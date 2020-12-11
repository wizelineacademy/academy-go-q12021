/**
Router Mux
*/
package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang-bootcamp-2020/config"

	"github.com/gorilla/mux"
)

// Controller interfaces
type Controller interface {
	ReadStudentsHandler(w http.ResponseWriter, r *http.Request)
	StoreStudentURLHandler(w http.ResponseWriter, r *http.Request)
}

// NewRouter new mux router
func NewRouter(controller Controller) {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	// GET students from csv
	apiRouter.PathPrefix("/readcsv").HandlerFunc(controller.ReadStudentsHandler).Methods("GET")
	//router.HandleFunc(
	//	"/readcsv",
	//	controller.ReadStudentsHandler,
	//).Methods("GET")

	// Get csv from url
	//router.HandleFunc(
	//	"/storedata",
	//	controller.StoreStudentURLHandler,
	//).Methods("GET")
	apiRouter.PathPrefix("/storedata").HandlerFunc(controller.StoreStudentURLHandler).Methods("GET")
	// Run server
	srv := server(apiRouter)
	fmt.Println("Server listen at " + config.C.GetServerAddr())
	log.Fatal("Fail router", srv.ListenAndServe())
}

// server obtain server setup
func server(router *mux.Router) *http.Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.GetServerAddr(),
		WriteTimeout: config.C.Server.Timeout * time.Second,
		ReadTimeout:  config.C.Server.Timeout * time.Second,
	}
	return srv
}
