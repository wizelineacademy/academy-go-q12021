package rest

type rickAndMortyApi struct {
}

type RickAndMortyApiRepository interface {
}

func NewRickAndMortyApiRepository() RickAndMortyApiRepository {
	return &rickAndMortyApi{}
}
