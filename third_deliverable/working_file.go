package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Movie struct {
	ImdbTitleId string `json:"imdb_title_id"`
    Title string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Year string `json:"year"`
}

// func worker(lines chan string, finished chan Movie) {
//   // defer waitgroup.Done()	// Let the main process know we're done.
//   for line := range lines {
//     // Do the work with the line.
  //   items := strings.Split(line, ",")
  //   newMovie := Movie{
  //       ImdbTitleId: items[0],
  //       Title: items[1],
  //       OriginalTitle: items[2],
  //       Year: items[3],
  //   }
  //   finished <- newMovie
  // }
// }

func processFile() int {
  file, err := os.Open("IMDb_movies.csv")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  jobs := make(chan string)
  results := make(chan Movie)

  wg := new(sync.WaitGroup)

  // start up some workers that will block and wait?
  const workers = 20
  for w := 1; w <= workers; w++ {
    wg.Add(1)
    go worker(jobs, results, wg)
  }

  // Go over a file line by line and queue up a ton of work
  go func() {
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      // Later I want to create a buffer of lines, not just line-by-line here ...
      jobs <- scanner.Text()
    }
    close(jobs)
  }()

  // Now collect all the results...
  // But first, make sure we close the result channel when everything was processed
  go func() {
    wg.Wait()
    close(results)
  }()

  // Add up the results from the results channel.
  counts := 0
  
  for movie := range results {
    _ = movie
    fmt.Println("#", counts, " ", movie.Title)
    counts++
  }

  return counts
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

func main() {
  start := time.Now()
  numberOfMovies := processFile()
  fmt.Println(" \t Movies Parsed: ", numberOfMovies, " \t TIME: " ,time.Since(start).Microseconds(), " Microseconds.")

}


