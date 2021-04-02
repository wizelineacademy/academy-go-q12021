package newsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jesus-mata/academy-go-q12021/infrastructure"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
)

//go:generate mockgen -package mocks -destination $ROOTDIR/mocks/$GOPACKAGE/mock_$GOFILE . NewsApiClient
type NewsApiClient interface {
	GetCurrentNews() ([]dto.NewItem, error)
}

type newsApiClient struct {
	host   string
	apiKey string
	client infrastructure.HTTPClient
}

func NewApiClient(host string, apiKey string, client infrastructure.HTTPClient) NewsApiClient {
	return &newsApiClient{host, apiKey, client}
}

func (c *newsApiClient) GetCurrentNews() ([]dto.NewItem, error) {
	req, err := http.NewRequest(http.MethodGet, c.host, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.apiKey)

	client := c.client
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error on API request %s", string(responseData)))
	}

	var responseObject dto.NewsApiResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	news := responseObject.News

	return news, nil
}
