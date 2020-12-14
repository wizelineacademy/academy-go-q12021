package service

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

type MockDatabase struct {
	mock.Mock
}

func (mock *MockDatabase) GetAll(tableName string) ([]map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]map[string]interface{}), args.Error(1)
}
func (mock *MockDatabase) GetItemByID(tableName string, id string) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}

func (mock *MockDatabase) DeleteItem(tableName string, id string) error {
	args := mock.Called()

	return args.Error(1)
}

func (mock *MockDatabase) UpdateItem(tableName string, id string, item map[string]interface{}) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}

func (mock *MockDatabase) AddItem(tableName string, item map[string]interface{}) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}

func TestGetAllLeagues(t *testing.T) {
	mockDB := new(MockDatabase)

	league := map[string]interface{}{
		"name":              "A",
		"country":           "B",
		"current_season_id": 123,
	}

	mockDB.On("GetAll").Return([]map[string]interface{}{league}, nil)

	testService := NewLeagueService(mockDB, nil)
	result, _ := testService.GetAllLeagues()

	mockDB.AssertExpectations(t)

	assert.Equal(t, "A", result[0]["name"])
	assert.Equal(t, "B", result[0]["country"])
	assert.Equal(t, 123, result[0]["current_season_id"])
}

func TestValidateEmptyLeague(t *testing.T) {
	testService := NewLeagueService(nil, nil)
	err := testService.ValidateLeague(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The league is empty", err.Error())
}

func TestValidateEmptyLeagueName(t *testing.T) {
	league := map[string]interface{}{
		"name": "",
	}
	testService := NewLeagueService(nil, nil)
	err := testService.ValidateLeague(league)
	assert.NotNil(t, err)
	assert.Equal(t, "The league name can't be empty", err.Error())
}

func TestAddLeague(t *testing.T) {
	mockDB := new(MockDatabase)

	league := map[string]interface{}{
		"name":              "A",
		"country":           "B",
		"current_season_id": 123,
	}

	mockDB.On("AddItem").Return(league, nil)

	testService := NewLeagueService(mockDB, nil)
	result, _ := testService.Addleague(league)

	mockDB.AssertExpectations(t)

	assert.Equal(t, "A", result["name"])
	assert.Equal(t, "B", result["country"])
	assert.Equal(t, 123, result["current_season_id"])
}
