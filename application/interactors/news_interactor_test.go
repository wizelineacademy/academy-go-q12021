package interactors

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jesus-mata/academy-go-q12021/domain"
	mocks "github.com/jesus-mata/academy-go-q12021/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	newsRepository := mocks.NewMockNewsArticleRepository(ctrl)

	news1, _ := domain.CreateNewsArticle("id1", "title 1", "description 1", "url", "author 1", "image", "language", "category", time.Now())
	news2, _ := domain.CreateNewsArticle("id2", "title 2", "description 2", "url", "author 2", "image", "language", "category", time.Now())
	arr := []*domain.NewsArticle{news1, news2}
	newsRepository.EXPECT().FindAll().Return(arr, nil)

	newsInteractor := NewNewsArticlesInteractor(newsRepository)

	_, err := newsInteractor.GetAll()
	if err != nil {
		t.Errorf("Test failed due to: %s", err)
	}

	for i, v := range arr {
		assert.Equal(t, fmt.Sprintf("id%v", i+1), v.Id)
		assert.Equal(t, fmt.Sprintf("title %v", i+1), v.Title)
	}
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	newsRepository := mocks.NewMockNewsArticleRepository(ctrl)

	news1, _ := domain.CreateNewsArticle("id1", "title 1", "description 1", "url", "author 1", "image", "language", "category", time.Now())

	newsRepository.EXPECT().FindByID("id1").Return(news1, nil)

	newsInteractor := NewNewsArticlesInteractor(newsRepository)

	res, err := newsInteractor.GetByID("id1")
	if err != nil {
		t.Errorf("Test failed due to: %s", err)
	}

	assert.Equal(t, "id1", res.Id)
	assert.Equal(t, "title 1", res.Title)

}

func TestFetchAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	newsRepository := mocks.NewMockNewsArticleRepository(ctrl)

	newsRepository.EXPECT().FetchCurrent().Return(nil)

	newsInteractor := NewNewsArticlesInteractor(newsRepository)

	err := newsInteractor.FetchAll()
	if err != nil {
		t.Errorf("Test failed due to: %s", err)
	}

}
