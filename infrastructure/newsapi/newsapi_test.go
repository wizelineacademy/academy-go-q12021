package newsapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/jesus-mata/academy-go-q12021/mocks/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentNews(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHttpClient := mocks.NewMockHTTPClient(ctrl)

	t.Log("Running test")

	// build response JSON
	json := `{"status":"ok","news":[{"id":"e1749cf0-8a49-4729-88b2-e5b4d03464ce","title":"US House speaker Nancy Pelosi backs congressional legislation on Hong Kong","description":"US House speaker Nancy Pelosi on Wednesday threw her support behind legislation meant to back Hong Kong's anti-government protesters.Speaking at a news conference featuring Hong Kong activists Joshua Wong Chi-fung and Denise Ho, who testified before the Congressional-Executive Commission on China (C...","url":"https://www.scmp.com/news/china/politics/article/3027994/us-house-speaker-nancy-pelosi-backs-congressional-legislation","author":"Robert Delaney","image":"None","language":"en","category":["world"],"published":"2019-09-18 21:08:58 +0000"}]}`
	// create a new reader with that JSON
	respBody := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mockHttpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       respBody,
	}, nil)

	newApi := NewApiClient("http://test", "tet_api_key", mockHttpClient)

	data, err := newApi.GetCurrentNews()
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	assert.Equal(t, len(data), 1)
}

func TestGetCurrentFail(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHttpClient := mocks.NewMockHTTPClient(ctrl)

	t.Log("Running test")

	// build response JSON
	json := `Authorized Required`
	// create a new reader with that JSON
	respBody := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mockHttpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 401,
		Body:       respBody,
	}, nil)

	newApi := NewApiClient("http://test", "tet_api_key", mockHttpClient)

	_, err := newApi.GetCurrentNews()
	if err != nil {
		t.Logf("Test ok: %s", err)
	}
}

func TestGetCurrentRequestTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHttpClient := mocks.NewMockHTTPClient(ctrl)

	t.Log("Running test")

	// build response JSON
	json := `Timeout`
	// create a new reader with that JSON
	respBody := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mockHttpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 408,
		Body:       respBody,
	}, errors.New("Timeout Error"))

	newApi := NewApiClient("http://test", "tet_api_key", mockHttpClient)

	_, err := newApi.GetCurrentNews()
	if err != nil && strings.Contains(err.Error(), "Timeout") {
		t.Logf("Timeout test ok: %s", err)
	}
}

func TestGetCurrentBadResponseFormat(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHttpClient := mocks.NewMockHTTPClient(ctrl)

	t.Log("Running test")

	// build response JSON
	json := `Timeout`
	// create a new reader with that JSON
	respBody := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mockHttpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       respBody,
	}, nil)

	newApi := NewApiClient("http://test", "tet_api_key", mockHttpClient)

	_, err := newApi.GetCurrentNews()
	if err != nil && strings.Contains(err.Error(), "invalid character") {
		t.Logf("Timeout test ok: %s", err)
	}
}
