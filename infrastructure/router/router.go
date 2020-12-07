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

type Controller interface {
	GetStudentsHandler(w http.ResponseWriter, r *http.Request)
	GetStudentURLHandler(w http.ResponseWriter, r *http.Request)
}

// router
func NewRouter(controller Controller) {
	router := mux.NewRouter()

	// GET students from csv
	router.HandleFunc(
		"/readcsv",
		controller.GetStudentsHandler,
	).Methods("GET")

	// Get csv from url
	router.HandleFunc(
		"/storedata",
		controller.GetStudentURLHandler,
	).Methods("GET")

	// Run server
	srv := server(router)
	fmt.Println("Server listen at " + config.C.GetServerAddr())
	log.Fatal("Fail router", srv.ListenAndServe())
}

func server(router *mux.Router) *http.Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.GetServerAddr(),
		WriteTimeout: config.C.Server.Timeout * time.Second,
		ReadTimeout:  config.C.Server.Timeout * time.Second,
	}
	return srv
}
