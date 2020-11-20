package services

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/fallentemplar/goenv"
)

const apiURL = "https://api.happi.dev/v1/music/"

type happiService struct {
	c chan model.Song
}

type HappiService interface {
	SearchSongLyrics(songID string) (*model.Song, error)
	SearchSongsByArtist(artistID string) ([]*model.Song, error)
	SearchArtist() // "songs/"
}

func NewHappiService() HappiService {
	c := make(chan model.Song)
	return &happiService{c}
}

//SearchArtist calls the Happi API and retrieves the artist song
func (hs *happiService) SearchArtist() {
	//TODO: Implementar búsqueda en BD/API aquí

	resp, err := http.Get(fmt.Sprintf("%s", apiURL))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Print(string(body))
}

func (hs *happiService) SearchSongsByArtist(artistID string) ([]*model.Song, error) {
	return nil, nil
}

func (hs *happiService) SearchSongLyrics(songID string) (*model.Song, error) {
	return nil, nil
}

func setHeaders(req *http.Request) {
	req.Header.Set("x-happi-key", goenv.GetString("HAPPI_API_KEY", ""))
}

func (hs *happiService) Start() error {
	return nil
}

func (hs *happiService) Stop() error {
	return nil
}
func (hs *happiService) Status() error {
	return nil
}
