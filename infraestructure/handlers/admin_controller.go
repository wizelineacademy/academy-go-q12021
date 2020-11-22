package handlers

import (
	"net/http"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/registry"
)

//GetLogs is the handler for retrieving the stored logs in a period of time
// and returning it as json
func GetLogs(w http.ResponseWriter, r *http.Request) {
	ai := registry.NewAdminInteractor()
	ai.GetLogs("users", time.Time{}, time.Time{})
}
