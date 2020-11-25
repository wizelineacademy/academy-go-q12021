package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang-bootcamp-2020/interface/controller"
)

// router
func Newrouter() {

	router := mux.NewRouter()

	// route /
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	})

	// GET students from csv
	router.HandleFunc("/showcsv", GetStudents).Methods("GET")

	// Get students from MongoDB
	router.HandleFunc("/download", DownloadDb).Methods("GET")

	//http.Handle("/", router)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

// GetStudents list
func GetStudents(writer http.ResponseWriter, request *http.Request) {
	students := controller.GetStudents()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(students)
}


func DownloadDb(writer http.ResponseWriter, request *http.Request) {

	json.NewEncoder(writer).Encode(map[string]bool{
		"ok": true,
		"downloaded":true,
	})
}