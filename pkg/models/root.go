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

// Represents a resource list for a list endpoint
type ResourceList struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Resource `json:"results"`
}

// Represents a unnamed resource
type Resource struct {
	URL string `json:"url"`
}

// Represents a named resource list for a list endpoint
type NamedResourceList struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []NamedResource `json:"results"`
}

// Represents a named resource
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Represents the localized description for a resource in specific Language
type Description struct {
	Description string        `json:"description"`
	Language    NamedResource `json:"language"`
}

// Represents the localized effect text for a resource in specific language
type Effect struct {
	Effect   string        `json:"effect"`
	Language NamedResource `json:"language"`
}

// Represents the detailed information about a specific encounter
type Encounter struct {
	MinLevel        int             `json:"min_level"`
	MaxLevel        int             `json:"max_level"`
	ConditionValues []NamedResource `json:"condition_values"`
	Chance          int             `json:"chance"`
	Method          NamedResource   `json:"method"`
}

// Represents the localized flavor text for a resource in specific language
type FlavorText struct {
	FlavorText string        `json:"flavor_text"`
	Language   NamedResource `json:"language"`
	Version    NamedResource `json:"version"`
}

// Represents a game index reference for a resource
type GenerationGameIndex struct {
	GameIndex  int           `json:"game_index"`
	Gereration NamedResource `json:"generation"`
}

// Represents a single machine that teaches a move
type MachineVersionDetail struct {
	Machine      Resource      `json:"machine"`
	VersionGroup NamedResource `json:"version_group"`
}

// Represents a localized name for a resource in specific language
type Name struct {
	Name     string        `json:"name"`
	Language NamedResource `json:"language"`
}

// Represents long description of an effect
type VerboseEffect struct {
	Effect      string        `json:"effect"`
	ShortEffect string        `json:"short_effect"`
	Language    NamedResource `json:"language"`
}

// Represents encounter details for a specific version
type VersionEncounterDetail struct {
	Version          NamedResource `json:"version"`
	MaxChance        int           `json:"max_chance"`
	EncounterDetails []Encounter   `json:"encounter_details"`
}

// Represents an game index reference for a specific version
type VersionGameIndex struct {
	GameIndex int           `json:"game_index"`
	Version   NamedResource `json:"version"`
}

// Respresents the flavor text for a version group
type VersionGroupFlavorText struct {
	Text         string        `json:"text"`
	Language     NamedResource `json:"language"`
	VersionGroup NamedResource `json:"version_group"`
}
