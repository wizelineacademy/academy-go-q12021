package csvservice

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"pokeapi/model"
)

const pathFile = "./csv/pokemon.csv"

type CsvService struct{}

type NewCsvService interface {
	GetPokemons(f *os.File) ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int, f *os.File) (model.Pokemon, *model.Error)
	AddLine(f *os.File, l [][]string, newPokes *[]model.SinglePokeExternal) *model.Error
	Open(path string) (*os.File, error)
	OpenAndWrite(path string) (*os.File, error)
	ReadAllLines(f *os.File) ([][]string, *model.Error)
}

func New() *CsvService {
	return &CsvService{}
}

func (s *CsvService) Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

func (s *CsvService) OpenAndWrite(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

func Read(f *os.File) ([]model.Pokemon, *model.Error) {

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	var pokemons []model.Pokemon = nil
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
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
	defer f.Close()

	return pokemons, nil
}

func (s *CsvService) ReadAllLines(f *os.File) ([][]string, *model.Error) {
	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		e := model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error trying to read the lines of the file",
		}
		return nil, &e
	}

	defer f.Close()

	return lines, nil
}

func (s *CsvService) AddLine(f *os.File, lines [][]string, newPokes *[]model.SinglePokeExternal) *model.Error {

	linesNumber := len(lines) + 1

	w := csv.NewWriter(f)
	for _, pokemon := range *newPokes {
		w.Write([]string{strconv.Itoa(linesNumber), pokemon.Name, pokemon.URL})
		linesNumber = linesNumber + 1
	}
	defer w.Flush()

	return nil
}

func (s *CsvService) GetPokemon(pokemonId int, f *os.File) (model.Pokemon, *model.Error) {

	pokes, err := Read(f)

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

func (s *CsvService) GetPokemons(f *os.File) ([]model.Pokemon, *model.Error) {

	pokes, err := Read(f)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	return pokes, nil
}
