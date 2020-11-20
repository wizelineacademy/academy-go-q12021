package repository

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type SongRepository interface {
	Find(song *model.Song) (*model.Song, error)
	FindByPattern(queryParams map[string]string) ([]*model.Song, error)
}
