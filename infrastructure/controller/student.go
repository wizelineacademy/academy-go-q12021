package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-bootcamp-2020/domain/model"
)

// Usecase interface
type Usecase interface {
	GetStudentsService() ([]model.Student, error)
	GetURLService() ([]model.Student, error)
}

// Students Use case struct
type Students struct {
	students Usecase
}

// NewController
func NewController(u Usecase) *Students {
	return &Students{students: u}
}

// GetStudentsHandler 	Handler for: /readcsv
func (s *Students) GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.GetStudentsService()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		respondWithJSON(w, http.StatusOK, students)
	}
}

// GetStudentUrlHandler	 Handler for: /storedata
func (s *Students) GetStudentURLHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.GetURLService()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		total := strconv.Itoa(len(students))
		respondWithJSON(w, http.StatusOK, map[string]string{"ok": "true", "msj": "Csv created", "total": total})
	}
}

// respondWithError response with error code and message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"msj": msg, "ok": "false"})
}

// respondWithJSON  respond message in JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
