package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

var csvFile *os.File

func readCsv() {
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
		tempPokemon := Pokemon{
			Name:   line[1],
			Weight: line[2],
			Height: line[3],
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
}

func openCsv() {
	f, err := os.Open("./csv/pokemon.csv")
	csvFile = f
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
	}

}
