package repository

import (
	"log"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type songRepository struct {
	db     services.HappiService
	logger services.Logger
}

//NewSongRepository creates a new Song Repository
func NewSongRepository(songAPI services.HappiService, logger services.Logger) repository.SongRepository {
	return &songRepository{songAPI, logger}
}

func (sr *songRepository) Find(song *model.Song) (*model.Song, error) {
	song, err := sr.db.SearchSongLyrics(song.InterpreterID, song.AlbumID, song.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sr.logger.Append(song.ToString())
	return song, nil
}

func (sr *songRepository) FindByPattern(queryParams map[string]string) ([]*model.Song, error) {
	log.Println("Here in song_repository.FindByPattern")
	song, err := sr.db.SearchSongsByQuery(queryParams) //TODO: Make this dynamic
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return song, nil
}
