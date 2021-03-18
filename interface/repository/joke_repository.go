package repository

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/repository"
	"database/sql"
	"fmt"
)

type jokeRepository struct {
	db *sql.DB
}

func NewJokeRepository(db *sql.DB) repository.JokeRepository {
	return &jokeRepository{db}
}

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