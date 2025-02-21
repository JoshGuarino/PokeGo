package pokego

import (
	"github.com/JoshGuarino/PokeGo/pkg/resources/berries"
	"github.com/JoshGuarino/PokeGo/pkg/resources/contests"
)

// PokeGo API wrapper client
type PokeGo struct {
	Berries  berries.Berries
	Contests contests.Contests
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		Berries:  berries.NewBerriesGroup(),
		Contests: contests.NewContestsGroup(),
	}
}
