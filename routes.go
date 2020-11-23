package main

import (
	"encoding/json"
	"net/http"
)

func getCentres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := db.Centres.GetAll()

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}