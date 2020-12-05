package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang-bootcamp-2020/config"
)

type Controller interface {
	GetStudentsHandler(w http.ResponseWriter, r *http.Request)
	GetStudentUrlHandler(w http.ResponseWriter, r *http.Request)
}

// router
func NewRouter(controller Controller)() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	})

	// GET students from csv
	router.HandleFunc("/readcsv", controller.GetStudentsHandler).Methods("GET")

	// Get students from url
	router.HandleFunc("/storedata",controller.GetStudentUrlHandler).Methods("GET")

	// get server
	srv := server(router)
	//run server
	log.Fatal("Fail router",srv.ListenAndServe())
}


func server(router *mux.Router) *http.Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.Server.Address + ":" + strconv.Itoa(config.C.Server.Port),
		WriteTimeout: config.C.Server.Timeout * time.Second,
		ReadTimeout:  config.C.Server.Timeout * time.Second,
	}
	return srv
}
