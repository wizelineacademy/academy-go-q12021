package main

import (
	"models"
	"utils"
	"fmt"
	"log"
	"net/http"
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)

	viperenv := utils.GetEnvVar("MONGO_URL")

	fmt.Printf("viper : %s = %s \n", "MONGO_URL", viperenv)

	log.Fatal(server)
}

func ReadCSV() models.PokemonList {
	var pokeList models.PokemonList

	recordFile, err := os.Open("pokemon.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return pokeList
	}

	reader := csv.NewReader(recordFile)

	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return pokeList
	}

	for _, pokemon := range allRecords {
		id, err := strconv.Atoi(pokemon[0])
		if err != nil {
			 // handle error
		}
		poke := models.Pokemon{Id:id, Name:pokemon[1], Types:pokemon[2], Region:pokemon[3]}
		pokeList = append(pokeList, poke)
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return pokeList
	}

	return pokeList
}

// func csvReaderRow() {
// 	// Open the file
// 	recordFile, err := os.Open("pokemon.csv")
// 	if err != nil {
// 		fmt.Println("An error encountered ::", err)
// 		return
// 	}

// 	// Setup the reader
// 	reader := csv.NewReader(recordFile)

// 	// Read the records
// 	header, err := reader.Read()
// 	if err != nil {
// 		fmt.Println("An error encountered ::", err)
// 		return
// 	}
// 	fmt.Printf("Headers : %v \n", header)

// 	for i:= 0 ;; i = i + 1 {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break // reached end of the file
// 		} else if err != nil {
// 			fmt.Println("An error encountered ::", err)
// 			return
// 		}

// 		fmt.Printf("Row %d : %v \n", i, record)
// 	}

// 	// Note: Each time Read() is called, it reads the next line from the file
// 	// r1, _ := reader.Read() // Reads the first row, useful for headers
// 	// r2, _ := reader.Read() // Reads the second row
// }
