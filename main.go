package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/fatih/structs"
)

type SearchResponse struct {
	Doc []Doc `json:"docs"`
	NumFound int `json:"numFound"`
	Start int `json:"start"`
}

type Doc struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Type string `json:"type"`
	Published string `json:"first_published_year"`

}

func main() {
	fmt.Println("Init")

	response, err := http.Get("http://openlibrary.org/search.json?q=the+lord+of+the+rings&page=1")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject SearchResponse

	json.Unmarshal(responseData, &responseObject)

	file, err := os.Create("/home/omar/Desktop/saved.csv")
	defer file.Close()

	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	SaveCsv(responseObject, *file)
}

func SaveCsv(responseObject SearchResponse, file os.File) {
	writer := csv.NewWriter(&file)
	defer writer.Flush()
	for _, doc := range responseObject.Doc {
		err := writer.Write(interfaceToString(structs.Values(doc)))
		if err != nil {
			log.Fatal("cannot write CSV", err)
		}
	}
}

func interfaceToString(record []interface{}) []string {
	var a []string

	for _, row := range record {
		a = append(a,row.(string))
	}
	return a
}
