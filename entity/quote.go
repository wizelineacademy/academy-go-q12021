package entity

//Quote represents a quote from https://quote-garden.herokuapp.com/api/v2/quotes/random
type Quote struct {
	ID          string `json:"_id"`
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	QuoteGenre  string `json:"quoteGenre"`
}

//QuoteStorage is access operations use case
type QuoteStorage interface {
	ReadAll() ([]Quote, error) //Fetches all items from data source
	Write(quote *Quote) error  //Writes an item into data source
}

//QuoteSource is quote source use case
type QuoteSource interface {
	Fetch() (Quote, error) //Retrieves a Quote
}
