package quotegarden

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/etyberick/golang-bootcamp-2020/entity"
)

const randomQuoteURL = "https://quote-garden.herokuapp.com/api/v3/quotes/random"

type client struct {
}

// Client interacts with quote garden
type Client interface {
	GetQuote() (qg entity.QuoteGarden, err error)
}

// NewClient returns a new instance of a client
func NewClient() Client {
	return &client{}
}

// GetQuote makes a request to https://quote-garden.herokuapp.com/api/v3/quotes/random
// and returns the response as a QuoteGarden structure.
func (c *client) GetQuote() (qg entity.QuoteGarden, err error) {
	resp, err := http.Get(randomQuoteURL)
	if err != nil {
		err = fmt.Errorf("error accessing %v - %v", randomQuoteURL, err)
		return
	}
	defer resp.Body.Close()

	//convert response to JSON
	responseMessage, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseMessage, &qg)
	return
}
