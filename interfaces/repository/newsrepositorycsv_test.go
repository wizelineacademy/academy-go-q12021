package repository

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
	infra "github.com/jesus-mata/academy-go-q12021/mocks/infrastructure"
	newsapi "github.com/jesus-mata/academy-go-q12021/mocks/newsapi"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	csvSource := infra.NewMockCsvSource(ctrl)

	api := newsapi.NewMockNewsApiClient(ctrl)

	uuid := "687a455e-b866-48f8-bc2d-85aff2bad30c"

	data := [][]string{
		{"f5433032-6d9b-4c28-a353-1517adb2fecd", "Title 1", "Descritpion 1", "http://page/page.com", "Author 1", "https://storage/image1.jpeg", "en", "technology", "2021-03-17 15:57:47 +0000"},
		{uuid, "Title 2", "Descritpion 2", "http://page/page.com", "Author 2", "https://storage/image1.jpeg", "en", "general", "2021-03-17 15:57:47 +0000"},
	}

	csvSource.EXPECT().GetAllLines().Return(data, nil)

	repo := NewNewsArticleRepository(csvSource, api, log.Default())
	newsArticle, err := repo.FindByID(uuid)
	if err != nil {
		t.Errorf("Test failed due to %s", err)
	}

	assert.Equal(t, newsArticle.Id, uuid)
	assert.Equal(t, newsArticle.Title, "Title 2")

}

func TestFindByIDNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	csvSource := infra.NewMockCsvSource(ctrl)

	api := newsapi.NewMockNewsApiClient(ctrl)

	uuid := "687a455e-b866-48f8-bc2d-85aff2bad30c"

	data := [][]string{
		{"f5433032-6d9b-4c28-a353-1517adb2fecd", "Title 1", "Descritpion 1", "http://page/page.com", "Author 1", "https://storage/image1.jpeg", "en", "technology", "2021-03-17 15:57:47 +0000"},
		{"d5a0fa3e-24a7-461f-9236-cf59894bfea7", "Title 2", "Descritpion 2", "http://page/page.com", "Author 2", "https://storage/image1.jpeg", "en", "general", "2021-03-17 15:57:47 +0000"},
	}

	errorMsg := fmt.Sprintf("News Article with ID %v does not exist.", uuid)

	csvSource.EXPECT().GetAllLines().Return(data, nil)

	repo := NewNewsArticleRepository(csvSource, api, log.Default())
	_, err := repo.FindByID(uuid)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errorMsg)

}

func TestFindByIDNotFounds(t *testing.T) {

	ctrl := gomock.NewController(t)
	csvSource := infra.NewMockCsvSource(ctrl)

	api := newsapi.NewMockNewsApiClient(ctrl)

	uuid := "687a455e-b866-48f8-bc2d-85aff2bad30c"

	csvSource.EXPECT().GetAllLines().Return(nil, errors.New("Cannot parse CSV"))

	repo := NewNewsArticleRepository(csvSource, api, log.Default())
	_, err := repo.FindByID(uuid)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Cannot parse CSV")
}

func TestFindAll(t *testing.T) {

	ctrl := gomock.NewController(t)
	csvSource := infra.NewMockCsvSource(ctrl)

	api := newsapi.NewMockNewsApiClient(ctrl)

	uuids := []string{"f5433032-6d9b-4c28-a353-1517adb2fecd", "687a455e-b866-48f8-bc2d-85aff2bad30c"}

	data := [][]string{
		{uuids[0], "Title 1", "Descritpion 1", "http://page/page.com", "Author 1", "https://storage/image1.jpeg", "en", "technology", "2021-03-17 15:57:47 +0000"},
		{uuids[1], "Title 2", "Descritpion 2", "http://page/page.com", "Author 2", "https://storage/image1.jpeg", "en", "general", "2021-03-17 15:57:47 +0000"},
	}

	csvSource.EXPECT().GetAllLines().Return(data, nil)

	repo := NewNewsArticleRepository(csvSource, api, log.Default())
	newsArticles, err := repo.FindAll()
	if err != nil {
		t.Errorf("Test failed due to %s", err)
	}

	for i, v := range newsArticles {
		assert.Equal(t, v.Id, uuids[i])
		assert.Equal(t, v.Title, fmt.Sprintf("Title %v", i+1))
	}

}

func TestFetchCurrent(t *testing.T) {

	ctrl := gomock.NewController(t)
	csvSource := infra.NewMockCsvSource(ctrl)

	api := newsapi.NewMockNewsApiClient(ctrl)

	data := []dto.NewItem{{Id: "f5433032-6d9b-4c28-a353-1517adb2fecd", Title: "Title 1", Description: "Description 1", Url: "http://url.com", Author: "Some Ahthor", Image: "htpp://storage/image.png", Language: "en", Category: []string{"general"}, Published: "2021-03-17 15:57:47 +0000"}}
	api.EXPECT().GetCurrentNews().Return(data, nil)

	csvSource.EXPECT().WriteLines(data)

	repo := NewNewsArticleRepository(csvSource, api, log.Default())
	err := repo.FetchCurrent()

	assert.Nil(t, err, "Fetch Current news was not successful")
}
