package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"golang-bootcamp-2020/domain/model"
	"os"
	"strconv"
)

type dbRepository struct {
}

type DataBaseRepository interface {
	CreateCharactersCSV(characters []model.Character) error
}

func NewDbRepository() DataBaseRepository {
	return &dbRepository{}
}

func (db *dbRepository) CreateCharactersCSV(characters []model.Character) error {
	file, err := os.Create("./characters.csv")
	// TODO: handle this error
	defer file.Close()

	if err != nil {
		return errors.New("error writing file")
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, ch := range characters {
		var row []string
		row = append(row, strconv.Itoa(ch.Id))
		row = append(row, ch.Name)
		row = append(row, ch.Status)
		row = append(row, ch.Species)
		row = append(row, ch.Type)
		row = append(row, ch.Gender)
		row = append(row, ch.Origin.Name)
		row = append(row, ch.Origin.Url)
		row = append(row, ch.Location.Name)
		row = append(row, ch.Location.Url)
		row = append(row, ch.Image)

		var episodesString string
		for _, episode := range ch.Episodes {
			episodesString += fmt.Sprintf("+%s", episode)
		}

		row = append(row, episodesString)
		//TODO: handle error
		writer.Write(row)
	}
	return nil
}
