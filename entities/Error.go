package entities

type HTTPError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
