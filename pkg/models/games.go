package models

// Represents a single generation resource
type Generation struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Abilities      []NamedResource `json:"abilities"`
	Names          []Name          `json:"names"`
	MainRegion     NamedResource   `json:"main_region"`
	Moves          []NamedResource `json:"moves"`
	PokemonSpecies []NamedResource `json:"pokemon_species"`
	Types          []NamedResource `json:"types"`
	VersionGroups  []NamedResource `json:"version_groups"`
}

// Represents a single pokedex resource
type Pokedex struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	IsMainSeries   bool            `json:"is_main_series"`
	Descriptions   []Description   `json:"descriptions"`
	Names          []Name          `json:"names"`
	PokemonEntries []PokemonEntry  `json:"pokemon_entries"`
	Region         NamedResource   `json:"region"`
	VersionGroups  []NamedResource `json:"version_groups"`
}

// Represents a pokemon in the specific pokedex
type PokemonEntry struct {
	EntryNumber    int           `json:"entry_number"`
	PokemonSpecies NamedResource `json:"pokemon_species"`
}

// Represents a single version resource
type Version struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Names        []Name        `json:"names"`
	VersionGroup NamedResource `json:"version_group"`
}

// Represents is a single version group resource
type VersionGroup struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Order            int             `json:"order"`
	Generation       NamedResource   `json:"generation"`
	MoveLearnMethods []NamedResource `json:"move_learn_methods"`
	Pokedexes        []NamedResource `json:"pokedexes"`
	Regions          []NamedResource `json:"regions"`
	Versions         []NamedResource `json:"versions"`
}
