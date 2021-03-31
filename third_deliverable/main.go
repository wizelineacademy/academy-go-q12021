package main

// dataset gathered from: https://www.kaggle.com/stefanoleone992/imdb-extensive-dataset

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
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
	// date_published string
	// genre string
	// duration string
	// country string
	// language string
	// director string
	// writer string
	// production_company string
	// actors string
	// description string
	// avg_vote string
	// votes string
	// budget string
	// usa_gross_income string
	// worlwide_gross_income string
	// metascore string
	// reviews_from_users string
	// reviews_from_critics string
}
type PageData struct {
    PageTitle string
    Movies     []Movie
}
type Response struct {
	Title string `json:"title"`
	Message string `json:"message"`
}


var movie Movie = Movie{
	ImdbTitleId: "1",
    Title: "The Hunger Games: Catching Fire",
	OriginalTitle: "Catching Fire",
	Year: "1996",
}

func ConvertStructToJSON(obj interface{}) string {
    e, err := json.Marshal(obj)
    if err != nil {
        return err.Error()
    }
    return string(e)
}

func worker(jobs <-chan string, results chan<- Movie, wg *sync.WaitGroup) {
  // Decreasing internal counter for wait-group as soon as goroutine finishes
  defer wg.Done()
  // eventually I want to have a []string channel to work on a chunk of lines not just one line of text
  for line := range jobs {
    items := strings.Split(line, ",")
    newMovie := Movie{
        ImdbTitleId: items[0],
        Title: items[1],
        OriginalTitle: items[2],
        Year: items[3],
    }
    results <- newMovie
  }
}
var movies []Movie

func GetMoviesConcurrently() {
    file, err := os.Open("IMDb_movies_short.csv")
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
  
    // Convert channel to slice of Movie and send
    for movie := range results {
		movies = append(movies,movie)
    }
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))

	start := time.Now() 
	GetMoviesConcurrently()

	log.Println(" \t Movies Parsed: ", len(movies), " Movies:", movies)
	func () {
		data := PageData{
			PageTitle: "IMDb Movies",
			Movies: movies,
		}
		tmpl.Execute(w, data)
	}()

	log.Println(" \t TIME: " ,time.Since(start).Microseconds(), " Microseconds.")	
}

func main() {
  http.HandleFunc("/", getMovies)
  //http.HandleFunc("/getMovie", getMovies)
  log.Println("Server running succesfully on port 8080!")
  log.Fatal(http.ListenAndServe(":8080", nil))
}


