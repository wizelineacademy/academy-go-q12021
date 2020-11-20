package repository

import (
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
	sr.db.SearchArtist()
	// err := ur.db.Find()
	return nil, nil
}
