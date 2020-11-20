package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/registry"
	"github.com/gorilla/mux"
)

//GetSongData retrieves the information of a song by it's name
func GetSongData(w http.ResponseWriter, r *http.Request) {
	variables := r.URL.Query()
	log.Println(variables)
	queryVariable, ok1 := variables["q"]
	typeVariable, ok2 := variables["type"]

	if !ok1 || !ok2 {
		writeError(w, r, http.StatusBadRequest, "Malformed query")
		return
	}

	if typeVariable[0] != "track" {
		writeError(w, r, http.StatusNotImplemented, "Unsupported operation")
	}

	urlQuery := make(map[string]string)
	urlQuery["q"] = queryVariable[0]
	urlQuery["type"] = typeVariable[0]

	songInteractor := registry.NewSongInteractor()
	songs, err := songInteractor.GetAll(urlQuery)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "Error while processing request")
		return
	}
	jsonWritter(w, r, http.StatusFound, songs)

}

//GetSongLyrics is the handler that manages the retrieval of a song's lyrics
func GetSongLyrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	artist, err1 := strconv.Atoi(vars["artistId"])
	album, err2 := strconv.Atoi(vars["albumId"])
	track, err3 := strconv.Atoi(vars["trackId"])

	if err1 != nil || err2 != nil || err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	songInteractor := registry.NewSongInteractor()
	song := &model.Song{
		InterpreterID: artist,
		AlbumID:       album,
		ID:            track,
	}
	song, err := songInteractor.Get(song)
	if err != nil {
		if strings.Contains(err.Error(), "Song not found or contains no lyrics") {
			writeError(w, r, http.StatusNoContent, "Song not found or it contains no lyrics")
		} else {
			writeError(w, r, http.StatusFailedDependency, "Cannot connect to external server")
		}
		return
	}
	jsonWritter(w, r, http.StatusFound, song)

}
