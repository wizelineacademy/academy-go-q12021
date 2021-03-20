package datastore

import (
	"log"
	"fmt"
	"io"
	"os"
	"errors"
	"encoding/csv"
	"strconv"
	"github.com/halarcon-wizeline/academy-go-q12021/domain"
)

// Read csv file
func readCsvPokemons(file string) ([]domain.Pokemon, error) {

	var pokemons []domain.Pokemon

	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		return pokemons, errors.New("Couldn't open the csv file")
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("Reading pokemon: %s %s\n", record[0], record[1])
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return pokemons, errors.New("Error: Pokemon does not have a valid ID\n")
		}
		pokemon := domain.Pokemon {ID:id, Name:record[1]}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func GetPokemonDB() ([]domain.Pokemon, error) {

	pokemons, err := readCsvPokemons("./infrastructure/datastore/pokemons.csv")

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

// Write csv file
func writeCsvPokemons(file string, pokemons []domain.Pokemon) (string, error) {

	// Create the file
	f, err := os.Create(file)
	if err != nil {
		return "", errors.New("Couldn't create the file")
	}
	defer f.Close()

	for _, pok := range pokemons {
		strPokemon := strconv.Itoa(pok.ID) + ", " + pok.Name
		fmt.Fprintln(f, strPokemon)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Error writing file")
		}
	}

	return file, nil
}

func CreatePokemonDB(file string, pokemons []domain.Pokemon) (string, error) {

	_, err := writeCsvPokemons(file, pokemons)

	if err != nil {
		return "", err
	}

	return file, nil
}