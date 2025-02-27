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
)
