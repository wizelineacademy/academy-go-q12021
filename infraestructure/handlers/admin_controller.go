package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/datastore"
	"github.com/alexis-aguirre/golang-bootcamp-2020/registry"
)

//GetLogs is the handler for retrieving the stored logs in a period of time
// and returning it as json
func GetLogs(w http.ResponseWriter, r *http.Request) {
	log.Println("Get logs handler")
	ai := registry.NewAdminInteractor()
	logs, err := ai.GetLogs("users", time.Time{}, time.Time{})
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "")
		return
	}

	log.Println(datastore.UserLogToJson(logs))
	//body, err := json.Marshal(logs)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "")
		return
	}
	bytesWritter(w, r, http.StatusOK, []byte(datastore.UserLogToJson(logs)))
}
