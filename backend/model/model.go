package model

// Movie struct (Model): meant to be used for a single movie request.
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

// MovieSummary struct (Model): meant to be used for a list of movies.
type MovieSummary struct {
	ImdbTitleId string `json:"imdb_title_id"`
	OriginalTitle string `json:"original_title"`
	Year string `json:"year"`
	Poster string `json:"poster"`
}
