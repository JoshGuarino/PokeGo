package models

// Root is a list of availble resource endpoints
type Root struct {
	Ability                 string `json:"ability"`
	Berry                   string `json:"berry"`
	BerryFirmness           string `json:"berry-firmness"`
	BerryFlavor             string `json:"berry-lavor"`
	Characteristic          string `json:"characteristic"`
	ContestEffect           string `json:"contest-effect"`
	ContestType             string `json:"contest-type"`
	EggGroup                string `json:"egg-group"`
	EncounterCondition      string `json:"encounter-condition"`
	EncounterConditionValue string `json:"encounter-condition-value"`
	EncounterMethod         string `json:"encounter-method"`
	EvolutionChain          string `json:"evolution-chain"`
	EvolutionTrigger        string `json:"evolution-trigger"`
	Gender                  string `json:"gender"`
	Generation              string `json:"generation"`
	GrowthRate              string `json:"growth-rate"`
	Item                    string `json:"item"`
	ItemAttribute           string `json:"item-attribute"`
	ItemCategory            string `json:"item-category"`
	ItemFlingEffect         string `json:"item-fling-effect"`
	ItemPocket              string `json:"item-pocket"`
	Language                string `json:"language"`
	Location                string `json:"location"`
	LocationArea            string `json:"location-area"`
	Machine                 string `json:"machine"`
	Move                    string `json:"move"`
	MoveAilment             string `json:"move-ailment"`
	MoveBattleStyle         string `json:"move-battle-style"`
	MoveCategory            string `json:"move-category"`
	MoveDamageClass         string `json:"move-damage-class"`
	MoveLearnMethod         string `json:"move-learn-method"`
	MoveTarget              string `json:"move-target"`
	Nature                  string `json:"nature"`
	PalParkArea             string `json:"pal-park-area"`
	PokeathlonStat          string `json:"pokeathlon-stat"`
	Pokedex                 string `json:"pokedex"`
	Pokemon                 string `json:"pokemon"`
	PokemonColor            string `json:"pokemon-color"`
	PokemonForm             string `json:"pokemon-form"`
	PokemonHabitat          string `json:"pokemon-habitat"`
	PokemonShape            string `json:"pokemon-shape"`
	PokemonSpecies          string `json:"pokemon-species"`
	Region                  string `json:"region"`
	Stat                    string `json:"stat"`
	SuperContestEffect      string `json:"super-contest-effect"`
	Type                    string `json:"type"`
	Version                 string `json:"version"`
	VersionGroup            string `json:"version-group"`
}

// Represents a ResourceList for an endpoint
type ResourceList struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

// Represents a ResourceList Result
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Represents PaginationOptions for a list endpoint
type PaginationOptions struct {
	Limit  int
	Offest int
}
