package pokemon

// Pokemon group interface
type IPokemon interface {
}

// Pokemon group struct
type Pokemon struct{}

// Return an instance of Pokmon resource group struct
func NewPokemonGroup() Pokemon {
	return Pokemon{}
}
