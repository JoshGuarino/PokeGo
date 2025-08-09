package endpoints

import (
	"testing"

	"github.com/charmbracelet/log"
)

const (
	// Base URLs
	ProdBaseURL  = "https://pokeapi.co/api/v2"
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
	Item            = "/item/"
	ItemAttribute   = "/item-attribute/"
	ItemCategory    = "/item-category/"
	ItemFlingEffect = "/item-fling-effect/"
	ItemPocket      = "/item-pocket/"

	// Locations group resource endpoints
	Location     = "/location/"
	LocationArea = "/location-area/"
	PalParkArea  = "/pal-park-area/"
	Region       = "/region/"

	// Machines group resource endpoints
	Machine = "/machine/"

	// Moves group resource endpoints
	Move            = "/move/"
	MoveAilment     = "/move-ailment/"
	MoveBattleStyle = "/move-battle-style/"
	MoveCategory    = "/move-category/"
	MoveDamageClass = "/move-damage-class/"
	MoveLearnMethod = "/move-learn-method/"
	MoveTarget      = "/move-target/"

	// Pokemon group resource endpoints
	Ability        = "/ability/"
	Characteristic = "/characteristic/"
	EggGroup       = "/egg-group/"
	Gender         = "/gender/"
	GrowthRate     = "/growth-rate/"
	Nature         = "/nature/"
	PokeathlonStat = "/pokeathlon-stat/"
	Pokemon        = "/pokemon/"
	PokemonColor   = "/pokemon-color/"
	PokemonForm    = "/pokemon-form/"
	PokemonHabitat = "/pokemon-habitat/"
	PokemonShape   = "/pokemon-shape/"
	PokemonSpecies = "/pokemon-species/"
	Stat           = "/stat/"
	Type           = "/type/"

	// Utility group resource endpoints
	Language = "/language/"
)

// PokeAPI base URL set to Production environment by default
var BaseURL string = ProdBaseURL
var Env string = "prod"

// Initialize function
func init() {
	// If testing, use the staging environment
	if testing.Testing() {
		BaseURL = StageBaseURL
		Env = "stage"
	}
	log.Info("Base URL initialized", "Env", Env, "BaseURL", BaseURL)
}
