package store

import "context"

type (
	Store struct {
		PersonStore
	}

	PersonStore struct {
	}
)
// Todo: add here the database connection
func NewStore(ctx context.Context) (*Store, error) {
	return &Store{
		PersonStore: PersonStore{},
	}, nil
}