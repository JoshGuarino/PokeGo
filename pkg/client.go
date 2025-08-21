package pokego

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/internal/logger"
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
	"github.com/JoshGuarino/PokeGo/pkg/resources/moves"
	"github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
	"github.com/JoshGuarino/PokeGo/pkg/resources/utility"
)

// PokeGo API wrapper interface
type IPokeGo interface {
	Root() (*models.Root, error)
	GetBaseURL() string
}

// PokeGo API wrapper client
type PokeGo struct {
	// Chache and environmen
	Cache *cache.Cache
	Env   *environment.Environment
	Log   *logger.Logger

	// resources
	Berries    berries.Berries
	Contests   contests.Contests
	Encounters encounters.Encounters
	Evolution  evolution.Evolution
	Games      games.Games
	Items      items.Items
	Locations  locations.Locations
	Machines   machines.Machines
	Moves      moves.Moves
	Pokemon    pokemon.Pokemon
	Utility    utility.Utility
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		// initialize cache, environment, and logger
		Cache: cache.CACHE,
		Env:   environment.ENV,
		Log:   logger.LOG,

		// initialize resources
		Berries:    berries.NewBerriesGroup(),
		Contests:   contests.NewContestsGroup(),
		Encounters: encounters.NewEncountersGroup(),
		Evolution:  evolution.NewEvolutionGroup(),
		Games:      games.NewGamesGroup(),
		Items:      items.NewItemsGroup(),
		Locations:  locations.NewLocationsGroup(),
		Machines:   machines.NewMachinesGroup(),
		Moves:      moves.NewMovesGroup(),
		Pokemon:    pokemon.NewPokemonGroup(),
		Utility:    utility.NewUtilityGroup(),
	}
}

// Return an instance of API Root list of available resources
func (p PokeGo) Root() (*models.Root, error) {
	root, err := request.GetResource[models.Root](p.GetBaseURL())
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Return the base URL for PokeAPI
func (p PokeGo) GetBaseURL() string {
	return p.Env.URL()
}
