package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-bootcamp-2020/domain/model"
	_errors "golang-bootcamp-2020/utils/error"
	"net/http"
	"net/http/httptest"
	"testing"
)

type appServiceMock struct {
	mock.Mock
}

type restError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	characters = []model.Character{
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
	character = model.Character{
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
	}
)

func (a appServiceMock) FetchData(maxPages int) ([]model.Character, _errors.RestError) {
	args := a.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]model.Character), nil
	}
	return args.Get(0).([]model.Character), args.Get(1).(_errors.RestError)
}

func (a appServiceMock) GetCharacterById(id string) (*model.Character, _errors.RestError) {
	args := a.Called()
	if args.Get(1) == nil {
		return args.Get(0).(*model.Character), nil
	}
	return args.Get(0).(*model.Character), args.Get(1).(_errors.RestError)
}

func (a appServiceMock) GetAllCharacters() ([]model.Character, _errors.RestError) {
	args := a.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]model.Character), nil
	}
	return args.Get(0).([]model.Character), args.Get(1).(_errors.RestError)
}

func (a appServiceMock) GetCharacterIdByName(name string) (string, _errors.RestError) {
	args := a.Called()
	if args.Get(1) == nil {
		return args.String(0), nil
	}
	return args.String(0), args.Get(1).(_errors.RestError)
}

func TestAppController_FetchData(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "maxPages", Value: "2"},
	}

	mockService.On("FetchData").Return(characters, nil)

	appController.FetchData(c)
	var apiResponse []model.Character
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Equal(t, characters, apiResponse)
}

func TestAppController_GetCharacter(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	mockService.On("GetCharacterById").Return(&character, nil)

	appController.GetCharacterById(c)
	var apiResponse model.Character
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Equal(t, character, apiResponse)
}

func TestAppController_GetCharacter_InvalidParam(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "id", Value: " "},
	}

	mockService.On("GetCharacterById").Return(nil, nil)

	appController.GetCharacterById(c)
	var apiResponse restError
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	errExpected := _errors.NewBadRequestError("id is required")

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, errExpected.Code(), apiResponse.Code)
	assert.EqualValues(t, errExpected.Message(), apiResponse.Message)
}

func TestAppController_GetCharacters(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "id", Value: " "},
	}

	mockService.On("GetAllCharacters").Return(characters, nil)

	appController.GetCharacters(c)
	var apiResponse []model.Character
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Equal(t, characters, apiResponse)
}

func TestAppController_GetCharacterIdByName(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "name", Value: "Rick Sanchez"},
	}

	mockService.On("GetCharacterIdByName").Return("1", nil)
	responseExpected := idResponse{"1"}

	appController.GetCharacterIdByName(c)
	var apiResponse idResponse
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Equal(t, responseExpected, apiResponse)
}

func TestAppController_GetCharacterIdByName_InvalidParam(t *testing.T) {
	var mockService = &appServiceMock{}

	appController := NewAppController(mockService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "name", Value: " "},
	}

	appController.GetCharacterIdByName(c)
	var apiResponse restError
	err := json.Unmarshal(response.Body.Bytes(), &apiResponse)

	errExpected := _errors.NewBadRequestError("name is required")

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, errExpected.Code(), apiResponse.Code)
	assert.EqualValues(t, errExpected.Message(), apiResponse.Message)
}
