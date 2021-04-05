package model

// QueryParameters (Model): meant to be used for a movie request.
type QueryParameters struct {
	ItemPerWorkers int    `json:"item_per_workers"`
	Items          int    `json:"items"`
	Type           string `json:"type"`
}

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

// MovieSummary struct (Model): meant to be used for a list of movies with only the neccesary fields.
type MovieSummary struct {
	ImdbTitleId   string `json:"imdb_title_id"`
	OriginalTitle string `json:"original_title"`
	Year          string `json:"year"`
	Poster        string `json:"poster"`
}

// Response_All useful for parsing the response from server with a list of MovieSummary as Data (Third deliverable)
type Response_All struct {
	Title         string         `json:"title"`
	Message       string         `json:"message"`
	Results       int            `json:"results"`
	Data          []MovieSummary `json:"data"`
	Errors        []string       `json:"errors"`
	ExecutionTime string         `json:"execution_time"`
}

// Response_Single useful for parsing the response from server with a single Movie as Data (Third deliverable)
type Response_Single struct {
	Title         string   `json:"title"`
	Message       string   `json:"message"`
	Results       int      `json:"results"`
	Data          Movie    `json:"data"`
	Errors        []string `json:"errors"`
	ExecutionTime string   `json:"execution_time"`
}

// Page_MovieDetails struct will render a page with a page title and a list (slice/array) of MovieSummary (Third deliverable)
type Page_AllMovies struct {
	PageTitle string
	Movies    []MovieSummary
}

// Page_MovieDetails struct will render a page with a single movie item (Third deliverable)
type Page_MovieDetails struct {
	PageTitle string
	Movie     Movie
}

// PageData struct will render a page with a single interface struct item (Second deliverable)
type PageData struct {
	PageTitle     string
	TechStackItem TechStackItem
}

// model.TechStackItem struct meant for the tech stack items (Second deliverable)
type TechStackItem struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Years string `json:"years"`
}
