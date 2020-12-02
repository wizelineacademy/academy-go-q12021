package entity

// QuoteGarden represents the response from https://quote-garden.herokuapp.com/api/v2/quotes/random
type QuoteGarden struct {
	StatusCode int   `json:"statusCode"`
	Quote      Quote `json:"quote"`
}
