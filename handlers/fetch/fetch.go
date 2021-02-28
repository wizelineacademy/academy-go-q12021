package fetch

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Joke Struct
type Joke struct {
	ID        int    `json:id`
	Setup     string `json:setup`
	Punchline string `json:punchline`
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

//GetData Updates the CSV from the target API
func GetData(w http.ResponseWriter, r *http.Request) {
	url := "https://official-joke-api.appspot.com/random_ten"
	request, error := http.NewRequest("GET", url, nil)
	if error != nil {
		fmt.Println(error)
	}
	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	// Unmarshal JSON data
	var jsonData []Joke
	json.Unmarshal([]byte(data), &jsonData)

	pwd, err := os.Getwd()
	csvFile, err := os.OpenFile(pwd+"/data/data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
	records := readCsvFile(pwd + "/data/data.csv")
	counter := len(records) + 1

	writer := csv.NewWriter(csvFile)

	for i, usance := range jsonData {
		var row []string
		row = append(row, strconv.Itoa(counter+i))
		row = append(row, strings.Replace(usance.Setup, "\n", "", -1))
		row = append(row, usance.Punchline)
		writer.Write(row)
	}

	// remember to flush!
	writer.Flush()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
