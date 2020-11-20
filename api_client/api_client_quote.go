package api

import (
	"net/url"

	"github.com/etyberick/golang-bootcamp-2020/entity"
)

type quoteSourceClient struct {
	source *url.URL
}

//NewQuoteSourceClient will create an entity that represents the entity.QuoteSource interface
func NewQuoteSourceClient(url *url.URL) entity.QuoteSource {
	return &quoteSourceClient{url}
}

func (quoteSourceClient *quoteSourceClient) Fetch() (entity.Quote, error) {
	panic("Fetch not implemented") // TODO: Implement
}
