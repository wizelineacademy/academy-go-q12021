package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
	variables := r.URL.Query()
	log.Println(variables)
	queryVariable, ok1 := variables["q"]
	typeVariable, ok2 := variables["type"]

	if !ok1 || !ok2 {
		Error(w, r, http.StatusBadRequest, "Malformed query")
		return
	}
	fmt.Println(queryVariable, typeVariable)
}
