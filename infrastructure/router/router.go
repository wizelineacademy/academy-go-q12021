// router package
package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang-bootcamp-2020/config"

	"github.com/gorilla/mux"
)

// Controller interface
type Controller interface {
	ReadStudentsHandler(w http.ResponseWriter, r *http.Request)
	StoreStudentURLHandler(w http.ResponseWriter, r *http.Request)
}

// NewRouter mux router
func NewRouter(controller Controller) {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	// endpoint to Get students from URL
	apiRouter.PathPrefix("/storedata").
		HandlerFunc(controller.StoreStudentURLHandler).
		Methods("GET").
		Name("storedata")

	// endpoint to GET students from csv
	apiRouter.PathPrefix("/readcsv").
		HandlerFunc(controller.ReadStudentsHandler).
		Methods("GET").
		Name("readcsv")

	// Run GetServer
	server := GetServer(apiRouter)
	fmt.Println("Server listen at " + config.C.GetServerAddr())
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// GetServer obtain GetServer setup
func GetServer(r *mux.Router) *http.Server {
	srv := &http.Server{
		Handler:      r,
		Addr:         config.C.GetServerAddr(),
		WriteTimeout: config.C.Server.Timeout * time.Second,
		ReadTimeout:  config.C.Server.Timeout * time.Second,
	}
	return srv
}
