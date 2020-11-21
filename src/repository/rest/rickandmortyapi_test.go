package rest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-bootcamp-2020/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRickAndMortyApi_FetchData_SinglePage(t *testing.T) {
	restRepo := NewRickAndMortyApiRepository()

	var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    "info": {
        "count": 671,
        "pages": 34,
        "next": "https://rickandmortyapi.com/api/character/?page=2",
        "prev": null
    },
    "results": [
        {
            "id": 1,
            "name": "Rick Sanchez",
            "status": "Alive",
            "species": "Human",
            "type": "",
            "gender": "Male",
            "origin": {
                "name": "Earth (C-137)",
                "url": "https://rickandmortyapi.com/api/location/1"
            },
            "location": {
                "name": "Earth (Replacement Dimension)",
                "url": "https://rickandmortyapi.com/api/location/20"
            },
            "image": "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
            "episode": [
                "https://rickandmortyapi.com/api/episode/1",
                "https://rickandmortyapi.com/api/episode/2"
            ],
            "url": "https://rickandmortyapi.com/api/character/1",
            "created": "2017-11-04T18:48:46.250Z"
        },
        {
            "id": 2,
            "name": "Morty Smith",
            "status": "Alive",
            "species": "Human",
            "type": "",
            "gender": "Male",
            "origin": {
                "name": "Earth (C-137)",
                "url": "https://rickandmortyapi.com/api/location/1"
            },
            "location": {
                "name": "Earth (Replacement Dimension)",
                "url": "https://rickandmortyapi.com/api/location/20"
            },
            "image": "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
            "episode": [
                "https://rickandmortyapi.com/api/episode/1",
                "https://rickandmortyapi.com/api/episode/2"
            ],
            "url": "https://rickandmortyapi.com/api/character/2",
            "created": "2017-11-04T18:50:21.651Z"
        }
    ]}`))
	}))
	defer mockServer.Close()

	expectedCharacters := []model.Character{
		{
			Id:       1,
			Name:     "Rick Sanchez",
			Status:   "Alive",
			Species:  "Human",
			Type:     "",
			Gender:   "Male",
			Origin:   model.Nested{Name: "Earth (C-137)", Url: "https://rickandmortyapi.com/api/location/1"},
			Location: model.Nested{Name: "Earth (Replacement Dimension)", Url: "https://rickandmortyapi.com/api/location/20"},
			Image:    "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
			Episodes: []string{"https://rickandmortyapi.com/api/episode/1", "https://rickandmortyapi.com/api/episode/2"},
		},
		{
			Id:       2,
			Name:     "Morty Smith",
			Status:   "Alive",
			Species:  "Human",
			Type:     "",
			Gender:   "Male",
			Origin:   model.Nested{Name: "Earth (C-137)", Url: "https://rickandmortyapi.com/api/location/1"},
			Location: model.Nested{Name: "Earth (Replacement Dimension)", Url: "https://rickandmortyapi.com/api/location/20"},
			Image:    "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
			Episodes: []string{"https://rickandmortyapi.com/api/episode/1", "https://rickandmortyapi.com/api/episode/2"},
		},
	}

	apiCharacters = mockServer.URL

	characters, err := restRepo.FetchData(0)

	assert.NotNil(t, characters)
	assert.Nil(t, err)
	assert.Equal(t, expectedCharacters, characters)
}

func TestRickAndMortyApi_FetchData_MultiplePage(t *testing.T) {
	restRepo := NewRickAndMortyApiRepository()

	var mockServerPage2 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
    "info": {
        "count": 671,
        "pages": 34,
        "next": "https://rickandmortyapi.com/api/character/?page=2",
        "prev": null
    },
    "results": [
        {
            "id": 2,
            "name": "Morty Smith",
            "status": "Alive",
            "species": "Human",
            "type": "",
            "gender": "Male",
            "origin": {
                "name": "Earth (C-137)",
                "url": "https://rickandmortyapi.com/api/location/1"
            },
            "location": {
                "name": "Earth (Replacement Dimension)",
                "url": "https://rickandmortyapi.com/api/location/20"
            },
            "image": "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
            "episode": [
                "https://rickandmortyapi.com/api/episode/1",
                "https://rickandmortyapi.com/api/episode/2"
            ],
            "url": "https://rickandmortyapi.com/api/character/2",
            "created": "2017-11-04T18:50:21.651Z"
        }
    ]}`))
	}))
	defer mockServerPage2.Close()

	var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{
    "info": {
        "count": 671,
        "pages": 34,
        "next": "%s",
        "prev": null
    },
    "results": [
        {
            "id": 1,
            "name": "Rick Sanchez",
            "status": "Alive",
            "species": "Human",
            "type": "",
            "gender": "Male",
            "origin": {
                "name": "Earth (C-137)",
                "url": "https://rickandmortyapi.com/api/location/1"
            },
            "location": {
                "name": "Earth (Replacement Dimension)",
                "url": "https://rickandmortyapi.com/api/location/20"
            },
            "image": "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
            "episode": [
                "https://rickandmortyapi.com/api/episode/1",
                "https://rickandmortyapi.com/api/episode/2"
            ],
            "url": "https://rickandmortyapi.com/api/character/1",
            "created": "2017-11-04T18:48:46.250Z"
        }
    ]}`, mockServerPage2.URL)))
	}))
	defer mockServer.Close()

	expectedCharacters := []model.Character{
		{
			Id:       1,
			Name:     "Rick Sanchez",
			Status:   "Alive",
			Species:  "Human",
			Type:     "",
			Gender:   "Male",
			Origin:   model.Nested{Name: "Earth (C-137)", Url: "https://rickandmortyapi.com/api/location/1"},
			Location: model.Nested{Name: "Earth (Replacement Dimension)", Url: "https://rickandmortyapi.com/api/location/20"},
			Image:    "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
			Episodes: []string{"https://rickandmortyapi.com/api/episode/1", "https://rickandmortyapi.com/api/episode/2"},
		},
		{
			Id:       2,
			Name:     "Morty Smith",
			Status:   "Alive",
			Species:  "Human",
			Type:     "",
			Gender:   "Male",
			Origin:   model.Nested{Name: "Earth (C-137)", Url: "https://rickandmortyapi.com/api/location/1"},
			Location: model.Nested{Name: "Earth (Replacement Dimension)", Url: "https://rickandmortyapi.com/api/location/20"},
			Image:    "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
			Episodes: []string{"https://rickandmortyapi.com/api/episode/1", "https://rickandmortyapi.com/api/episode/2"},
		},
	}

	apiCharacters = mockServer.URL

	characters, err := restRepo.FetchData(2)

	assert.NotNil(t, characters)
	assert.Nil(t, err)
	assert.Equal(t, expectedCharacters, characters)
}

func TestRickAndMortyApi_FetchData_Error(t *testing.T) {
	restRepo := NewRickAndMortyApiRepository()

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
