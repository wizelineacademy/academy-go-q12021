package controllers

import (
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
	s := interactor.NewSongInteractor()

}
