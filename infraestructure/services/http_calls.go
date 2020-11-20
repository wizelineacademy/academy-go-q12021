package services

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func readResponseBody(resp *http.Response, url string) ([]byte, error) {
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Println("Error reading the response data from server in query ", url)
		return nil, readErr
	}

	return body, nil
}

func makeHTTPCall(requestType string, url string, body io.Reader, queryParams map[string]string) (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(requestType, url, body)
	setHeaders(request)
	if len(queryParams) != 0 {
		setQueryValues(request, queryParams)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("Error retrieving data from server in query ", url)
		return nil, err
	}
	return resp, nil
}

func setQueryValues(request *http.Request, queryParams map[string]string) {
	query := request.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
}
