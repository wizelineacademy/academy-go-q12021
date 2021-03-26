package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

/* Generic functions and structure */
type Response struct {
	Title string `json:"title"`
	Message string `json:"message"`
}
type Item struct {
	Id string `json:"id"`
    Title string `json:"title"`
	Years string `json:"years"`
}

func ConvertStructToJSON(obj interface{}) string {
    e, err := json.Marshal(obj)
    if err != nil {
        return err.Error()
    }
    return string(e)
}

func displayError(w http.ResponseWriter, message string) {
    log.Println(message)
	fmt.Fprintf(w, "%s", ConvertStructToJSON(Response{Title: "Error", Message: message}))

}

func GetDataFromCSVFile(filePath string) ([][] string)  {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("\n- An error ocurred while reading the file. \n", err)
	} else {
		fmt.Println("\nSuccessfully Opened CSV file") 
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
		log.Println("\n- An error ocurred while reading the file. \n", err)
		return nil
    }
	return csvLines
}

/* Non Generic functions */

func ParseCSVDataToItemsList(csvLines [][]string) (listOfItems []Item ) { 
	// Convert csv lines to a generic item structure and append them to the array of items
    for _, line := range csvLines {
		newItem := Item{
            Id: line[0],
            Title: line[1],
			Years: line[2],
        }
        listOfItems = append(listOfItems, newItem)
        log.Println(newItem.Id + " " + newItem.Title + " ")
    }
	return 
}

/* Endpoint Functions */

func getLanguages(w http.ResponseWriter, r *http.Request) {
	csvLines := GetDataFromCSVFile("data.csv")
	listOfItems :=  ParseCSVDataToItemsList(csvLines)   
	fmt.Fprintf(w, "%s", ConvertStructToJSON(listOfItems))
}

func getLanguageById(w http.ResponseWriter, r *http.Request) {
	csvLines := GetDataFromCSVFile("data.csv")
	listOfItems :=  ParseCSVDataToItemsList(csvLines)   
	// Obtain the query param id number from URL
	keys, ok := r.URL.Query()["id"]
    if !ok || len(keys[0]) < 1 {
		displayError(w, "Url Param 'id' is missing!")
        return
    }
	// Casting the string number to an integer
    id, err := strconv.Atoi(keys[0])
	if err != nil {
		displayError(w, "The Id provided is wrong, please check it!")
		return
	}
	// Validations: number is positive and that exists as index in the slice 
	if (id >= len(listOfItems) || id < 0) {
		displayError(w, "The Id doesn't seem to exist!")
		return
	}
	// Get the object from slice using the id as index
	obj := listOfItems[id]
    fmt.Fprintf(w, "%s", ConvertStructToJSON(obj))
}

func main() {
    http.HandleFunc("/getLanguages", getLanguages)
	http.HandleFunc("/getLanguageById", getLanguageById)
	log.Println("Server running succesfully on port 8080!")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

