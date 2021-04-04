package model

type QueryParameters struct {
	ItemPerWorkers int    `json:"item_per_workers"`
	Items          int    `json:"items"`
	Type           string `json:"type"`
}

/* Movie structure */
type ShortMovie struct {
	ImdbTitleId   string `json:"imdb_title_id"`
	OriginalTitle string `json:"original_title"`
	Year          string `json:"year"`
	Poster        string `json:"poster"`
}

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

type Response_All struct {
	Title         string       `json:"title"`
	Message       string       `json:"message"`
	Results       int          `json:"results"`
	Data          []ShortMovie `json:"data"`
	Errors        []string     `json:"errors"`
	ExecutionTime string       `json:"execution_time"`
}
type Response_Single struct {
	Title         string   `json:"title"`
	Message       string   `json:"message"`
	Results       int      `json:"results"`
	Data          Movie    `json:"data"`
	Errors        []string `json:"errors"`
	ExecutionTime string   `json:"execution_time"`
}

type Page_AllMovies struct {
	PageTitle string
	Movies    []ShortMovie
}

type Page_MovieDetails struct {
	PageTitle string
	Movie     Movie
}
