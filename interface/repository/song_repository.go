package repository

import (
	"log"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type songRepository struct {
	db services.HappiService
}

//NewSongRepository creates a new Song Repository
func NewSongRepository() repository.SongRepository {
	happiService := services.NewHappiService()
	return &songRepository{happiService}
}

func (sr *songRepository) Find(song *model.Song) (*model.Song, error) {
	log.Println("Here in song_repository.Find")
	song, err := sr.db.SearchSongLyrics(song.InterpreterID, song.AlbumID, song.ID) //TODO: Make this dynamic
	if err != nil {
		log.Println(err)
		return nil, err
	}
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
