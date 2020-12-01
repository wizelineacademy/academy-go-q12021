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
	GetStudents(w http.ResponseWriter, r *http.Request)
	//DownloadCsv(w http.ResponseWriter, r *http.Request)
}

//type C struct{
//	controller Controller
//}
// router
func NewRouter(controller Controller)() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	})

	// GET students from csv
	router.HandleFunc("/readcsv", controller.GetStudents  ).Methods("GET")

	// Get students from url
	//router.HandleFunc(
	//	"/storedata",
	//	controller.DownloadCsv,
	//).Methods("GET")

	srv := runServer(router)

	log.Fatal(srv.ListenAndServe())

}


func runServer(router *mux.Router) *http.Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.Server.Address + ":" + strconv.Itoa(config.C.Server.Port),
		WriteTimeout: config.C.Server.Timeout * time.Second,
		ReadTimeout:  config.C.Server.Timeout * time.Second,
	}
	return srv
}
