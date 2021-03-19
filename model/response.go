package model

import "fmt"

type Response struct {
	Error  error     `json:"error,omitempty"`
	Result []Pokemon `json:"result"`
	Total  int       `json:"total"`
	Page   int       `json:"page"`
	Count  int       `json:"count"`
}

func (r Response) String() string {
	return fmt.Sprintf("\n{\n\tError: %v,\n\tResult: %v\n,\tTotal: %v,\n\tPage: %v,\n\tCount: %v\n}\n", r.Error, r.Result, r.Total, r.Page, r.Count)
}
