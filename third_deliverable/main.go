package main

// dataset gathered from: https://www.kaggle.com/stefanoleone992/imdb-extensive-dataset

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

/* Generic functions and structure */
type Response struct {
	Title string `json:"title"`
	Message string `json:"message"`
}

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

var movie Movie = Movie{
	ImdbTitleId: "1",
    Title: "The Hunger Games: Catching Fire",
	OriginalTitle: "Catching Fire",
	Year: "1996",
}

// func displayError(w http.ResponseWriter, message string) {
//     log.Println(message)
// 	fmt.Fprintf(w, "%s", ConvertStructToJSON(Response{Title: "Error", Message: message}))

// }

// func getMovieById(w http.ResponseWriter, r *http.Request) {
// 	csvLines := GetDataFromCSVFile("IMDb_movies.csv")
// 	listOfMovies :=  ParseCSVDataToMovieList(csvLines)   
// 	// Obtain the query param id number from URL
// 	keys, ok := r.URL.Query()["id"]
//     if !ok || len(keys[0]) < 1 {
// 		displayError(w, "Url Param 'id' is missing!")
//         return
//     }
// 	// Casting the string number to an integer
//     id, err := strconv.Atoi(keys[0])
// 	if err != nil {
// 		displayError(w, "The Id provided is wrong, please check it!")
// 		return
// 	}
// 	// Validations: number is positive and that exists as index in the slice 
// 	if (id >= len(listOfMovies) || id < 0) {
// 		displayError(w, "The Id doesn't seem to exist!")
// 		return
// 	}
// 	// Get the object from slice using the id as index
// 	obj := listOfMovies[id]
//     fmt.Fprintf(w, "%s", ConvertStructToJSON(obj))
// }


func ConvertStructToJSON(obj interface{}) string {
    e, err := json.Marshal(obj)
    if err != nil {
        return err.Error()
    }
    return string(e)
}

func GetDataFromCSVFile(filePath string) ([][] string)  {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("\n- An error ocurred while reading the file. \n", err)
	} else {
		fmt.Println("\nSuccessfully Opened CSV file") 
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
		log.Println("\n- An error ocurred while reading the file. \n", err)
		return nil
    }
	return csvLines
}

func ParseCSVDataToMovieList(csvLines [][]string) (listOfMovies []Movie ) { 
	// Convert csv lines to a movie structure and append them to the array of items
    for _, line := range csvLines {
		newMovie := Movie{
            ImdbTitleId: line[0],
            Title: line[1],
			OriginalTitle: line[2],
			Year: line[3],
        }
        listOfMovies = append(listOfMovies, newMovie)
        log.Println(newMovie.ImdbTitleId + " " + newMovie.Title)
    }
	return 
}


func worker(ID int, jobs <-chan Movie, results chan<- Movie) {
	for job := range jobs {
		fmt.Println("Worker ", ID, " is working on job ", job)
		duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Worker ", ID, " completed work on job ", job, " within ", duration)
		results <- movie
	}
}

func WorkerPool() {
	jobs := make(chan Movie)
	results := make(chan Movie)
	// 3 Workers
	for x := 1; x <= 3; x++ {
		go worker(x, jobs, results)
	}
	// Give them jobs
	for j := 1; j <= 6; j++ {
		jobs <- movie
	}
	close(jobs)
	// Wait for the results
	for r := 1; r <= 6; r++ {
		fmt.Println("Result received from worker: ", <-results)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	go WorkerPool()
	// csvLines := GetDataFromCSVFile("IMDb_movies.csv")
	// listOfMovies :=  ParseCSVDataToMovieList(csvLines)   
	// fmt.Fprintf(w, "%s", ConvertStructToJSON(listOfMovies))
}

func main() {
    http.HandleFunc("/getMovies", getMovies)
	// http.HandleFunc("/getMovieById", getMovieById)
	log.Println("Server running succesfully on port 8080!")
    log.Fatal(http.ListenAndServe(":8080", nil))
}