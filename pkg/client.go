package pokego

import (
	"github.com/JoshGuarino/PokeGo/pkg/resources/berries"
)

// PokeGo client
type PokeGo struct {
	Berries berries.Berries
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		Berries: berries.NewBerriesGroup(),
	}
}
