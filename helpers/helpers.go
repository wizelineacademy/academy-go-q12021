package helpers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// ServerError writes a 500 error to w and outputs a debug trace
func ServerError(w http.ResponseWriter, errorLog *log.Logger, err error) {
	//Debug trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description to the user.
func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound is simply a convenience wrapper around clientError which sends a 404 Not Found response to the user.
func NotFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}
