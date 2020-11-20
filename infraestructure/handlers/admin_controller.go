package handlers

import (
	"net/http"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/registry"
)

//GetLogs is the handler for
func GetLogs(w http.ResponseWriter, r *http.Request) {
	ai := registry.NewAdminInteractor()
	ai.GetLogs("users", time.Time{}, time.Time{})
}
