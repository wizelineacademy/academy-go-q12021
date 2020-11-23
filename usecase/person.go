package usecase

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pankecho/golang-bootcamp-2020/entity"
)

type Headers int

const (
	// CSV Headers
	ID 			Headers = iota
	FirstName
	LastName

	NumOfHeaders = 3
)

type PersonStore interface {
	CreatePerson(ctx echo.Context, person *entity.Person) (*entity.Person, error)
}

type Person struct {
	PersonStore	PersonStore
}

func NewPerson(store PersonStore) Person {
	return Person{
		PersonStore: store,
	}
}

func (p Person) UploadCSV(ctx echo.Context, data [][]string) ([]*entity.Person, error) {
	if len(data) <= 1 {
		return nil, errors.New("empty csv")
	}

	people := []*entity.Person{}
	for i, record := range data {
		if i == 0 {
			if !validHeaders(record) {
				return nil, errors.New("invalid csv headers")
			}
			continue
		}
		if len(record) != NumOfHeaders {
			return nil, errors.New("invalid record")
		}

		people = append(people, &entity.Person{
			ID: 		record[ID],
			FirstName:	record[FirstName],
			LastName:	record[LastName],
		})
	}
	for _, person := range people {
		_, err := p.PersonStore.CreatePerson(ctx, person)
		if err != nil {
			return people, err
		}
	}
	return people, nil
}

func (h Headers) String() string {
	return [...]string {
		"ID",
		"First Name",
		"Last Name",
	}[h]
}

func validHeaders(record []string) bool {
	if len(record) != NumOfHeaders {
		return false
	}
	for i := range record {
		if record[i] != Headers(i).String() {
			return false
		}
	}
	return true
}