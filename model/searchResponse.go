package model

// SearchResponse holds the structure for the response of external API
type SearchResponse struct {
	Doc      []Doc `json:"docs"`
	NumFound int   `json:"numFound"`
	Start    int   `json:"start"`
}

// Doc represents the actual result document from external API.
type Doc struct {
	Key       string `json:"key"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Published string `json:"first_published_year"`
}
