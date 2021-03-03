package model

//Book contains information about each book
type Book struct {
	ID       int    `json:"id"`
	Isbn     string `json:"isbn"`
	Authors  string `json:"authors"`
	Year     int    `json:"yer"`
	ImageURL string `json:"imageURL"`
}

// FindBookByID returns the index that corresponds to an ID (if exists)
func FindBookByID(books []Book, ID int) int {

	for index, item := range books {
		if item.ID == ID {
			return index
		}
	}

	return -1
}
