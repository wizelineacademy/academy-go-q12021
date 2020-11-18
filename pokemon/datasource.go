package pokemon

// DataSource is an interface for collecting pokemon data
type DataSource interface {
	GetPokemonByName(name string) (Pokemon, error)
	Close()
}
