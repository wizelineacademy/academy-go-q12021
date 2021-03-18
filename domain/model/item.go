package model

type Item struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Item) TableName() string { return "`items.csv`" }