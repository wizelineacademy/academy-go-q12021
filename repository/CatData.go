package repository

import (
	"bufio"
	"encoding/csv"
	"github.com/wl-project/academy-go-q12021/entities"
	"io"
	"log"
	"os"
)

func LoadData() []entities.CatFact {
	// Initialize with data
	var catFacts []entities.CatFact
	csvFile, _ := os.Open("catfacts.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		catFacts = append(catFacts, entities.CatFact{
			Id: line[0],
			Fact:  line[1],
		})
	}
	return catFacts
}