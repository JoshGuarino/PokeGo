package models

// Represents a single location resource
type Location struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Region      NamedResource         `json:"region"`
	Names       []Name                `json:"names"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedResource       `json:"areas"`
}

// Represents a single location area resource
type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedResource         `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

// Represents a method to encounter pokemon for an area
type EncounterMethodRate struct {
	EncounterMethod NamedResource             `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

// Represents a version detail for an encounter method rate
type EncounterVersionDetails struct {
	Rate    int           `json:"rate"`
	Version NamedResource `json:"version"`
}

// Respresents the pokemon that can be encountered in an area
type PokemonEncounter struct {
	Pokemon        NamedResource            `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

// Represents a single pal park area resource
type PalParkArea struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

// Represents a pokemon species that can be encountered in a pal park area
type PalParkEncounterSpecies struct {
	BaseScore      int           `json:"base_score"`
	Rate           int           `json:"rate"`
	PokemonSpecies NamedResource `json:"pokemon_species"`
}

// Represents a single region resource
type Region struct {
	ID             int             `json:"id"`
	Locations      []NamedResource `json:"locations"`
	Name           string          `json:"name"`
	Names          []Name          `json:"names"`
	MainGeneration NamedResource   `json:"main_generation"`
	Pokedexes      []NamedResource `json:"pokedexes"`
	VersionGroups  []NamedResource `json:"version_groups"`
}
