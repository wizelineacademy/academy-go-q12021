package model

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Format string  `json:"format"`
	Price  float64 `json:"price"`
}
