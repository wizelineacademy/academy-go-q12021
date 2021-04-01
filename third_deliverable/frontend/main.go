package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
)
type QueryParameters struct {
	ItemPerWorkers int `json:"item_per_workers"`
    Items int `json:"items"`
    Type string `json:"type"`
}

/* Movie structure */
type ShortMovie struct {
	ImdbTitleId string `json:"imdb_title_id"`
	OriginalTitle string `json:"original_title"`
	Year string `json:"year"`
	Poster string `json:"poster"`
}

type Movie struct {
	ImdbTitleId string `json:"imdb_title_id"`
    Title string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Year string `json:"year"`
	DatePublished string `json:"date_published"`
	Genre string `json:"genre"`
	Duration string `json:"duration"`
	Country string `json:"country"`
	Language string `json:"language"`
	Director string `json:"director"`
	Writer string `json:"writer"`
	ProductionCompany string `json:"production_company"`
	Actors string `json:"actors"`
	Description string `json:"description"`
	AvgVote string `json:"avg_vote"`
	Votes string `json:"votes"`
	Budget string `json:"budget"`
	UsaGrossIncome string `json:"usa_gross_income"`
	WorlwideGrossIncome string `json:"worlwide_gross_income"`
	Metascore string `json:"metascore"`
	ReviewsFromUsers string `json:"reviews_from_users"`
	ReviewsFromCritics string `json:"reviews_from_critics"`
	Poster string `json:"poster"`
}

type Response struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Results int `json:"results"`
	Data []Movie `json:"data"`
	Errors []string `json:"errors"`
	ExecutionTime string `json:"execution_time"`
}

type PageData struct {
    PageTitle string
    Movies     []Movie
}

func GetMovies(queryParams QueryParameters) (response Response) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovies")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("\n\n ITEMS:", queryParams.Items)

	parameters := url.Values{}
	// parameters.Add("type", queryParams.Type)
	itemsString := strconv.Itoa(queryParams.Items) // parse items to string
	parameters.Add("items", itemsString)
	// parameters.Add("item_per_workers", string(queryParams.ItemPerWorkers))

	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	log.Println(Url.String())

	if err != nil {
		defer resp.Body.Close()
		log.Fatalf(err.Error())
		var response Response
		return response
	} 

	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status, resp.Body)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		json.Unmarshal([]byte(scanner.Text()), &response) // items slice
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return response
	 
}

func GetMoviesById(id string) (response Response) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovieById")
	if err != nil {
		// requestErrors = append(requestErrors,err.Error())
		log.Fatal(err.Error())
	}
	parameters := url.Values{}
	parameters.Add("id", id)
	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	if err != nil {
		// requestErrors = append(requestErrors, err.Error())
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	
	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status)

	// Print the first 5 lines of the response body.
	// bufio.Reader.ReadLine
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
    	json.Unmarshal([]byte(scanner.Text()), &response) // items slice
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return response
}

func GetQueryParams(r *http.Request) (queryParams QueryParameters) {
	keys := r.URL.Query()

	// if val, ok := keys["type"]; ok {
	// 	log.Println("Type query provided")
	// 	queryParams.Type = val[0]
	// } else {
	// 	log.Println("Type not provided as query param.")
	// }
	// if val, ok := keys["item_per_workers"]; ok {
	// 	IntItemPerWorkers, err := strconv.Atoi(val[0]) // parse string to int
	// 	if err != nil {
	// 		queryParams.ItemPerWorkers = 1
	// 	} else {
	// 		log.Println("item_per_workers query provided")
	// 		queryParams.ItemPerWorkers = IntItemPerWorkers	
	// 	}
	// } else {
	// 	log.Println("item_per_workers not provided as query param")
	// }

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
	log.Println(queryParams)

	response := GetMovies(QueryParameters{Items: queryParams.Items, ItemPerWorkers: 1, Type: "odd" })
	

	tmpl := template.Must(template.ParseFiles("html/index.html"))
	data := PageData{
		PageTitle: "IMDb Movies",
		Movies: response.Data,
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
	data := PageData{
		PageTitle: "IMDb Movie",
		Movies: response.Data,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	


	if err := tmpl.Execute(w, response.Data[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}


func main() {
  http.HandleFunc("/", RenderMovies)
  http.HandleFunc("/getMovieById", RenderMovieById)
  log.Println("Server running succesfully on port 3000!")
  log.Fatal(http.ListenAndServe(":3000", nil))
}


