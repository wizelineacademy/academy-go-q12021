package model

type Response interface {
	NewErrorResponse(errorMessage string) Response
	GetError() error
}
