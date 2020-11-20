package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/fallentemplar/goenv"
)

const apiURL = "https://api.happi.dev/v1/music"

type happiService struct {
	c chan model.Song
}

//HappiLyricResponse maps the response from the HAPPI API
type HappiLyricResponse struct {
	Success bool        `json:"success"`
	Length  int         `json:"length"`
	Result  lyricResult `json:"result"`
}

type lyricResult struct {
	Artist          string `json:"artist"`
	IDArtist        int    `json:"id_artist"`
	Track           string `json:"track"`
	IDTrack         int    `json:"id_track"`
	IDAlbum         int    `json:"id_album"`
	Album           string `json:"album"`
	Lyrics          string `json:"lyrics"`
	APIArtist       string `json:"api_artist"`
	APIAlbums       string `json:"api_albums"`
	APIAlbum        string `json:"api_album"`
	APITracks       string `json:"api_tracks"`
	APITrack        string `json:"api_track"`
	APILyrics       string `json:"api_lyrics"`
	CopyrightLabel  string `json:"copyright_label"`
	CopyrightNotice string `json:"copyright_notice"`
	CopyrightString string `json:"copyright_text"`
}

//HappiService is the interface for the HappiService
type HappiService interface {
	SearchSongLyrics(artistID, albumID, trackID int) (*model.Song, error)
	SearchSongsByArtist(artistID string) ([]*model.Song, error)
	SearchArtist() // "songs/"
}

//NewHappiService creates a new instance of HappiService
func NewHappiService() HappiService {
	c := make(chan model.Song)
	return &happiService{c}
}

//SearchArtist calls the Happi API and retrieves the artist song
func (hs *happiService) SearchArtist() {
	//TODO: Implementar búsqueda en BD/API aquí

	resp, err := http.Get(fmt.Sprintf("%sartists/%d/albums/%d/tracks/%d/lyrics", apiURL))
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

func (hs *happiService) SearchSongLyrics(artistID, albumID, trackID int) (*model.Song, error) {
	log.Println("Im in SearchSongLyric")

	URL := fmt.Sprintf("%s/artists/%d/albums/%d/tracks/%d/lyrics", apiURL, artistID, albumID, trackID)

	resp, err := makeHTTPCall("GET", URL, nil)
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, errors.New("Song not found or contains no lyrics")
	}

	body, err := readResponseBody(resp, URL)

	lyric := &HappiLyricResponse{}

	err = json.Unmarshal(body, lyric)
	if err != nil {
		log.Println("Error parsing the response data from server in query ", URL)
		log.Println(err)
		return nil, err
	}
	song := lyricResponseToSong(lyric)

	return song, nil
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
