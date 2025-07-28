package endpoints

import (
	"fmt"
	"testing"
)

const (
	// PokeApi base URL
	ProdBaseURL = "https://pokeapi.co/api/v2"

	// Staging base URL
	StageBaseURL = "https://staging.pokeapi.co/api/v2"

	// Berry group resource endpoints
	Berry         = "/berry/"
	BerryFirmness = "/berry-firmness/"
	BerryFlavor   = "/berry-flavor/"

	// Contests group resource endpoints
	ContestType        = "/contest-type/"
	ContestEffect      = "/contest-effect/"
	SuperContestEffect = "/super-contest-effect/"

	// Encounters group resource endpoints
	EncounterMethod         = "/encounter-method/"
	EncounterCondition      = "/encounter-condition/"
	EncounterConditionValue = "/encounter-condition-value/"

	// Evolution group resource endpoints
	EvolutionChain   = "/evolution-chain/"
	EvolutionTrigger = "/evolution-trigger/"

	// Games group resource endpoints
	Generation   = "/generation/"
	Pokedex      = "/pokedex/"
	Version      = "/version/"
	VersionGroup = "/version-group/"

	// Items group resource endpoints
	Item            = ProdBaseURL + "/item/"
	ItemAttribute   = ProdBaseURL + "/item-attribute/"
	ItemCategory    = ProdBaseURL + "/item-category/"
	ItemFlingEffect = ProdBaseURL + "/item-fling-effect/"
	ItemPocket      = ProdBaseURL + "/item-pocket/"

	// Locations group resource endpoints
	Location     = ProdBaseURL + "/location/"
	LocationArea = ProdBaseURL + "/location-area/"
	PalParkArea  = ProdBaseURL + "/pal-park-area/"
	Region       = ProdBaseURL + "/region/"

	// Machines group resource endpoints
	Machine = "/machine/"

	// Moves group resource endpoints
	Move            = ProdBaseURL + "/move/"
	MoveAilment     = ProdBaseURL + "/move-ailment/"
	MoveBattleStyle = ProdBaseURL + "/move-battle-style/"
	MoveCategory    = ProdBaseURL + "/move-category/"
	MoveDamageClass = ProdBaseURL + "/move-damage-class/"
	MoveLearnMethod = ProdBaseURL + "/move-learn-method/"
	MoveTarget      = ProdBaseURL + "/move-target/"

	// Pokemon group resource endpoints
	Ability        = ProdBaseURL + "/ability/"
	Characteristic = ProdBaseURL + "/characteristic/"
	EggGroup       = ProdBaseURL + "/egg-group/"
	Gender         = ProdBaseURL + "/gender/"
	GrowthRate     = ProdBaseURL + "/growth-rate/"
	Nature         = ProdBaseURL + "/nature/"
	PokeathlonStat = ProdBaseURL + "/pokeathlon-stat/"
	Pokemon        = ProdBaseURL + "/pokemon/"
	PokemonColor   = ProdBaseURL + "/pokemon-color/"
	PokemonForm    = ProdBaseURL + "/pokemon-form/"
	PokemonHabitat = ProdBaseURL + "/pokemon-habitat/"
	PokemonShape   = ProdBaseURL + "/pokemon-shape/"
	PokemonSpecies = ProdBaseURL + "/pokemon-species/"
	Stat           = ProdBaseURL + "/stat/"
	Type           = ProdBaseURL + "/type/"

	// Utility group resource endpoints
	Language = "/language/"
)

// PokeAPI base URL set to prod env by default
var BaseURL string = ProdBaseURL

// Initialize function
func init() {
	// If testing, use the staging URL
	if testing.Testing() {
		BaseURL = StageBaseURL
	}
	fmt.Println("Base URL initialized ->", BaseURL)
}
