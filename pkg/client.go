package pokego

import (
	"github.com/JoshGuarino/PokeGo/pkg/resources/berries"
	"github.com/JoshGuarino/PokeGo/pkg/resources/contests"
	"github.com/JoshGuarino/PokeGo/pkg/resources/encounters"
)

// PokeGo API wrapper client
type PokeGo struct {
	Berries    berries.Berries
	Contests   contests.Contests
	Encounters encounters.Encounters
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		Berries:    berries.NewBerriesGroup(),
		Contests:   contests.NewContestsGroup(),
		Encounters: encounters.NewEncountersGroup(),
	}
}
