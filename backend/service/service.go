package service

import (
	"encoding/csv"
	"fmt"
	"os"

	"main/model"
)

// Service
type Service struct {
	csvr *os.File
	csvw *csv.Writer
}

// New creates a new Service layer
func New(
	csvr *os.File,
	csvw *csv.Writer) (*Service, error) {

	return &Service{csvr, csvw}, nil
}


// GetMovies -
func (s *Service) GetMovies() ([]*model.Movie, error) {
	out := []*model.Movie{}

	csvr := csv.NewReader(s.csvr)

	data, err := csvr.ReadAll()
	if err != nil {
		fmt.Print("Error reading")
		return nil, err
	}

	for _, row := range data {
		movie := model.Movie{
			ImdbTitleId: row[0],
			Title: row[1],
			OriginalTitle: row[2],
			Year: row[3],
			DatePublished: row[4],
			Genre: row[5],
			Duration: row[6],
			Country: row[7],
			Language: row[8],
			Director: row[9],
			Writer: row[10],
			ProductionCompany: row[11],
			Actors: row[12],
			Description: row[13],
			AvgVote: row[14],
			Votes: row[15],
			Budget: row[16],
			UsaGrossIncome: row[17],
			WorlwideGrossIncome: row[18],
			Metascore: row[19],
			ReviewsFromUsers: row[20],
			ReviewsFromCritics: row[21],
			Poster: row[22],
		}
		out = append(out, &movie)
	}
	s.csvr.Seek(0, 0)

	return out, nil
}

// GetMovieById -
func (s *Service) GetMovieById(movieID string) (*model.Movie, error) {
	out := &model.Movie{}

	csvr := csv.NewReader(s.csvr)

	data, err := csvr.ReadAll()
	if err != nil {
		fmt.Print("Error reading")
		return nil, err
	}

	for _, row := range data {
		if row[0] == movieID {
			out.ImdbTitleId = row[0]
			out.Title = row[1]
			out.Year = row[2]
		}
	}
	s.csvr.Seek(0, 0)

	return out, nil
}
