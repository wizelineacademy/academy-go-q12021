package model

//Book contains information about each book
type Book struct {
	ID       int    `json:"id"`
	Isbn     string `json:"isbn"`
	Authors  string `json:"authors"`
	Year     int    `json:"yer"`
	ImageURL string `json:"imageURL"`
}
