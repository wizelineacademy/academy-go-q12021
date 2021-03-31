package main

// dataset gathered from: https://www.kaggle.com/stefanoleone992/imdb-extensive-dataset

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"text/template"
	"time"
)

/* Generic functions and structure */


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

type PageData struct {
    PageTitle string
    Movies     []Movie
}
type Response struct {
	Title string `json:"title"`
	Message string `json:"message"`
}

var movies []Movie

func ConvertStructToJSON(obj interface{}) string {
    e, err := json.Marshal(obj)
    if err != nil {
        return err.Error()
    }
    return string(e)
}

// following function from: https://play.golang.org/p/f5jceIm4nbE
func SplitAtCommas(s string) []string {
    res := []string{}
    var beg int
    var inString bool

    for i := 0; i < len(s); i++ {
        if s[i] == ',' && !inString {
            res = append(res, s[beg:i])
            beg = i+1
        } else if s[i] == '"' {
            if !inString {
                inString = true
            } else if i > 0 && s[i-1] != '\\' {
                inString = false
            }
        }
    }
    return append(res, s[beg:])
}

func worker(jobs <-chan string, results chan<- Movie, wg *sync.WaitGroup) {
  // Decreasing internal counter for wait-group as soon as goroutine finishes
  defer wg.Done()
  // eventually I want to have a []string channel to work on a chunk of lines not just one line of text
  for line := range jobs {
    items := SplitAtCommas(line)
    newMovie := Movie{
        ImdbTitleId: items[0],
        Title: items[1],
        OriginalTitle: items[2],
        Year: items[3],
		DatePublished: items[4],
		Genre: items[5],
		Duration: items[6],
		Country: items[7],
		Language: items[8],
		Director: items[9],
		Writer: items[10],
		ProductionCompany: items[11],
		Actors: items[12],
		Description: items[13],
		AvgVote: items[14],
		Votes: items[15],
		Budget: items[16],
		UsaGrossIncome: items[17],
		WorlwideGrossIncome: items[18],
		Metascore: items[19],
		ReviewsFromUsers: items[20],
		ReviewsFromCritics: items[21],
    }
    results <- newMovie
  }
}

func GetMoviesConcurrently() {
    file, err := os.Open("IMDb_movies.csv")
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()
  
    jobs := make(chan string)
    results := make(chan Movie)
  
    wg := new(sync.WaitGroup)
  
    // start workers
    const workers = 1
    for w := 1; w <= workers; w++ {
      wg.Add(1)
      go worker(jobs, results, wg)
    }
  
    // scan the file into the string channel
    go func() {
      scanner := bufio.NewScanner(file)
      for scanner.Scan() {
        // Later I want to create a buffer of lines, not just line-by-line here ...
        jobs <- scanner.Text()
      }
      close(jobs)
    }()
  
    // Collect all the results,  make sure we close the result channel when everything was processed
    go func() {
      wg.Wait()
      close(results)
    }()

	movies = nil
    // Convert channel to slice of Movie and send
    for movie := range results {
		movies = append(movies,movie)
    }
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// Get query params
	// keys, ok := r.URL.Query()["type"]
    // if !ok || len(keys[0]) < 1 {
	// 	errorMessage := "Url Param 'type' is missing"
	// 	log.Println(errorMessage)
	// 	fmt.Fprintf(w, "%s", errorMessage)
    //     return
    // }

	start := time.Now() 
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	GetMoviesConcurrently()
	log.Println(" \t Number of Parsed Movies: ", len(movies))
	func () {
		data := PageData{
			PageTitle: "IMDb Movies",
			Movies: movies,
		}
		tmpl.Execute(w, data)
	}()
	log.Println(" \t TIME: " ,time.Since(start).Microseconds(), " Microseconds.")	
}

func RenderMovie(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
    if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
        return
    }
	// Casting the string number to an integer
	id := keys[0]
	item := GetMovieById(id)

	tmpl := template.Must(template.ParseFiles("html/item.html"))

	if err := tmpl.Execute(w, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}

func GetMovieById(id string) Movie {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	// var url string = "http://localhost:8080/getLanguageById?id=" + id
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// // Print the HTTP response status.
	// fmt.Println("\n\tResponse status:", resp.Status)

	// // Print the first 5 lines of the response body.
	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
    // 	json.Unmarshal([]byte(scanner.Text()), &movie) // items slice
	// }
	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
	var selectedMovie Movie
	for _, movie := range movies {
		if movie.ImdbTitleId == id {
			// Found!
			selectedMovie = movie
			break
		}
	}	
	imgUrl := GetMoviePoster(selectedMovie.OriginalTitle, selectedMovie.Year)


	selectedMovie.Poster = imgUrl
	return selectedMovie
}

func GetMoviePoster(title string, year string) (imageUrl string) {
	// Consume the api of omdbapi
	// urlSlice := []string{"http://www.omdbapi.com/?apikey=", "43502af4","&t=", title, "&y=", year }
	// var url string = strings.Join(urlSlice, "")
	// log.Println("url: ", url)

	Url, err := url.Parse("http://www.omdbapi.com/")
    if err != nil {
        panic("boom")
    }
    parameters := url.Values{}
    parameters.Add("apikey", "43502af4")
    parameters.Add("t", title)
    parameters.Add("y", year)
    Url.RawQuery = parameters.Encode()
    fmt.Printf("Encoded URL is %q\n", Url.String())

	resp, err := http.Get(Url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		var movie Movie
    	json.Unmarshal([]byte(scanner.Text()), &movie) 
		imageUrl = movie.Poster
		log.Println("Movie Found on OMDBAPI: ", movie)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return imageUrl
}

func main() {
  http.HandleFunc("/", GetMovies)
  http.HandleFunc("/getMovieById", RenderMovie)

  log.Println("Server running succesfully on port 8080!")
  log.Fatal(http.ListenAndServe(":8080", nil))
}


