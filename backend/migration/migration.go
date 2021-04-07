package migrations

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
)

// Errors that will be sent to the api response
var requestErrors []string

// Method used to gather images from the Omdb api
func GetMoviePosterFromOmdbApi(title string, year string) (imageUrl string) {
	// Consume the api of omdbapi
	Url, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		requestErrors = append(requestErrors, err.Error())
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
		var movie model.Movie
		json.Unmarshal([]byte(scanner.Text()), &movie)
		imageUrl = movie.Poster
		log.Println("model.Movie Found on OMDBAPI: ", movie)
	}
	if err := scanner.Err(); err != nil {
		requestErrors = append(requestErrors, err.Error())
		log.Println(err.Error())
	}
	return imageUrl
}

// The following method was used to populate the .csv with images since it came without those an no one likes a ui without images.
func WriteDataToCSVFile(fileName string, movies []model.Movie) {
	log.Println("Data: ", movies)

	csvfile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error creating file csv", err)
	}
	var writter *csv.Writer = csv.NewWriter(csvfile)

	for _, movie := range movies {
		strSlice := []string{
			movie.ImdbTitleId,
			movie.Title,
			movie.OriginalTitle,
			movie.Year,
			movie.DatePublished,
			movie.Genre,
			movie.Duration,
			movie.Country,
			movie.Language,
			movie.Director,
			movie.Writer,
			movie.ProductionCompany,
			movie.Actors,
			movie.Description,
			movie.AvgVote,
			movie.Votes,
			movie.Budget,
			movie.UsaGrossIncome,
			movie.WorlwideGrossIncome,
			movie.Metascore,
			movie.ReviewsFromUsers,
			movie.ReviewsFromCritics,
			movie.Poster,
		}
		fmt.Println(strSlice)
		writter.Write(strSlice)
	}
	// Write any buffered movies data to the underlying writer (standard output).
	writter.Flush()

	if err := writter.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}
