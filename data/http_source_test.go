package data

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/grethelBello/academy-go-q12021/model"
)

type mockResponseData struct {
	Body          string
	StatusCode    int
	ErrorResponse string
}
type httpClientMock struct{}

func (hcm httpClientMock) Do(request *http.Request) (*http.Response, error) {
	return doMock(request)
}

var doMock func(*http.Request) (*http.Response, error)

func TestSourceNotDefinedError(t *testing.T) {
	source := initSource("", "", mockResponseData{})

	_, notDefinedSourceError := source.GetData()
	if notDefinedSourceError == nil {
		t.Errorf("Source should return an error when there are not required fields defined: %v", source)
	} else if !strings.Contains(notDefinedSourceError.Error(), "Source not defined") {
		t.Errorf("Source should return an error when there are not required fields defined: %v", notDefinedSourceError)
	}
}
func TestCreateRequestError(t *testing.T) {
	source := initSource("GE", "https://my-url.com", mockResponseData{})

	_, createRequestError := source.GetData()
	if createRequestError == nil {
		t.Errorf("Source should return an error when there are not required fields defined: %v", source)
	} else if !strings.Contains(createRequestError.Error(), "Error Response") {
		// TODO: The error string should be 'Error sending request'
		t.Errorf("Source should return an error when there are not required fields defined: %v", createRequestError)
	}
}
func TestHttpSourceErrorResponse(t *testing.T) {
	responseMock := mockResponseData{
		Body:       "Testing",
		StatusCode: http.StatusNotFound,
	}
	source := initSource("GET", "https://my-url.com", responseMock)

	response, errorResponse := source.GetData()
	if errorResponse == nil {
		t.Errorf("Source should return an error when the response is different of 200: %v", response)
	} else if !strings.Contains(errorResponse.Error(), "Error Response") {
		t.Errorf("Source should return an error when there are not required fields defined: %v", errorResponse)
	}
}
func TestHttpSourceSuccessResponse(t *testing.T) {
	const testBody = "Testing"
	responseMock := mockResponseData{
		Body:       testBody,
		StatusCode: http.StatusOK,
	}
	source := initSource("GET", "https://my-url.com", responseMock)

	httpData, errorResponse := source.GetData()
	if errorResponse != nil {
		t.Errorf("Source should not return an error when the response status code is 200: %v", errorResponse)
	} else if httpData.HttpData != testBody {
		t.Errorf("Body response expected '%v', got '%v'", testBody, httpData.HttpData)
	}
}

func initSource(method, url string, response mockResponseData) *HttpSource {
	doMock = func(req *http.Request) (*http.Response, error) {
		if response.ErrorResponse != "" {
			return nil, errors.New(response.ErrorResponse)
		}

		bodyMock := ioutil.NopCloser(bytes.NewReader([]byte(response.Body)))
		return &http.Response{
			StatusCode: response.StatusCode,
			Body:       bodyMock,
		}, nil
	}

	clientMock := httpClientMock{}
	source := &HttpSource{
		Data: model.HttpData{
			Url:    url,
			Method: method,
		},
		Client: clientMock,
	}

	return source
}
