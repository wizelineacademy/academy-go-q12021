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
	GetUrlService() ([]model.Student, error)
}

// Student struct
type Students struct {
	students Usecase
}

func NewController(u Usecase) *Students {
	return &Students{students: u}
}

// GetStudentsHandler 	route: /readcsv
func (s *Students) GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.GetStudentsService()
	if err != nil {
		log.Fatal("Fail read csv", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, students)
}

// GetStudentUrlHandler		route: /storedata
func (s *Students) GetStudentUrlHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.GetUrlService()
	if err != nil {
		log.Fatal("Fail Get from url ", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	//fmt.Println("Get info from url: ", students)
	//respondWithJSON(w, http.StatusOK, map[string]bool{"download": true})
	respondWithJSON(w, http.StatusOK, students)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
