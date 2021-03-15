// go:generate mockgen -source=csv.go -destination=mock/csv_mock.go -package mock
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
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	SavePokemons(*[]model.SinglePokeExternal) *model.Error
	GetPokemonsConcurrently(items int, itemsPerWorker int) ([]model.Pokemon, *model.Error)
}

func New() *CsvService {
	return &CsvService{}
}

func Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

func OpenAndWrite(path string) (*os.File, error) {
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

func ReadConcurrently(f *os.File, items int, itemsPerWorker int) ([]model.Pokemon, *model.Error) {

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	var pokemons []model.Pokemon = nil
	for i := 0; i < items; i++ {
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

func ReadAllLines(f *os.File) ([][]string, *model.Error) {
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

func AddLine(f *os.File, lines [][]string, newPokes *[]model.SinglePokeExternal) *model.Error {

	linesNumber := len(lines) + 1

	w := csv.NewWriter(f)
	for _, pokemon := range *newPokes {
		w.Write([]string{strconv.Itoa(linesNumber), pokemon.Name, pokemon.URL})
		linesNumber = linesNumber + 1
	}
	defer w.Flush()

	return nil
}

func (s *CsvService) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	f, err := Open(pathFile)

	if err != nil {
		return model.Pokemon{}, &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	pokes, errorReading := Read(f)

	if errorReading != nil {
		errorReading := model.Error{
			Code:    http.StatusInternalServerError,
			Message: errorReading,
		}
		return model.Pokemon{}, &errorReading
	}

	var tempPokemon model.Pokemon

	for _, pokemon := range pokes {
		if pokemon.ID == pokemonId {
			tempPokemon = pokemon
			break
		}
	}

	if tempPokemon == (model.Pokemon{}) { //Check for unexisting pokemon
		return model.Pokemon{}, &model.Error{
			Code:    http.StatusAccepted,
			Message: "The pokemon does not exists",
		}
	}

	return tempPokemon, nil
}

func (s *CsvService) GetPokemons() ([]model.Pokemon, *model.Error) {
	f, err := Open(pathFile)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	pokes, errorReading := Read(f)

	if errorReading != nil {
		errorReading := model.Error{
			Code:    http.StatusInternalServerError,
			Message: errorReading,
		}
		return nil, &errorReading
	}

	return pokes, nil
}

func (s *CsvService) SavePokemons(newPokemons *[]model.SinglePokeExternal) *model.Error {
	f, _ := Open(pathFile) //Read only
	lines, _ := ReadAllLines(f)
	fileOpenAndWrite, _ := OpenAndWrite(pathFile) // Write

	err := AddLine(fileOpenAndWrite, lines, newPokemons)
	if err != nil {
		return err
	}

	return nil
}

func (s *CsvService) GetPokemonsConcurrently(items int, itemsPerWorker int) ([]model.Pokemon, *model.Error) {
	f, err := Open(pathFile)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	pokes, errorReading := ReadConcurrently(f, items, itemsPerWorker)

	if errorReading != nil {
		errorReading := model.Error{
			Code:    http.StatusInternalServerError,
			Message: errorReading,
		}
		return nil, &errorReading
	}

	return pokes, nil
}
