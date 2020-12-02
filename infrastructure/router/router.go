package router

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const badRequest = 405

// HandleRequest is the handler of my routes
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method)

	switch r.Method {
	case http.MethodGet:
		readCSV(w, r)
		break
	default:
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
		break
	}
}

func readCSV(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("infrastructure/datastore/astronauts.csv")
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var data [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading line %v", err)
		}

		data = append(data, record)
	}

	fmt.Println(data)
}
