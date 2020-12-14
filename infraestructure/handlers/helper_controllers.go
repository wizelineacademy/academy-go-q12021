package handlers

import (
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, r *http.Request, httpStatusCode int, errorMessage string) {
	response := &responseError{
		StatusCode: httpStatusCode,
		Message:    errorMessage,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func jsonWritter(w http.ResponseWriter, r *http.Request, httpStatusCode int, sendableData interface{}) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(sendableData)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "Error while parsing response object")
		return
	}
}

func bytesWritter(w http.ResponseWriter, r *http.Request, httpStatusCode int, sendableData []byte) {
	w.WriteHeader(httpStatusCode)
	w.Write(sendableData)
}
