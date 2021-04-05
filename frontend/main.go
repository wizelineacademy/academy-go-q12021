package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"text/template"

	"main/config"
	"main/model"
)

// Abnormal exit constants
const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCSVFile
)

//
func GetTechStackList() (items []model.TechStackItem) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	resp, err := http.Get("http://localhost:8080/getTechStack")
	if err != nil {
		log.Fatalf(err.Error())
		items = []model.TechStackItem{{Title: "", Id: "", Years: ""}}
		defer resp.Body.Close()
		return
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
			panic(err)
		}
		return items
	}
}

func GetTechStackById(id string) (item model.TechStackItem) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	var url string = "http://localhost:8080/getTechStackById?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return item
}

func RenderItem(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
		return
	}
	// Casting the string number to an integer
	id := keys[0]
	item := GetTechStackById(id)

	tmpl := template.Must(template.ParseFiles("html/item.html"))

	if err := tmpl.Execute(w, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := GetMovies()
	tmpl := template.Must(template.ParseFiles("html/movies.html"))
	data := model.PageData{
		PageTitle:     "My Tech Stack",
		TechStackItem: techStackItem,
	}
	tmpl.Execute(w, data)
	WriteDataToCSVFile("result.csv", items)
}

func WriteDataToCSVFile(fileName string, items []model.TechStackItem) {
	log.Println("Data: ", items)

	csvfile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error creating file csv", err)
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

func GetMovies(queryParams model.QueryParameters) (response model.Response_All) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovies")
	if err != nil {
		log.Fatal(err.Error())
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
	log.Println(Url.String())

	if err != nil {
		defer resp.Body.Close()
		log.Fatalf(err.Error())
		var response model.Response_All
		return response
	}

	defer resp.Body.Close()

	// Print the HTTP response status.
	// fmt.Println("\n\tResponse status:", resp.Status, resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	return response
}

func GetMoviesById(id string) (response model.Response_Single) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovieById")
	if err != nil {
		log.Fatal(err.Error())
	}
	parameters := url.Values{}
	parameters.Add("id", id)
	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	return response
}

func GetQueryParams(r *http.Request) (queryParams model.QueryParameters) {
	keys := r.URL.Query()

	if val, ok := keys["item_per_workers"]; ok {
		IntItemPerWorkers, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			queryParams.ItemPerWorkers = 1
		} else {
			log.Println("item_per_workers query provided")
			queryParams.ItemPerWorkers = IntItemPerWorkers
		}
	} else {
		log.Println("item_per_workers not provided as query param")
	}

	if val, ok := keys["type"]; ok {
		queryParams.Type = val[0]
	}
	if val, ok := keys["items"]; ok {
		IntItems, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			log.Fatal(err.Error())
			queryParams.Items = 1
		} else {
			queryParams.Items = IntItems
			log.Println("items query provided: value ", IntItems)
		}
	} else {
		queryParams.Items = 1
	}
	return
}

func RenderMovies(w http.ResponseWriter, r *http.Request) {
	// Casting the string number to an integer
	queryParams := GetQueryParams(r)

	response := GetMovies(model.QueryParameters{Items: queryParams.Items, ItemPerWorkers: 1, Type: queryParams.Type})

	tmpl := template.Must(template.ParseFiles("html/movies.html"))
	data := model.Page_AllMovies{
		PageTitle: "Cine+",
		Movies:    response.Data,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderMovieById(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
		return
	}
	// Casting the string number to an integer
	id := keys[0]
	response := GetMoviesById(id)

	tmpl := template.Must(template.ParseFiles("html/item.html"))

	movieMock := model.Movie{
		ImdbTitleId:   "123123",
		Title:         "Hola  k ase",
		OriginalTitle: "Original title",
		Year:          "109123",
	}

	data := model.Page_MovieDetails{
		PageTitle: "Cine+",
		Movie:     movieMock,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, response.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var configFile string
	flag.StringVar(
		&configFile,
		"public-config-file",
		"config.yml",
		"Path to public config file",
	)
	flag.Parse()

	// Read config file
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	// Second deliverable
	http.HandleFunc("/getTechStack", GetAllItems)
	http.HandleFunc("/getTechStackById", RenderItem)

	// Third deliverable
	http.HandleFunc("/getMovies", RenderMovies)
	http.HandleFunc("/getMovieById", RenderMovieById)
	fmt.Printf("Web app running succesfully on port [%s].", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
