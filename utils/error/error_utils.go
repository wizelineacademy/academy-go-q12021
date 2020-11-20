package error

import (
	"fmt"
	"net/http"
)

type e struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"message"`
}

type RestError interface {
	Error() string
	Code() int
	Message() string
}

//Return custom error message
func (e e) Error() string {
	return fmt.Sprintf("message: %s - code: %d", e.ErrorMessage, e.ErrorCode)
}

//Return status code
func (e e) Code() int {
	return e.ErrorCode
}

//Return error messsage
func (e e) Message() string {
	return e.ErrorMessage
}

//Return internal server error given error message
func NewInternalServerError(message string) RestError {
	return e{
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: message,
	}
}

//Return bad request error given error message
func NewBadRequestError(message string) RestError {
	return &e{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: message,
	}
}

//Return not found error given error message
func NewNotFoundError(message string) RestError {
	return &e{
		ErrorCode:    http.StatusNotFound,
		ErrorMessage: message,
	}
}
