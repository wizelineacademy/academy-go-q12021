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

func (e e) Error() string {
	return fmt.Sprintf("message: %s - code: %d", e.ErrorMessage, e.ErrorCode)
}

func (e e) Code() int {
	return e.ErrorCode
}

func (e e) Message() string {
	return e.ErrorMessage
}

func NewInternalServerError(message string) RestError {
	return e{
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: message,
	}
}

func NewBadRequestError(message string) RestError {
	return &e{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewNotFoundError(message string) RestError {
	return &e{
		ErrorCode:    http.StatusNotFound,
		ErrorMessage: message,
	}
}
