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
	EncounterConditionEndpointValue = BaseUrl + "/encounter-condition-value/"
)
