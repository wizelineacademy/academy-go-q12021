package httputil

import (
	"encoding/json"
	"net/http"
)

// Response JSON-API v1.0 compliant object
type Response struct {
	Code     int         `json:"code"`
	Status   string      `json:"status,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// RespondJSON responds an HTTP request with JSON content
func RespondJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if code == 0 { // set default code
		code = http.StatusOK
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Response{
		Code:     code,
		Status:   http.StatusText(code),
		Data:     v,
		Metadata: nil,
	})
}
