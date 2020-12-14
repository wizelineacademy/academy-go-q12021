package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/fallentemplar/goenv"
)

const baseAPIURL = "https://api.happi.dev/v1/music"

type happiService struct {
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

//HappiLyricResponse maps the response from the HAPPI API
type HappiSearchResponse struct {
	Success bool           `json:"success"`
	Length  int            `json:"length"`
	Result  []searchResult `json:"result"`
}

type searchResult struct {
	Track     string `json:"track"`
	IDTrack   int    `json:"id_track"`
	Artist    string `json:"artist"`
	IDArtist  int    `json:"id_artist"`
	Album     string `json:"album"`
	IDAlbum   int    `json:"id_album"`
	HasLyrics bool   `json:"haslyrics"`
}

//HappiService is the interface for the HappiService
type HappiService interface {
	SearchSongLyrics(artistID, albumID, trackID int) (*model.Song, error)
	SearchSongsByQuery(queryParams map[string]string) ([]*model.Song, error)
}

//NewHappiService creates a new instance of HappiService
func NewHappiService() HappiService {
	return &happiService{}
}

func (hs *happiService) SearchSongsByQuery(queryParams map[string]string) ([]*model.Song, error) {
	URL := fmt.Sprintf("%s", baseAPIURL)

	resp, err := makeHTTPCall("GET", URL, nil, queryParams)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Println(resp.Status)

	if resp.StatusCode >= 300 {
		return nil, errors.New("No song found with that query")
	}

	body, err := readResponseBody(resp, URL)

	tracks := &HappiSearchResponse{}

	err = json.Unmarshal(body, tracks)
	if err != nil {
		log.Println("Error parsing the response data from server in query ", URL)
		log.Println(err)
		return nil, err
	}

	songs := seachResponseToSongSlice(tracks)
	return songs, nil
}

func (hs *happiService) SearchSongLyrics(artistID, albumID, trackID int) (*model.Song, error) {
	log.Println("Im in SearchSongLyric")

	URL := fmt.Sprintf("%s/artists/%d/albums/%d/tracks/%d/lyrics", baseAPIURL, artistID, albumID, trackID)

	resp, err := makeHTTPCall("GET", URL, nil, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

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
func (hs *happiService) Status() int {
	return 0
}
