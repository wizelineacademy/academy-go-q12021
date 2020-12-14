package interactor_test

import (
	"fmt"
	"testing"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var expectedSong *model.Song = &model.Song{
	InterpreterID: 22635,
	Interpreter:   "Mägo de Oz",
	AlbumID:       1455876,
	Album:         "Fiesta Pagana 2.0",
	ID:            14519164,
	Name:          "Fiesta Pagana 2.0",
	Length:        156,
}

type SongRepositoryMock struct {
	mock.Mock
}

func (sr *SongRepositoryMock) Find(song *model.Song) (*model.Song, error) {
	sr.Called(song)
	return expectedSong, nil
}
func (sr *SongRepositoryMock) FindByPattern(queryParams map[string]string) ([]*model.Song, error) {
	return nil, nil
}

func (sp *SongPresenterMock) ResponseSong(u *model.Song) *model.Song { return nil }

type SongPresenterMock struct {
	mock.Mock
}

func TestGetSongLyrics(t *testing.T) {

	songRepository := new(SongRepositoryMock)
	songPresenter := new(SongPresenterMock)

	requestedSong := &model.Song{
		InterpreterID: 22635,
		AlbumID:       1455876,
		ID:            14519164,
	}

	songRepository.On("Find", requestedSong).Return(expectedSong)

	songInteractor := interactor.NewSongInteractor(songRepository, songPresenter)

	foundSong, err := songInteractor.Get(requestedSong)
	fmt.Println("Is nil? ", foundSong == nil)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedSong, foundSong)
}
