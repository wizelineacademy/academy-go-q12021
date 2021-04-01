package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

/* Movie structure */
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

func GetMovies() (response Response) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080/getMovies")
	if err != nil {
		log.Fatal(err.Error())
	}
	parameters := url.Values{}
	parameters.Add("type", "")
	parameters.Add("items", "10")
	parameters.Add("item_per_workers", "1")

	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	resp, err := http.Get(Url.String())
	
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

func GetMoviesById(id string) (movie Movie) {
	// Get the http reponse from api localhost:8080 backend
	Url, err := url.Parse("http://localhost:8080")
	if err != nil {
		// requestErrors = append(requestErrors,err.Error())
		log.Fatal(err.Error())
	}
	parameters := url.Values{}
	parameters.Add("id", "")
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
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {

    	json.Unmarshal([]byte(scanner.Text()), &movie) // items slice
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return movie
}

func RenderMovies(w http.ResponseWriter, r *http.Request) {
	// Casting the string number to an integer
	response := GetMovies()

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
	movie := GetMoviesById(id)
	tmpl := template.Must(template.ParseFiles("html/item.html"))

	if err := tmpl.Execute(w, movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}


func main() {
  http.HandleFunc("/", RenderMovies)
  http.HandleFunc("/getMovieById", RenderMovieById)
  log.Println("Server running succesfully on port 3000!")
  log.Fatal(http.ListenAndServe(":3000", nil))
}


