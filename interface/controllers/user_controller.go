package controllers

import (
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("GetUsers"))
}
