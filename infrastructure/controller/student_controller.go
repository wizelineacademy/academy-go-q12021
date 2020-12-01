package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-bootcamp-2020/domain/model"
)

type Usecase interface {
	GetStudentsFromCsv() ([]model.Student, error)
}

type Students struct{
	students Usecase
}

func New(u Usecase) *Students{
	return &Students{students: u }
}

// GetStudents
func (s *Students) GetStudents(writer http.ResponseWriter, request *http.Request) {
	students, err := s.students.GetStudentsFromCsv()
	if err != nil {
		log.Fatal(err)
		respondWithError(writer, 404, err.Error())
	}
	respondWithJson(writer, 200, students)
}

// dwonload csv
func DownloadCsv(writer http.ResponseWriter, request *http.Request) {
	//usecase.
	respondWithJson(writer, 200, map[string]bool{"download": true})
}

// Download Db from csv
//func DownloadDb(writer http.ResponseWriter, request *http.Request) {

//student, err := datastore.MongoDAO.FindAll()
//err != nil{
//	respondWithError(writer, http.StatusInternalServerError, err.Error()),
//	return
//}
//}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
