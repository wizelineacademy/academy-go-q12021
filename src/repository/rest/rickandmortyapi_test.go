package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-bootcamp-2020/domain/model"
	"golang-bootcamp-2020/repository/rest/testdata"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	name1        = "Rick Sanchez"
	image1       = "https://rickandmortyapi.com/api/character/avatar/1.jpeg"
	name2        = "Morty Smith"
	image2       = "https://rickandmortyapi.com/api/character/avatar/2.jpeg"
	episode1     = "https://rickandmortyapi.com/api/episode/1"
	episode2     = "https://rickandmortyapi.com/api/episode/2"
	status       = "Alive"
	species      = "Human"
	cType        = ""
	gender       = "Male"
	originName   = "Earth (C-137)"
	originURL    = "https://rickandmortyapi.com/api/location/1"
	locationName = "Earth (Replacement Dimension)"
	locationURL  = "https://rickandmortyapi.com/api/location/20"
)

func TestRickAndMortyAPI_FetchData_SinglePage(t *testing.T) {
	restRepo := NewRickAndMortyAPIRepository(resty.New())

	var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testdata.SinglePage))
	}))
	defer mockServer.Close()

	expectedCharacters := []model.Character{
		{
			ID:       1,
			Name:     name1,
			Status:   status,
			Species:  species,
			Type:     cType,
			Gender:   gender,
			Origin:   model.Nested{Name: originName, URL: originURL},
			Location: model.Nested{Name: locationName, URL: locationURL},
			Image:    image1,
			Episodes: []string{episode1, episode2},
		},
		{
			ID:       2,
			Name:     name2,
			Status:   status,
			Species:  species,
			Type:     cType,
			Gender:   gender,
			Origin:   model.Nested{Name: originName, URL: originURL},
			Location: model.Nested{Name: locationName, URL: locationURL},
			Image:    image2,
			Episodes: []string{episode1, episode2},
		},
	}

	apiCharacters = mockServer.URL

	characters, err := restRepo.FetchData(0)

	assert.NotNil(t, characters)
	assert.Nil(t, err)
	assert.Equal(t, expectedCharacters, characters)
}

func TestRickAndMortyAPI_FetchData_MultiplePage(t *testing.T) {
	restRepo := NewRickAndMortyAPIRepository(resty.New())

	var mockServerPage2 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testdata.Page2))
	}))
	defer mockServerPage2.Close()

	var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(testdata.Page1, mockServerPage2.URL)))
	}))
	defer mockServer.Close()

	expectedCharacters := []model.Character{
		{
			ID:       1,
			Name:     name1,
			Status:   status,
			Species:  species,
			Type:     cType,
			Gender:   gender,
			Origin:   model.Nested{Name: originName, URL: originURL},
			Location: model.Nested{Name: locationName, URL: locationURL},
			Image:    image1,
			Episodes: []string{episode1, episode2},
		},
		{
			ID:       2,
			Name:     name2,
			Status:   status,
			Species:  species,
			Type:     cType,
			Gender:   gender,
			Origin:   model.Nested{Name: originName, URL: originURL},
			Location: model.Nested{Name: locationName, URL: locationURL},
			Image:    image2,
			Episodes: []string{episode1, episode2},
		},
	}

	apiCharacters = mockServer.URL

	characters, err := restRepo.FetchData(2)

	assert.NotNil(t, characters)
	assert.Nil(t, err)
	assert.Equal(t, expectedCharacters, characters)
}

func TestRickAndMortyAPI_FetchData_Error(t *testing.T) {
	restRepo := NewRickAndMortyAPIRepository(resty.New())

	var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer mockServer.Close()

	apiCharacters = mockServer.URL

	characters, err := restRepo.FetchData(0)

	assert.Nil(t, characters)
	assert.NotNil(t, err)
	assert.EqualValues(t, 500, err.Code())
	assert.EqualValues(t, "error in rest response", err.Message())
}
