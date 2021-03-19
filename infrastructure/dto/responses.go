package dto

import "time"

type errorRespose struct {
	Message     string    `json:"message"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

func NewErrorResponse(msg string, desc string) errorRespose {
	return errorRespose{Message: msg, Description: desc, Timestamp: time.Now()}
}

type Respose struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}
