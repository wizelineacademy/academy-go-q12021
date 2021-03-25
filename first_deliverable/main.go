package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ProgrammingLanguage struct {
	Id string
    Title string 
}

func displayError(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, "%v", message)
    log.Println(message)
}

func OpenCSV(filePath string) (*os.File)  {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("\n- An error ocurred while reading the file. \n", err)
	} else {
		fmt.Println("\nSuccessfully Opened CSV file") 
	}
	return csvFile
}


func PrintDataFromCSVFile(filePath string) (listOfProgrammingLanguages []ProgrammingLanguage ) { 
    csvFile := OpenCSV(filePath)
    defer csvFile.Close()

    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
		fmt.Println("\n- An error ocurred while reading the file. \n", err)
		return
    }

    for _, line := range csvLines {
		newProgrammingLanguage := ProgrammingLanguage{
            Id: line[0],
            Title: line[1],
        }
        listOfProgrammingLanguages = append(listOfProgrammingLanguages, newProgrammingLanguage)
        fmt.Println(newProgrammingLanguage.Id + " " + newProgrammingLanguage.Title + " ")
    }
	csvFile.Close()
	return 
}


func getAllLanguages(w http.ResponseWriter, r *http.Request) {
	listOfProgrammingLanguages := PrintDataFromCSVFile("data.csv")
    for  _, value := range listOfProgrammingLanguages {
    	fmt.Fprintf(w, "\n\tId: %s, Title: %s", value.Id, value.Title)
	}
}



func getLanguageById(w http.ResponseWriter, r *http.Request) {
	listOfProgrammingLanguages := PrintDataFromCSVFile("data.csv")

	keys, ok := r.URL.Query()["id"]
    if !ok || len(keys[0]) < 1 {
		displayError(w, "Url Param 'id' is missing!")
        return
    }
    id, err := strconv.Atoi(keys[0])
	if err != nil {
		displayError(w, "The Id provided is wrong, please check it!")
		return
	}
	if (id >= len(listOfProgrammingLanguages) || id < 0) {
		displayError(w, "The Id doesn't seem to exist!")
		return
	}

    fmt.Fprintf(w, "\nId: %s \nTitle: %s", listOfProgrammingLanguages[id].Id, listOfProgrammingLanguages[id].Title)
}

func main() {
    http.HandleFunc("/getAllLanguages", getAllLanguages)
	http.HandleFunc("/getLanguageById", getLanguageById)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

