package model

// Movie struct (Model): meant to be used for a single movie request.
type Movie struct {
	ImdbTitleId         string `json:"imdb_title_id"`
	Title               string `json:"title"`
	OriginalTitle       string `json:"original_title"`
	Year                string `json:"year"`
	DatePublished       string `json:"date_published"`
	Genre               string `json:"genre"`
	Duration            string `json:"duration"`
	Country             string `json:"country"`
	Language            string `json:"language"`
	Director            string `json:"director"`
	Writer              string `json:"writer"`
	ProductionCompany   string `json:"production_company"`
	Actors              string `json:"actors"`
	Description         string `json:"description"`
	AvgVote             string `json:"avg_vote"`
	Votes               string `json:"votes"`
	Budget              string `json:"budget"`
	UsaGrossIncome      string `json:"usa_gross_income"`
	WorlwideGrossIncome string `json:"worlwide_gross_income"`
	Metascore           string `json:"metascore"`
	ReviewsFromUsers    string `json:"reviews_from_users"`
	ReviewsFromCritics  string `json:"reviews_from_critics"`
	Poster              string `json:"poster"`
}

// MovieSummary struct (Model): meant to be used for a list of movies.
type MovieSummary struct {
	ImdbTitleId   string `json:"imdb_title_id"`
	OriginalTitle string `json:"original_title"`
	Year          string `json:"year"`
	Poster        string `json:"poster"`
}

// model.QueryParameters struct (Model)
type QueryParameters struct {
	ItemPerWorkers int    `json:"item_per_workers"`
	Items          int    `json:"items"`
	Type           string `json:"type"`
}

// // Response_Single struct (Model)
type Response struct {
	Title         string        `json:"title"`
	Message       string        `json:"message"`
	Results       int           `json:"results"`
	ExecutionTime string        `json:"execution_time"`
	Data          []interface{} `json:"data"`
	Errors        []string      `json:"errors"`
}

type Response_All struct {
	Title         string          `json:"title"`
	Message       string          `json:"message"`
	Results       int             `json:"results"`
	Data          []*MovieSummary `json:"data"`
	Errors        []string        `json:"errors"`
	ExecutionTime string          `json:"execution_time"`
}
type Response_Single struct {
	Title         string   `json:"title"`
	Message       string   `json:"message"`
	Results       int      `json:"results"`
	Data          Movie    `json:"data"`
	Errors        []string `json:"errors"`
	ExecutionTime string   `json:"execution_time"`
}

type Item struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Years string `json:"years"`
}
