package repository

import (
	"bufio"
	"encoding/csv"
	"os"

	"github.com/ramrodo/golang-bootcamp-2020/model"
)

// FilmRepository - interface
type FilmRepository interface {
	FindAll() ([]model.Film, error)
}

// FindAll - reads CSV file and returns an array of Films
func FindAll() ([]model.Film, error) {
	films := []model.Film{}

	csvFile, err := os.Open("config/ghibliDB.csv")

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for i, row := range data {
		if i == 0 {
			continue
		}

		films = append(films, model.Film{
			ID:          row[0],
			Title:       row[1],
			Description: row[2],
			Director:    row[3],
			Producer:    row[4],
			ReleaseDate: row[5],
			RtScore:     row[6],
		})
	}

	return films, nil
}
