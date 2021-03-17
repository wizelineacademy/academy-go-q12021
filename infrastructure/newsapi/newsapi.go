package newsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jesus-mata/academy-go-q12021/infrastructure"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
)

type NewsApiClient interface {
	GetCurrentNews() ([]dto.NewItem, error)
}

type newsApiClient struct {
	client infrastructure.HTTPClient
}

func NewApiClient(client infrastructure.HTTPClient) NewsApiClient {
	return &newsApiClient{client}
}

func (c *newsApiClient) GetCurrentNews() ([]dto.NewItem, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.currentsapi.services/v1/latest-news", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "VVaqbbrkdYyMa4Kw92mRefEDZAAxwFudy8Mhew4I-2HmQS1P")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject dto.NewsApiResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	news := responseObject.News

	return news, nil
}
