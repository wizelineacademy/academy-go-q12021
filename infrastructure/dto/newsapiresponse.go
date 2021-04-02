package dto

type NewsApiResponse struct {
	Status string    `json:"status"`
	News   []NewItem `json:"news"`
}

type NewItem struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Author      string   `json:"author"`
	Image       string   `json:"image"`
	Language    string   `json:"language"`
	Category    []string `json:"category"`
	Published   string   `json:"published"`
}
