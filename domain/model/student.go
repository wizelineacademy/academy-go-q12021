package model

import (
	"log"
	"strconv"
)

// Student - struct to store an student
type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      int    `json:"zip"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}

func (s Student) ToSlice() []string {
	id := strconv.Itoa(s.ID)
	zip:= strconv.Itoa(s.Zip)
	row := make([]string, 9, 9) // Since you only have 1 field in the struct
	row[0] = id
	row[1] = s.Name
	row[2] = s.LastName
	row[3] = s.Gender
	row[4] = s.City
	row[5] = s.State
	row[6] = zip
	row[7] = s.Email
	row[8] = s.Age
	return row
}
func (s Student) ToStruct(data []string) (Student ,error) {
	id, err := strconv.Atoi(data[0])
	checkError("Can't convert ID to int", err)
	zip, err := strconv.Atoi(data[6])
	checkError("Can't convert Zip to int", err)
	// fill struct with data
	student :=  Student{
		ID:       id,
		Name:     data[1],
		LastName: data[2],
		Gender:   data[3],
		City:     data[4],
		State:    data[5],
		Zip:      zip,
		Email:    data[7],
		Age:      data[8],
	}
	return student,err
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}