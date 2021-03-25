package data

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/model/errs"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HttpSource struct {
	Data   model.HttpData
	Client HttpClient
}

func (hs HttpSource) String() string {
	return fmt.Sprintf("{Data: %v, Client: '%v'}", hs.Data, hs.Client)
}

func (hs HttpSource) GetData(httpConfig ...*model.SourceConfig) (*model.Data, error) {
	_httpConfig := hs.Data
	if len(httpConfig) > 0 {
		_httpConfig = *&httpConfig[0].HttpConfig
	}
	httpError := errs.HttpError{
		HttpData: _httpConfig,
	}

	if _httpConfig.Method == "" || _httpConfig.Url == "" {
		httpError.ErrorMessage = fmt.Sprintf("Source not defined, Method and Url are required, got %v", hs.Data)
		return &model.Data{}, httpError
	}

	req, reqError := http.NewRequest(_httpConfig.Method, _httpConfig.Url, nil)
	if reqError != nil {
		httpError.ErrorMessage = fmt.Sprintf("Request creation failed: %v", reqError)
		return &model.Data{}, httpError
	}

	response, responseError := hs.Client.Do(req)
	if responseError != nil {
		httpError.ErrorMessage = fmt.Sprintf("Error sending request: %v", responseError)
		return &model.Data{}, httpError
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		httpError.ErrorMessage = fmt.Sprintf("Error Response: %v", buf.String())
		httpError.StatusCode = response.StatusCode
		return &model.Data{}, httpError
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	responseData := model.NewHttpData(buf.String())
	return responseData, nil
}

func (hs HttpSource) SetData(data *model.Data) error {
	return nil
}

func (hs *HttpSource) NewData(data model.HttpData) {
	hs.Data = data
}
