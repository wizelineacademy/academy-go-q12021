package model

type Pokemon struct {
	Id            int    `csv:"ID"`
	Name          string `csv:"English"`
	JapaneseName  string `csv:"Japanese"`
	PrimaryType   string `csv:"Primary"`
	SecondaryType string `csv:"Secondary"`
	EvolvesTo     string `csv:"Evolves into"`
	Information   string `csv:"Notes"`
}
