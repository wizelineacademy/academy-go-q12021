package model

import (
	"errors"
	"fmt"
)

type ConcurrentResponse struct {
	Error  error     `json:"error,omitempty"`
	Result []Pokemon `json:"result"`
	Total  int       `json:"total"`
	Items  int       `json:"count"`
}

func (r ConcurrentResponse) String() string {
	return fmt.Sprintf("\n{\n\tError: %v,\n\tResult: %v,\n\tTotal: %v,\n\tCount: %v\n}\n", r.Error, r.Result, r.Total, r.Items)
}

func (r ConcurrentResponse) NewErrorResponse(errorMessage string) Response {
	return ConcurrentResponse{Error: errors.New(errorMessage)}
}

func (r ConcurrentResponse) GetError() error {
	return r.Error
}
