package csvstore

import (
  "github.com/pkg/errors"

  dataCard "github.com/gbrayhan/academy-go-q12021/data/card"
  domainCard "github.com/gbrayhan/academy-go-q12021/domain/card"
  domainErrors "github.com/gbrayhan/academy-go-q12021/domain/errors"
)

const (
  createError = "Error in creating new author"
  readError   = "Error in finding author in the database"
  listError   = "Error in getting authors from the database"
)

// Store struct manages interactions with authors store
type Store struct {
  csv       FileCSV
  booksRepo domainCard.CardRepository
}

func New() *Store {
  return &Store{
  }
}

func (s *Store) CreateCard(cardDom *domainCard.Card) (card *domainCard.Card, err error) {

  return
}

func (s *Store) ReadCard(id int) (card *domainCard.Card, err error) {

  dCard, err := s.csv.FindCardByID(id)

  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
    return
  }

  if dCard.Name == "" {
    err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
    return
  }

  card = dataCard.ToDomainModel(&dCard)
  return
}

func (s *Store) ListCards() (cards []domainCard.Card, err error) {
  var results []dataCard.Card
  err = s.csv.FindAllCards(&results)

  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
    return
  }

  cards = make([]domainCard.Card, len(results))

  for i, element := range results {
    cards[i] = *dataCard.ToDomainModel(&element)
  }

  return
}
