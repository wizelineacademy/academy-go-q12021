package client

import (
	"academy/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ConsultExternalService will call the Jokes API
func ConsultExternalService() []model.Joke {
	url := "https://official-joke-api.appspot.com/random_ten"
	request, error := http.NewRequest("GET", url, nil)
	if error != nil {
		fmt.Println(error)
	}
	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	// Unmarshal JSON data
	var jsonData []model.Joke
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		fmt.Println("Unable to handle json")
	}
	return jsonData
}
