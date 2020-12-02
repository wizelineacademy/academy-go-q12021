package entity

// Quote represents a quote from https://quote-garden.herokuapp.com/api/v2/quotes/random
type Quote struct {
	ID     string `json:"_id"`
	Text   string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
	Genre  string `json:"quoteGenre"`
}

// Slice convertion
func (q Quote) Slice() []string {
	s := make([]string, 4)
	s[0] = q.ID
	s[1] = q.Text
	s[2] = q.Author
	s[3] = q.Genre
	return s
}
