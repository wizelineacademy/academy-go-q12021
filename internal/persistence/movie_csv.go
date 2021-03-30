package persistence

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/marshal"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

const (
	totalItemsWorkersFilterKey = "items_per_worker"
	searchGenericFilterKey     = "type"
)

// MovieCSV handles all persistence Movie's operations locally using an specific `.csv` file
//	Implements Movie repository
type MovieCSV struct {
	mu  sync.RWMutex
	cfg infrastructure.Configuration
}

// NewMovieCSV allocates a Movie repository CSV concrete implementation
func NewMovieCSV(config infrastructure.Configuration) *MovieCSV {
	return &MovieCSV{
		mu:  sync.RWMutex{},
		cfg: config,
	}
}

// Get retrieves a Movie by its ID
func (m *MovieCSV) Get(ctx context.Context, id valueobject.MovieID) (*aggregate.Movie, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	file, err := os.Open(m.cfg.MoviesDataset)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	return m.searchMovieOnFile(csv.NewReader(file), id)
}

func (m *MovieCSV) searchMovieOnFile(r *csv.Reader, id valueobject.MovieID) (*aggregate.Movie, error) {
	isHeader := true
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		} else if isHeader {
			isHeader = false
			continue
		} else if err != nil {
			return nil, err
		}

		movie := aggregate.NewEmptyMovie()
		if err = marshal.UnmarshalMovieCSV(movie, records...); err != nil {
			return nil, err
		} else if movie.ID == id {
			return movie, nil
		}
	}

	return nil, nil
}

// Search retrieves a set of Movies by the given criteria filters, returns the set, a next page token or an error
func (m *MovieCSV) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	file, err := os.Open(m.cfg.MoviesDataset)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		err = file.Close()
	}()

	reader := csv.NewReader(file)
	if workerItems := criteria.Query.Filters[totalItemsWorkersFilterKey].Value.(string); workerItems != "" {
		return m.searchMoviesOnFileParallel(reader, criteria)
	}

	return m.searchMoviesOnFile(reader, criteria)
}

func (m *MovieCSV) searchMoviesOnFileParallel(r *csv.Reader, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	records, err := r.ReadAll()
	if err != nil {
		return nil, "", err
	} else if len(records) >= 2 {
		records = records[1:] // remove header
	}

	totalWorkers := len(records)
	if len(records) > 100 {
		totalWorkers = 100 // avoid more than 100 workers
	}

	totalItemsPerWorker, err := strconv.Atoi(criteria.Query.Filters[totalItemsWorkersFilterKey].Value.(string))
	if err != nil {
		return nil, "", err
	}

	movies := make([]*aggregate.Movie, 0)
	movieChan := make(chan *aggregate.Movie, len(records))
	jobs := make(chan []string, len(records))
	workerWg := new(sync.WaitGroup)
	workerWg.Add(totalWorkers)
	for i := 0; i < totalWorkers; i++ {
		go m.searchMoviesWorker(criteria.Query.Filters[searchGenericFilterKey].Value.(string),
			totalItemsPerWorker, jobs, workerWg, movieChan)
	}

	go m.enqueueSearchMoviesJobs(records, jobs)
	workerWg.Wait()
	close(movieChan)

	for movie := range movieChan {
		movies = append(movies, movie)
		if totalMovies := len(movies); totalMovies == criteria.Limit {
			break
		}
	}
	return movies, "", nil
}

func (m *MovieCSV) searchMoviesWorker(filter string, totalItems int, jobs <-chan []string, wg *sync.WaitGroup, movieChan chan<- *aggregate.Movie) {
	defer wg.Done()
	validItemsCount := 0
	for record := range jobs {
		if validItemsCount == totalItems {
			return
		}

		movie := aggregate.NewEmptyMovie()
		if err := marshal.UnmarshalMovieCSV(movie, record...); err == nil {
			validItemsCount++
		}

		cleanedID, err := strconv.Atoi(strings.TrimPrefix(string(movie.IMDbID), "tt"))
		if err != nil {
			continue
		} else if filter == "odd" && cleanedID%2 == 0 {
			continue
		} else if filter == "even" && cleanedID%2 != 0 {
			continue
		}
		movieChan <- movie
	}
}

func (m *MovieCSV) enqueueSearchMoviesJobs(records [][]string, jobs chan<- []string) {
	for _, record := range records {
		jobs <- record
	}
	close(jobs)
}

func (m *MovieCSV) searchMoviesOnFile(r *csv.Reader, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	isHeader := true
	totalCount := 0
	allowFetch := false
	movies := make([]*aggregate.Movie, 0)
	nextPageToken := ""
csvTraverse:
	for {
		record, err := r.Read()
		switch {
		case err == io.EOF:
			break csvTraverse
		case err != nil:
			return nil, "", err
		case isHeader:
			isHeader = false
			continue csvTraverse
		case criteria.NextPage != "" && record[0] != criteria.NextPage && allowFetch == false:
			continue csvTraverse
		}

		movie := aggregate.NewEmptyMovie()
		if err = marshal.UnmarshalMovieCSV(movie, record...); err != nil {
			return nil, "", err
		} else if totalCount >= criteria.Limit { // fetch one more item to get next page
			nextPageToken = string(movie.ID)
			break
		}

		movies = append(movies, movie)
		totalCount++
		allowFetch = true
	}

	return movies, nextPageToken, nil
}

// Save stores the current state of the given Movie
func (m *MovieCSV) Save(ctx context.Context, movie aggregate.Movie) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	file, err := os.OpenFile(m.cfg.MoviesDataset, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	return m.writeToFile(file, movie)
}

func (m *MovieCSV) writeToFile(file io.Writer, movie aggregate.Movie) error {
	w := csv.NewWriter(file)
	defer w.Flush()
	fields := []string{
		string(movie.ID),
		string(movie.DisplayName),
		string(valueobject.MarshalDirectorsString(movie.Directors...)),
		strconv.Itoa(int(movie.ReleaseYear)),
		string(movie.IMDbID),
	}
	if err := w.Write(fields); err != nil {
		return err
	}
	return w.Error()
}
