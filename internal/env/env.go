package env

import (
	"testing"

	"github.com/charmbracelet/log"
)

const (
	// Environment domains
	prodDomain  = "prod"
	stageDomain = "stage"

	// Base URLs
	prodURL  = "https://pokeapi.co/api/v2"
	stageURL = "https://staging.pokeapi.co/api/v2"

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

// Environment interface
type IEnv interface {
	GetDomain() string
	GetURL() string
	SetStage()
	SetProd()
}

// Environment struct
type Env struct {
	url    string
	domain string
}

// Environment global variable defaulting to production
var ENV *Env = prodEnv()

// Initialize function
func init() {
	// Set environment to stage if testing
	if testing.Testing() {
		ENV = stageEnv()
	}
	log.Info("Environment initialized", "domain", ENV.Domain(), "url", ENV.URL())
}

// Return an instance of Env struct
// pointing to the staging environment
func stageEnv() *Env {
	return &Env{
		url:    stageURL,
		domain: stageDomain,
	}
}

// Return an instance of Env struct
// pointing to the production environment
func prodEnv() *Env {
	return &Env{
		url:    prodURL,
		domain: prodDomain,
	}
}

// Return the environment domain
func (e *Env) Domain() string {
	return e.domain
}

// Return the url for the environment
func (e *Env) URL() string {
	return e.url
}

// Set the environment to stage
func (e *Env) SetStage() {
	e.domain = stageDomain
	e.url = stageURL
}

// Set the environment to production
func (e *Env) SetProd() {
	e.domain = prodDomain
	e.url = prodURL
}
