package repository

import (
	"bootcamp/domain/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const BaseURL = "https://icanhazdadjoke.com/"

// jokeClient struct for REST client
type JokeClient struct {
	BaseURL string
	HTTPClient *http.Client
}

// NewJokeClient exported constructor
func NewJokeClient() *JokeClient {
	return &JokeClient{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// GetJoke exported function to get al jokes from third party REST endpoint
func (c *JokeClient) GetJoke() ([]*model.Joke, error) {

	req, err := http.NewRequest("GET", BaseURL, nil)

	if err != nil {
		return nil, err
	}

	jokes, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return jokes, nil

}

// sendRequest sends and handles response
func (c *JokeClient) sendRequest(req *http.Request) ([]*model.Joke, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK  {
		return nil, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	joke := model.Joke{}
	var jokes []*model.Joke

	err = json.Unmarshal(respBody, &joke)
	if err != nil {
		return nil, err
	}

	jokes = append(jokes, &joke)

	return jokes, nil
}
