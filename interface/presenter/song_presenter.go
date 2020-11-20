package presenter

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type songPresenter struct {
}

type SongPresenter interface {
	ResponseSong(s *model.Song) *model.Song
}

func NewSongPresenter() SongPresenter {
	return &songPresenter{}
}

func (sp *songPresenter) ResponseSong(s *model.Song) *model.Song {
	return s
}
