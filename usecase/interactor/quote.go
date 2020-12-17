package interactor

import (
	"encoding/json"

	"github.com/etyberick/golang-bootcamp-2020/entity"
	"github.com/etyberick/golang-bootcamp-2020/service/quotegarden"
	"github.com/etyberick/golang-bootcamp-2020/service/repository"
)

type quoteInteractor struct {
	Repository repository.QuoteRepository
	Gateway    quotegarden.Client
}

// QuoteInteractor operates the business rules
type QuoteInteractor interface {
	GetAll() ([]byte, error)
	Update() ([]byte, error)
}

// NewQuoteInteractor returns a new instance of a quoteInteractor
func NewQuoteInteractor(c entity.Config) QuoteInteractor {
	return &quoteInteractor{
		Repository: repository.NewQuoteRepository(c.CSVFilepath),
		Gateway:    quotegarden.NewClient(),
	}
}

func (qi *quoteInteractor) GetAll() ([]byte, error) {
	// Get all quotes
	q, err := qi.Repository.ReadAll()
	if err != nil {
		return nil, err
	}

	// Convert them to JSON
	jq, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return jq, nil
}

func (qi *quoteInteractor) Update() ([]byte, error) {
	// Get a new quote
	q, err := qi.Gateway.GetQuote()
	if err != nil {
		return nil, err
	}

	// Write quote into database
	qi.Repository.Write(q.Data[0])
	jm, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return jm, nil
}
