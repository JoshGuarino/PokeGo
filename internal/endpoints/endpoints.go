package endpoints

const (
	// PokeApi base URL
	BaseUrl = "https://pokeapi.co/api/v2"

	// Berry group resource endpoints
	Berry         = BaseUrl + "/berry/"
	BerryFirmness = BaseUrl + "/berry-firmness/"
	BerryFlavor   = BaseUrl + "/berry-flavor/"

	// Contests group resource endpoints
	ContestType        = BaseUrl + "/contest-type/"
	ContestEffect      = BaseUrl + "/contest-effect/"
	SuperContestEffect = BaseUrl + "/super-contest-effect/"

	// Encounters group resource endpoints
	EncounterMethod         = BaseUrl + "/encounter-method/"
	EncounterCondition      = BaseUrl + "/encounter-condition/"
	EncounterConditionValue = BaseUrl + "/encounter-condition-value/"

	// Evolution group resource endpoints
	EvolutionChain   = BaseUrl + "/evolution-chain/"
	EvolutionTrigger = BaseUrl + "/evolution-trigger/"

	// Games group resource endpoints
	Generation   = BaseUrl + "/generation/"
	Pokedex      = BaseUrl + "/pokedex/"
	Version      = BaseUrl + "/version/"
	VersionGroup = BaseUrl + "/version-group/"

	// Items group resource endpoints
	Item            = BaseUrl + "/item/"
	ItemAttribute   = BaseUrl + "/item-attribute/"
	ItemCategory    = BaseUrl + "/item-category/"
	ItemFlingEffect = BaseUrl + "/item-fling-effect/"
	ItemPocket      = BaseUrl + "/item-pocket/"

	// Locations group resource endpoints
	Location     = BaseUrl + "/location/"
	LocationArea = BaseUrl + "/location-area/"
	PalParkArea  = BaseUrl + "/pal-park-area/"
	Region       = BaseUrl + "/region/"

	// Machines group resource endpoints
	Machine = BaseUrl + "/machine/"

	// Moves group resource endpoints
	Move            = BaseUrl + "/move/"
	MoveAilment     = BaseUrl + "/move-ailment/"
	MoveBattleStyle = BaseUrl + "/move-battle-style/"
	MoveCategory    = BaseUrl + "/move-category/"
	MoveDamageClass = BaseUrl + "/move-damage-class/"
	MoveLearnMethod = BaseUrl + "/move-learn-method/"
	MoveTarget      = BaseUrl + "/move-target/"

	// Pokemon group resource endpoints
	Ability        = BaseUrl + "/ability/"
	Characteristic = BaseUrl + "/characteristic/"
	EggGroup       = BaseUrl + "/egg-group/"
	Gender         = BaseUrl + "/gender/"
	GrowthRate     = BaseUrl + "/growth-rate/"
	Nature         = BaseUrl + "/nature/"
	PokeathlonStat = BaseUrl + "/pokeathlon-stat/"
	Pokemon        = BaseUrl + "/pokemon/"
	PokemonColor   = BaseUrl + "/pokemon-color/"
	PokemonForm    = BaseUrl + "/pokemon-form/"
	PokemonHabitat = BaseUrl + "/pokemon-habitat/"
	PokemonShape   = BaseUrl + "/pokemon-shape/"
	PokemonSpecies = BaseUrl + "/pokemon-species/"
	Stat           = BaseUrl + "/stat/"
	Type           = BaseUrl + "/type/"

	// Utility group resource endpoints
	Language = BaseUrl + "/language/"
)
