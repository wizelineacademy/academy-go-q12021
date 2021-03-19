package model

type Pokemon struct {
	Id             int    `csv:"ID"`
	Name           string `csv:"English"`
	Height         int    `csv:"Height"`
	Weight         int    `csv:"Weight"`
	BaseExperience int    `csv:"Base experience"`
	PrimaryType    string `csv:"Primary"`
	SecondaryType  string `csv:"Secondary"`
}
