package pokego

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/JoshGuarino/PokeGo/pkg/resources/berries"
	"github.com/JoshGuarino/PokeGo/pkg/resources/contests"
	"github.com/JoshGuarino/PokeGo/pkg/resources/encounters"
	"github.com/JoshGuarino/PokeGo/pkg/resources/evolution"
	"github.com/JoshGuarino/PokeGo/pkg/resources/games"
	"github.com/JoshGuarino/PokeGo/pkg/resources/items"
	"github.com/JoshGuarino/PokeGo/pkg/resources/locations"
	"github.com/JoshGuarino/PokeGo/pkg/resources/machines"
)

// PokeGo API wrapper interface
type IPokeGo interface {
	Root() (models.Root, error)
}

// PokeGo API wrapper client
type PokeGo struct {
	Berries    berries.Berries
	Contests   contests.Contests
	Encounters encounters.Encounters
	Evolution  evolution.Evolution
	Games      games.Games
	Items      items.Items
	Locations  locations.Locations
	Machines   machines.Machines
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		Berries:    berries.NewBerriesGroup(),
		Contests:   contests.NewContestsGroup(),
		Encounters: encounters.NewEncountersGroup(),
		Evolution:  evolution.NewEvolutionGroup(),
		Games:      games.NewGamesGroup(),
		Items:      items.NewItemsGroup(),
		Locations:  locations.NewLocationsGroup(),
		Machines:   machines.NewMachinesGroup(),
	}
}

// Return an instance of API Root list of available resources
func (p PokeGo) Root() (*models.Root, error) {
	root, err := request.GetSpecificResource[models.Root](constants.BaseUrl)
	if err != nil {
		return nil, err
	}
	return root, nil
}
