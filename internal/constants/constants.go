package constants

const (
	// PokeApi base URL
	BaseUrl = "https://pokeapi.co/api/v2"

	// Berry group resource endpoints
	BerryEndpoint         = BaseUrl + "/berry/"
	BerryFirmnessEndpoint = BaseUrl + "/berry-firmness/"
	BerryFlavorEndpoint   = BaseUrl + "/berry-flavor/"

	// Contests group resource endpoints
	ContestTypeEndpoint        = BaseUrl + "/contest-type/"
	ContestEffectEndpoint      = BaseUrl + "/contest-effect/"
	SuperContestEffectEndpoint = BaseUrl + "/super-contest-effect/"

	// Encounters group resource endpoints
	EncounterMethodEndpoint         = BaseUrl + "/encounter-method/"
	EncounterConditionEndpoint      = BaseUrl + "/encounter-condition/"
	EncounterConditionValueEndpoint = BaseUrl + "/encounter-condition-value/"

	// Evolution group resource endpoints
	EvolutionChainEndpoint   = BaseUrl + "/evolution-chain/"
	EvolutionTriggerEndpoint = BaseUrl + "/evolution-trigger/"

	// Games group resource endpoints
	GenerationEndpoint   = BaseUrl + "/generation/"
	PokedexEndpoint      = BaseUrl + "/pokedex/"
	VersionEndpoint      = BaseUrl + "/version/"
	VersionGroupEndpoint = BaseUrl + "/version-group/"

	// Items group resource endpoints
	ItemEndpoint            = BaseUrl + "/item/"
	ItemAttributeEndpoint   = BaseUrl + "/item-attribute/"
	ItemCategoryEndpoint    = BaseUrl + "/item-category/"
	ItemFlingEffectEndpoint = BaseUrl + "/item-fling-effect/"
	ItemPocketEndpoint      = BaseUrl + "/item-pocket/"

	// Locations group resource endpoints
	LocationEndpoint     = BaseUrl + "/location/"
	LocationAreaEndpoint = BaseUrl + "/location-area/"
	PalParkAreaEndpoint  = BaseUrl + "/pal-park-area/"
	RegionEndpoint       = BaseUrl + "/region/"

	// Machines group resource endpoints
	MachineEndpoint = BaseUrl + "/machine/"

	// Moves group resource endpoints
	MoveEndpoint            = BaseUrl + "/move/"
	MoveAilmentEndpoint     = BaseUrl + "/move-ailment/"
	MoveBattleStyleEndpoint = BaseUrl + "/move-battle-style/"
	MoveCategoryEndpoint    = BaseUrl + "/move-category/"
	MoveDamageClassEndpoint = BaseUrl + "/move-damage-class/"
	MoveLearnMethodEndpoint = BaseUrl + "/move-learn-method/"
	MoveTargetEndpoint      = BaseUrl + "/move-target/"

	// Pokemon group resource endpoints
	AbilityEndpoint        = BaseUrl + "/ability/"
	CharacteristicEndpoint = BaseUrl + "/characteristic/"
	EggGroupEndpoint       = BaseUrl + "/egg-group/"
	GenderEndpoint         = BaseUrl + "/gender/"
	GrowthRateEndpoint     = BaseUrl + "/growth-rate/"
	NatureEndpoint         = BaseUrl + "/nature/"
	PokeathlonStatEndpoint = BaseUrl + "/pokeathlon-stat/"
	PokemonEndpoint        = BaseUrl + "/pokemon/"
	PokemonColorEndpoint   = BaseUrl + "/pokemon-color/"
	PokemonFormEndpoint    = BaseUrl + "/pokemon-form/"
	PokemonHabitatEndpoint = BaseUrl + "/pokemon-habitat/"
	PokemonShapeEndpoint   = BaseUrl + "/pokemon-shape/"
	PokemonSpeciesEndpoint = BaseUrl + "/pokemon-species/"
	StatEndpoint           = BaseUrl + "/stat/"
	TypeEndpoint           = BaseUrl + "/type/"

	// Utility group resource endpoints
	LanguageEndpoint = BaseUrl + "/language/"
)
