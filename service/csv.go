package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"pokeapi/model"
	"strconv"
)

const pathFile = "./csv/pokemon.csv"

var pokemons []model.Pokemon = nil
var csvFile *os.File

type CsvService struct{}

type NewCsvService interface {
	GetPokemons() []model.Pokemon
	GetPokemon(pokemonId int) model.Pokemon
}

func NewCsv() *CsvService {
	return &CsvService{}
}

func openCsv() {
	f, err := os.Open(pathFile)
	csvFile = f
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
	}
}

func readCsv() []model.Pokemon {
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
		}
		tempPokemon := model.Pokemon{
			Name: line[1],
			URL:  line[2],
		}

		if line[0] != "" {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Printf("There was an error trying to process the ID: %v\n", err)
			}
			tempPokemon.ID = id
		}

		pokemons = append(pokemons, tempPokemon)
	}
	defer csvFile.Close()
	return pokemons
}

func addLineCsv(newPokes *[]model.SinglePokeExternal) {
	openCsv()
	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	linesNumber := len(lines) + 1
	defer csvFile.Close()

	f, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	for _, pokemon := range *newPokes {
		w.Write([]string{strconv.Itoa(linesNumber), pokemon.Name, pokemon.URL})
		linesNumber = linesNumber + 1
	}
	w.Flush()
}

func (s *CsvService) GetPokemon(pokemonId int) model.Pokemon {
	pokes := readCsv()
	var tempPokemon model.Pokemon

	for _, pokemon := range pokes {
		if pokemon.ID == pokemonId {
			tempPokemon = pokemon
		}
	}

	return tempPokemon
}

func (s *CsvService) GetPokemons() []model.Pokemon {
	pokes := readCsv()
	return pokes
}
