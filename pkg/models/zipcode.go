package models

type ZipCodeCSV struct {
	Id           string `bson:"_id"`
	CodigoPostal string `bson:"codigoPostal"`
	Estado       string `bson:"estado"`
	EstadoISO    string `bson:"estadoISO"`
	Municipio    string `bson:"municipio"`
	Ciudad       string `bson:"ciudad"`
	Barrio       string `bson:"barrio"`
}
