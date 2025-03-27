package models

// Represents a single item resource
type Item struct {
	ID                int                      `json:"id"`
	Name              string                   `json:"name"`
	Cost              int                      `json:"cost"`
	FlingPower        int                      `json:"fling_power"`
	FlingEffect       NamedResource            `json:"fling_effect"`
	Attributes        []NamedResource          `json:"attributes"`
	Category          NamedResource            `json:"category"`
	EffectEntries     []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices       []GenerationGameIndex    `json:"game_indices"`
	Names             []Name                   `json:"names"`
	Sprites           ItemSprites              `json:"sprites"`
	HeldByPokemon     []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerFor    Resource                 `json:"baby_trigger_for"`
	Machines          []MachineVersionDetail   `json:"machines"`
}

// Represents the default sprite for an item
type ItemSprites struct {
	Default string `json:"default"`
}

// Represents a pokemon that holds an item
type ItemHolderPokemon struct {
	Pokemon        NamedResource                    `json:"pokemon"`
	VersionDetails []ItemHolderPokemonVersionDetail `json:"version_details"`
}

// Represents the version details for a pokemon that holds an item
type ItemHolderPokemonVersionDetail struct {
	Rarity  int           `json:"rarity"`
	Version NamedResource `json:"version"`
}

// Represents a single item attribute resource
type ItemAttribute struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Items        []NamedResource `json:"items"`
	Names        []Name          `json:"names"`
	Descriptions []Description   `json:"descriptions"`
}

// Represents a single item category resource
type ItemCategory struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Items  []NamedResource `json:"items"`
	Names  []Name          `json:"names"`
	Pocket NamedResource   `json:"pocket"`
}

// Represents a single item fling effect resource
type ItemFlingEffect struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	EffectEntries []Effect        `json:"effect_entries"`
	Items         []NamedResource `json:"items"`
}

// Represents a single item pocket resource
type ItemPocket struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Categories []NamedResource `json:"categories"`
	Names      []Name          `json:"names"`
}
