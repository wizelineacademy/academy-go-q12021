package dataload

import (
	"academy/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var allJokes []model.Joke

//Load this will load the data from the csv
func LoadData() {
	pwd, _ := os.Getwd()
	csvFile, err := os.Open(pwd + "/data/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(records) == 0 {
		fmt.Println("Currently there are not Jokes Stored... this is not funny!")
		os.Exit(1)
	}
	for _, rec := range records {
		var insert model.Joke
		insert.ID, err = strconv.Atoi(rec[0])
		insert.Setup = rec[1]
		insert.Punchline = rec[2]
		allJokes = append(allJokes, insert)
	}

}

//ReadData will return the data from the file
func ReadData() []model.Joke {
	return allJokes
}
