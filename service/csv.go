package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"pokeapi/model"
	"strconv"
)

const pathFile = "./csv/pokemon.csv"

var pokemons []model.Pokemon = nil
var csvFile *os.File

type CsvService struct{}

type NewCsvService interface {
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	AddLineCsv(newPokes *[]model.SinglePokeExternal) *model.Error
}

func New() *CsvService {
	return &CsvService{}
}

func openCsv() {
	f, err := os.Open(pathFile)
	csvFile = f
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
	}
}

func readCsv() ([]model.Pokemon, *model.Error) {
	openCsv()
	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("There was an error reading something in line: %v\n", err)

			err := model.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return nil, &err
		}
		tempPokemon := model.Pokemon{
			Name: line[1],
			URL:  line[2],
		}

		if line[0] != "" {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Printf("There was an error trying to process the ID: %v\n", err)

				err := model.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				return nil, &err
			}
			tempPokemon.ID = id
		}

		pokemons = append(pokemons, tempPokemon)
	}
	defer csvFile.Close()

	return pokemons, nil
}

func (s *CsvService) AddLineCsv(newPokes *[]model.SinglePokeExternal) *model.Error {
	openCsv()
	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)

		e := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return &e
	}
	linesNumber := len(lines) + 1
	fmt.Println("Number of lines", linesNumber)
	defer csvFile.Close()

	f, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)

		e := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return &e
	}
	w := csv.NewWriter(f)
	for _, pokemon := range *newPokes {
		fmt.Println("Pokemon: ", pokemon.Name)
		w.Write([]string{strconv.Itoa(linesNumber), pokemon.Name, pokemon.URL})
		linesNumber = linesNumber + 1
	}
	defer w.Flush()

	return nil
}

func (s *CsvService) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	pokes, err := readCsv()

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return model.Pokemon{}, &err
	}

	var tempPokemon model.Pokemon

	for _, pokemon := range pokes {
		if pokemon.ID == pokemonId {
			tempPokemon = pokemon
		}
	}

	return tempPokemon, nil
}

func (s *CsvService) GetPokemons() ([]model.Pokemon, *model.Error) {
	pokes, err := readCsv()

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	return pokes, nil
}
