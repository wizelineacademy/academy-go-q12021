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
	"strconv"
	"sync"
	"time"
)

/* Generic functions and structure */

const MaxUint64 = ^uint64(0) 
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

type QueryParameters struct {
	ItemPerWorkers string `json:"item_per_workers"`
    Items uint64 `json:"items"`
    Type string `json:"type"`
}
var movies []Movie
var requestErrors []string

func ConvertStructToJSON(obj interface{}) string {
    e, err := json.Marshal(obj)
    if err != nil {
		requestErrors = append(requestErrors, err.Error())
        return err.Error()
    }
    return string(e)
}

func Even(number int) bool {
    return number%2 == 0
}

func Odd(number int) bool {
    return !Even(number)
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

func worker(jobs <-chan string, results chan<- Movie, wg *sync.WaitGroup, queryParams QueryParameters) {
	itemsToDisplay := queryParams.Items
	numberType := queryParams.Type
	log.Println("\nItems per response: ", itemsToDisplay, "\nItems per worker: ", 0,"\nType: ", numberType,)

	defer wg.Done()

	var moviesAddedCounter uint64

	for line := range jobs {
		lineItems := SplitAtCommas(line)
		newMovie := Movie{
			ImdbTitleId: lineItems[0],
			Title: lineItems[1],
			OriginalTitle: lineItems[2],
			Year: lineItems[3],
			DatePublished: lineItems[4],
			Genre: lineItems[5],
			Duration: lineItems[6],
			Country: lineItems[7],
			Language: lineItems[8],
			Director: lineItems[9],
			Writer: lineItems[10],
			ProductionCompany: lineItems[11],
			Actors: lineItems[12],
			Description: lineItems[13],
			AvgVote: lineItems[14],
			Votes: lineItems[15],
			Budget: lineItems[16],
			UsaGrossIncome: lineItems[17],
			WorlwideGrossIncome: lineItems[18],
			Metascore: lineItems[19],
			ReviewsFromUsers: lineItems[20],
			ReviewsFromCritics: lineItems[21],
		}
		// get id from Movie struct and parse the string to a number
		inputFmt := newMovie.ImdbTitleId[2:len(newMovie.ImdbTitleId)] // get substring of id
		id, err := strconv.Atoi(inputFmt) // parse substring to int

		log.Println(moviesAddedCounter < itemsToDisplay, moviesAddedCounter, itemsToDisplay)
		if err != nil {
			requestErrors = append(requestErrors, err.Error())
			return
		}	else if moviesAddedCounter < itemsToDisplay {
			if numberType ==  "odd" && Odd(id) {
				log.Println("The Id is Odd: ", id)
				moviesAddedCounter++
				results <- newMovie
			}
			if numberType ==  "even" && Even(id) {
				log.Println("The Id is Odd: ", id)
				moviesAddedCounter++
				results <- newMovie
			}
			if numberType !=  "even" && numberType !=  "odd" {
				log.Println("Display both even and odd numbers")
				moviesAddedCounter++
				results <- newMovie
			}
		}
	}
}



func GetMoviesFromFileConcurrently(queryParams QueryParameters) {
	itemsPerWorkers, err := strconv.Atoi(queryParams.ItemPerWorkers) // parse substring to int
	if err != nil {
		requestErrors = append(requestErrors, err.Error())
		log.Println(err.Error())
	}

    file, err := os.Open("IMDb_movies_short.csv")
    if err != nil {
		requestErrors = append(requestErrors, err.Error())
      	log.Fatal(err)
    }
    defer file.Close()
  
    jobs := make(chan string)
    results := make(chan Movie)
  
    wg := new(sync.WaitGroup)
  
    // start workers
    var workers = itemsPerWorkers

    for w := 1; w <= workers; w++ {
      wg.Add(1)
      go worker(jobs, results, wg, queryParams)
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
	start := time.Now() 
	w.Header().Set("Content-Type", "application/json")

	// GET QUERY PARAMS AND VALIDATE
	keys := r.URL.Query()
	var queryParams QueryParameters

	if val, ok := keys["type"]; ok {
		log.Println("Type query provided")
		queryParams.Type = val[0]
	} else {
		requestErrors = append(requestErrors, "`type` was not provided as query param. Should be rather odd or even.")
		log.Println("Type not provided as query param.")
	}

	if val, ok := keys["item_per_workers"]; ok {
		log.Println("item_per_workers query provided")
		queryParams.ItemPerWorkers = val[0]
	} else {
		requestErrors = append(requestErrors, "`items_per_workers` was not provided as query param.")
		log.Println("item_per_workers not provided as query param")
	}

	if val, ok := keys["items"]; ok {
		uIntItems, err := strconv.ParseUint(val[0],10,32) // parse string to uint
		if err != nil {
			requestErrors = append(requestErrors, err.Error() + ". Number should be positive integer. The items param will be considered as 0. ")
			queryParams.Items = 0
		} else {
			queryParams.Items = uIntItems
			log.Println("items query provided: value ", uIntItems)	
		}
	} else {
		requestErrors = append(requestErrors, "`items` was not provided as query param: MaxValue")
		queryParams.Items = MaxUint64
	}

	GetMoviesFromFileConcurrently(queryParams)

	totalTime :=  fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")

	jsonObject := Response{ 
		Title: "Response", 
		Results: len(movies),
		Message: "Data",
		Data: movies,
		Errors: requestErrors,
		ExecutionTime: totalTime,
	}
	jsonResult := ConvertStructToJSON(jsonObject)

	fmt.Fprintf(w, "%s", jsonResult)
	log.Println(" \t Number of Parsed Movies: ", len(movies), " \t TIME: " ,totalTime)	
	requestErrors = nil
}


func GetMovieById(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys) <= 0 {
		requestErrors = append(requestErrors, "Id query param is required but missing")
		log.Println("Id query param is required but missing")
	}
	// id := keys[0]

	// GetMoviesFromFileConcurrently(id)
	// var selectedMovie Movie
	// for _, movie := range movies {
	// 	if movie.ImdbTitleId == id {
	// 		// Found!
	// 		selectedMovie = movie
	// 		break
	// 	}
	// }	
	// imgUrl := GetMoviePosterFromOmdbApi(selectedMovie.OriginalTitle, selectedMovie.Year)
	// selectedMovie.Poster = imgUrl
	// return selectedMovie
}

func GetMoviePosterFromOmdbApi(title string, year string) (imageUrl string) {
	// Consume the api of omdbapi
	Url, err := url.Parse("http://www.omdbapi.com/")
    if err != nil {
		requestErrors = append(requestErrors,err.Error())
        log.Fatal(err.Error())
    }
    parameters := url.Values{}
    parameters.Add("apikey", "43502af4")
    parameters.Add("t", title)
    parameters.Add("y", year)
    Url.RawQuery = parameters.Encode()
    fmt.Printf("Encoded URL is %q\n", Url.String())

	resp, err := http.Get(Url.String())
	if err != nil {
		requestErrors = append(requestErrors, err.Error())
        log.Fatal(err.Error())
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
		requestErrors = append(requestErrors,err.Error())
        log.Println(err.Error())
	}
	return imageUrl
}

func main() {
	requestErrors = nil
	http.HandleFunc("/getMovies", GetMovies)
	http.HandleFunc("/getMovieById", GetMovieById)
	log.Println("Server running succesfully on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


