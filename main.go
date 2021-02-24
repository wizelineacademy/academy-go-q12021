
package main

import (
    "fmt"
    "log"
	"net/http"
	"io"
	"encoding/csv"
	"os"
)

func getInfo(w http.ResponseWriter, r *http.Request){
	keys, ok := r.URL.Query()["id"]
	found := false
    
    if !ok || len(keys[0]) < 1 {
        fmt.Fprintf(w, "Url Param 'id' is missing")
        log.Println("Url Param 'id' is missing")
        return
	}
	key := keys[0]

	// Open the file
	csvfile, err := os.Open("data.csv")
	if err != nil {
        fmt.Fprintf(w, "Couldn't open the csv file")
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	reader := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[0] == key {
			fmt.Fprintf(w, "%s", record[1])
			found = true
			break
		} 
	}

	if !found {
		fmt.Fprintf(w, "Not found data with id: %s", key)
	}
	
	fmt.Println("Endpoint Hit: get-info")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get-info", getInfo)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}