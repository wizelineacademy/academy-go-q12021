package model

type Centre struct {
	Id			int     `json:"id"`
	Name		string  `json:"name"`
	Address		string  `json:"address"`
	Email		string  `json:"email"`
	Phone		string  `json:"phone"`
	Line		string  `json:"line"`
	Capacity	int     `json:"capacity"`
	Openness	int     `json:"openness"`
}

func (Centre) TableName() string { return "centres" }