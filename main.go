package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

var Employees []Employee

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Wizeline challenge</h1><p>Golang Bootcamp! Let's go!</p>")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(Employees)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(w, string(jsonData))
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/csv", returnAllEmployees)
	//router.HandleFunc("/csv", returnJson)
	log.Fatal(http.ListenAndServe(":8001", router))
}

func readCsv(filename string) ([][]string, error) {
	csvFile, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func main() {
	fmt.Println("Golang Challange - 1st Deliverable")
	records, err := readCsv("sample.csv")
	if err != nil {
		panic(err)
	}
	var emp Employee
	for _, record := range records {
		emp.ID = record[0]
		emp.EmployeeName = record[1]
		emp.EmployeeSalary = record[2]
		emp.EmployeeAge = record[3]
		emp.ProfileImage = record[4]
		Employees = append(Employees, emp)
	}
	handleRequests()
}
