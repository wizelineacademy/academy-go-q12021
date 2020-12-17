package entity

// QuoteGarden represents the response from https://quote-garden.herokuapp.com/api/v3/quotes/random
type QuoteGarden struct {
	StatusCode  int        `json:"statusCode"`
	Message     string     `json:"message"`
	Pagination  Pagination `json:"pagination"`
	TotalQuotes int        `json:"totalQuotes"`
	Data        []Quote    `json:"data"`
}

// Pagination contains information about the current quotes "page"
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	NextPage    int `json:"nextPage"`
	TotalPages  int `json:"totalPages"`
}
