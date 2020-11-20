package handlers

type responseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"errorMessage"`
}
