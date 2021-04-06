package repository

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/repository"
	"database/sql"
	"fmt"
)

// jokeRepository struct for SQL DB
type jokeRepository struct {
	db *sql.DB
}

// NewJokeRepository returns a JokeRepository
func NewJokeRepository(db *sql.DB) repository.JokeRepository {
	return &jokeRepository{db}
}

// GetAll return all jokes fro third party REST endpoint
func (jr *jokeRepository) GetAll(jokes []*model.Joke) ([]*model.Joke, error) {

	c := NewJokeClient()
	jokes, err := c.GetJoke()

	if err != nil {
		return nil, err
	}

	for _, joke := range jokes {
		queryString := "INSERT INTO jokes(id, joke) VALUES ('"+joke.ID+"', \""+joke.Joke+"\");"
		fmt.Println("queryString=", queryString)
		rws, err := jr.db.Exec(queryString)

		if err != nil {
			return nil, err
		}

		lastInsertId, err := rws.LastInsertId()
		rowsAffected, err := rws.RowsAffected()

		fmt.Printf("LastInsertId = %d, %d \n", lastInsertId , rowsAffected)
	}




	return jokes, nil
}