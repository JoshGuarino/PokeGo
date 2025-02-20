package constants

const (
	// PokeApi base URL
	BaseUrl = "https://pokeapi.co/api/v2"

	// Berry group resource endpoints
	BerryEndpoint         = BaseUrl + "/berry/"
	BerryFirmnessEndpoint = BaseUrl + "/berry-firmness/"
	BerryFlavorEndpoint   = BaseUrl + "/berry-flavor/"

	// Contests group resource endpoints
	contestTypeEndpoint        = BaseUrl + "/contest-type/"
	contestEffectEndpoint      = BaseUrl + "/contest-effect/"
	superContestEffectEndpoint = BaseUrl + "/super-contest-effect/"
)
