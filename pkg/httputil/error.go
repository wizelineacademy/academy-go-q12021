package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/maestre3d/academy-go-q12021/internal/domain"
)

// Error JSON-API v1.0 compliant object
type Error struct {
	Code        int         `json:"code"`
	Status      string      `json:"status,omitempty"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	Source      sourceErr   `json:"source,omitempty"`
	Metadata    interface{} `json:"metadata,omitempty"`
}

type sourceErr struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter,omitempty"`
}

// GetErrStatus get an http status from the given error
func GetErrStatus(err error) int {
	domainErr, ok := err.(domain.Error)
	if !ok {
		return http.StatusInternalServerError
	}
	switch {
	case domainErr.IsAlreadyExists():
		return http.StatusConflict
	case domainErr.IsNotFound():
		return http.StatusNotFound
	case domainErr.IsDomain():
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// RespondErrJSON responds to an HTTP request with the given error as JSON
func RespondErrJSON(w http.ResponseWriter, r *http.Request, err error) {
	code := GetErrStatus(err)
	title := err.Error()
	if code == http.StatusInternalServerError {
		title = "something happened" // hide internal errors to clients, they are still logged out through stdout
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Error{
		Code:        code,
		Status:      http.StatusText(code),
		Title:       title,
		Description: "",
		Source: sourceErr{
			Pointer:   r.RequestURI,
			Parameter: r.URL.RawQuery,
		},
		Metadata: nil,
	})
}
