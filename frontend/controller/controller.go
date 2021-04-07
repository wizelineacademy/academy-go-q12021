package controller

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func WriteDataToCSVFile(fileName string, items []model.TechStackItem) {
	csvfile, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file csv", err.Error())

		return
	}
	var writter *csv.Writer = csv.NewWriter(csvfile)

	for _, item := range items {
		strSlice := []string{item.Id, item.Title, item.Years}
		fmt.Println(strSlice)
		writter.Write(strSlice)
	}
	// Write any buffered items data to the underlying writer (standard output).
	writter.Flush()

	if err := writter.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

//
func GetItems() (items []model.TechStackItem, err error) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	resp, err := http.Get("http://localhost:8080/getTechStack")
	if err != nil {
		log.Println("Error while getting a server response")
		return nil, err
	} else {
		defer resp.Body.Close()
		// Print the HTTP response status.
		fmt.Println("\n\tResponse status:", resp.Status)

		// Print the first 5 lines of the response body.
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			json.Unmarshal([]byte(scanner.Text()), &items) // items slice
		}
		if err := scanner.Err(); err != nil {
			log.Println(err.Error())
		}
		return items, err
	}
}

func GetTechStackItems(id string) (item model.TechStackItem, err error) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	var url string = "http://localhost:8080/getTechStackById?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		json.Unmarshal([]byte(scanner.Text()), &item) // items slice
	}
	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	return item, err
}

func GetMovies(queryParams model.QueryParameters, endpoint string) (response model.Response_All, err error) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/" + endpoint)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("\n\n TYPE:", queryParams.Type)

	parameters := url.Values{}
	parameters.Add("type", queryParams.Type)
	itemsString := strconv.Itoa(queryParams.Items) // parse items to string
	parameters.Add("items", itemsString)
	parameters.Add("item_per_workers", string(rune(queryParams.ItemPerWorkers)))

	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	// Print the HTTP response status.
	// fmt.Println("\n\tResponse status:", resp.Status, resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func GetMoviesById(id string) (response model.Response_Single, err error) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovieById")
	if err != nil {
		log.Println(err.Error())
		return
	}
	parameters := url.Values{}
	parameters.Add("id", id)
	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
	}
	return
}
