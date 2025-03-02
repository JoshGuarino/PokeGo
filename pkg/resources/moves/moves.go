package moves

// Moves group interface
type IMoves interface {
}

// Moves group struct
type Moves struct{}

// Return an instance of Pokmon resource group struct
func NewMovesGroup() Moves {
	return Moves{}
}
