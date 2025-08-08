package pokego

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
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
	"github.com/charmbracelet/log"
)

// PokeGo API wrapper interface
type IPokeGo interface {
	Root() (*models.Root, error)
	GetBaseURL() string
}

// PokeGo API wrapper client
type PokeGo struct {
	// base URL and cache
	BaseURL string
	Cache   *cache.Cache

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

// Initialize function
func init() {
	log.Info("PokeGo API wrapper initialized")
}

// Return an instance of the PokeGo API wrapper client
func NewClient() PokeGo {
	return PokeGo{
		// initialize base URL and cache
		BaseURL: endpoints.BaseURL,
		Cache:   cache.C,

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
	root, err := request.GetResource[models.Root](p.BaseURL)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Return the base URL for PokeAPI
func (p PokeGo) GetBaseURL() string {
	return p.BaseURL
}
