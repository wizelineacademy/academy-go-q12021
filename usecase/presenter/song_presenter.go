package presenter

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type SongPresenter interface {
	ResponseSong(u *model.Song) (*model.Song, error)
}
