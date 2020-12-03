package model

// Student - struct to store an student
type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      int    `json:"zip"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}
