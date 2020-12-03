package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-bootcamp-2020/domain/model"
)

// Usecase interface
type Usecase interface {
	GetStudentsService() ([]model.Student, error)
	//DownloadCsv()
}

// Student struct
type Students struct{
	students Usecase
}

func NewController(u Usecase) *Students{
	return &Students{students: u }
}

// GetStudents
func (s *Students) GetStudentsHandler(writer http.ResponseWriter, request *http.Request) {
	students, err := s.students.GetStudentsService()
	if err != nil {
		log.Fatal(err)
		respondWithError(writer, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(writer, http.StatusOK, students)
}

// dwonload csv
func DownloadCsv(writer http.ResponseWriter, request *http.Request) {
	//usecase.
	respondWithJSON(writer, http.StatusOK, map[string]bool{"download": true})
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
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
