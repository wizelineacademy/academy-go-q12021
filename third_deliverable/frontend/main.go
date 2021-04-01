package main

import (
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

type Response_All struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Results int `json:"results"`
	Data []ShortMovie `json:"data"`
	Errors []string `json:"errors"`
	ExecutionTime string `json:"execution_time"`
}
type Response_Single struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Results int `json:"results"`
	Data Movie `json:"data"`
	Errors []string `json:"errors"`
	ExecutionTime string `json:"execution_time"`
}


type Page_AllMovies struct {
    PageTitle string
    Movies []ShortMovie
}

type Page_MovieDetails struct {
    PageTitle string
    Movie Movie
}

func GetMovies(queryParams QueryParameters) (response Response_All) {
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
	// parameters.Add("item_per_workers", string(queryParams.ItemPerWorkers))

	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	log.Println(Url.String())

	if err != nil {
		defer resp.Body.Close()
		log.Fatalf(err.Error())
		var response Response_All
		return response
	} 

	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status, resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil { 
		panic(err)
	}
	return response	 
}

func GetMoviesById(id string) (response Response_Single) {
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
	fmt.Println("\n\tResponse status:", resp.Status, resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil { 
		panic(err)
	}
	return response	 
}

func GetQueryParams(r *http.Request) (queryParams QueryParameters) {
	keys := r.URL.Query()

	if val, ok := keys["type"]; ok {
		queryParams.Type = val[0]
	}
	
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

	response := GetMovies(QueryParameters{Items: queryParams.Items, ItemPerWorkers: 1, Type: queryParams.Type })
	

	tmpl := template.Must(template.ParseFiles("html/index.html"))
	data := Page_AllMovies{
		PageTitle: "Cine+",
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
	data := Page_MovieDetails{
		PageTitle: "Cine+",
		Movie: response.Data,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	

	if err := tmpl.Execute(w, response.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}


func main() {
  http.HandleFunc("/", RenderMovies)
  http.HandleFunc("/getMovieById", RenderMovieById)
  log.Println("Server running succesfully on port 3000!")
  log.Fatal(http.ListenAndServe(":3000", nil))
}


