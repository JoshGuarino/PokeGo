package models

// Represents a single ability resource
type Ability struct {
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	IsMainSeries      bool                  `json:"is_main_series"`
	Generation        NamedResource         `json:"generation"`
	Names             []Name                `json:"names"`
	EffectEntries     []VerboseEffect       `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

// Represents the effects an ability has
type AbilityEffectChange struct {
	EffectEntries []Effect      `json:"effect_entries"`
	VersionGroup  NamedResource `json:"version_group"`
}

// Represents the localized flavor text of an ability
type AbilityFlavorText struct {
	FlavorText   string        `json:"flavor_text"`
	Language     NamedResource `json:"language"`
	VersionGroup NamedResource `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Pokemon  NamedResource `json:"pokemon"`
}

// Represents a single characteristic resource
type Characteristic struct {
	ID             int           `json:"id"`
	GeneModulo     int           `json:"gene_modulo"`
	PossibleValues []int         `json:"possible_values"`
	HighestStat    NamedResource `json:"highest_stat"`
	Descriptions   []Description `json:"descriptions"`
}

// Represents a single egg group resource
type EggGroup struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Names          []Name          `json:"names"`
	PokemonSpecies []NamedResource `json:"pokemon_species"`
}

// Represents a single gender resource
type Gender struct {
	ID                    int                    `json:"id"`
	Name                  string                 `json:"name"`
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution  []NamedResource        `json:"required_for_evolution"`
}

// Represents a pokemmon species and the chance its a particular gender
type PokemonSpeciesGender struct {
	Rate           int           `json:"rate"`
	PokemonSpecies NamedResource `json:"pokemon_species"`
}

// Represents a single growth rate resource
type GrowthRate struct {
	ID             int                         `json:"id"`
	Name           string                      `json:"name"`
	Formula        string                      `json:"formula"`
	Descriptions   []Description               `json:"descriptions"`
	Levels         []GrowthRateExperienceLevel `json:"levels"`
	PokemonSpecies []NamedResource             `json:"pokemon_species"`
}

// Represents a level within a growth rate
type GrowthRateExperienceLevel struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`
}

// Represents a single nature resource
type Nature struct {
	ID                         int                         `json:"id"`
	Name                       string                      `json:"name"`
	DecreasedStat              NamedResource               `json:"decreased_stat"`
	IncreasedStat              NamedResource               `json:"increased_stat"`
	HatesFlavor                NamedResource               `json:"hates_flavor"`
	LikesFlavor                NamedResource               `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []Name                      `json:"names"`
}

// Represents how a nature changes the referenced stat
type NatureStatChange struct {
	MaxChange      int           `json:"max_change"`
	PokeathlonStat NamedResource `json:"pokeathlon_stat"`
}

// Represents a move battle style and how likely a pokemon is to use it
type MoveBattleStylePreference struct {
	LowHpPreference  int           `json:"low_hp_preference"`
	HighHpPreference int           `json:"high_hp_preference"`
	MoveBattleStyle  NamedResource `json:"move_battle_style"`
}

// Represents a single pokéathlon stat resource
type PokeathlonStat struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	Names            []Name                         `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

// Represents how a pokéathlon stat affects natures
type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

// Represents how a pokéathlon stat affects a natures
type NaturePokeathlonStatAffect struct {
	MaxChange int           `json:"max_change"`
	Nature    NamedResource `json:"nature"`
}

// Represents a single pokemon resource
type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []NamedResource    `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
	Sprites                PokemonSprites     `json:"sprites"`
	Species                NamedResource      `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
}

// Represents a single pokemon ability
type PokemonAbility struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Ability  NamedResource `json:"ability"`
}

// Represents a single pokemon type
type PokemonType struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

type PokemonFormType struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

// Represents past types of a pokemon in a previous generation
type PokemonTypePast struct {
	Generation NamedResource `json:"generation"`
	Types      []PokemonType `json:"types"`
}

// Represents a held item by a pokemon
type PokemonHeldItem struct {
	Item           NamedResource            `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

// Represents the version details of a held item by a pokemon
type PokemonHeldItemVersion struct {
	Version NamedResource `json:"version"`
	Rarity  int           `json:"rarity"`
}

// Represents a single pokemon move
type PokemonMove struct {
	Move                NamedResource        `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

// Represents the version details of a move learned by a pokemon
type PokemonMoveVersion struct {
	MoveLearnMethod NamedResource `json:"move_learn_method"`
	VersionGroup    NamedResource `json:"version_group"`
	LevelLearnedAt  int           `json:"level_learned_at"`
}

// Represents a single pokemon stat
type PokemonStat struct {
	Stat     NamedResource `json:"stat"`
	Effort   int           `json:"effort"`
	BaseStat int           `json:"base_stat"`
}

// Represents the sprites of a pokemon
type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// Represents a single location area encounter
type LocationAreaEncounter struct {
	LocationArea   NamedResource            `json:"location_area"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

// Represents a single pokemon color resource
type PokemonColor struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Names          []Name          `json:"names"`
	PokemonSpecies []NamedResource `json:"pokemon_species"`
}

// Represents a single pokemon form resource
type PokemonForm struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Order        int                `json:"order"`
	FormOrder    int                `json:"form_order"`
	IsDefault    bool               `json:"is_default"`
	IsBattleOnly bool               `json:"is_battle_only"`
	IsMega       bool               `json:"is_mega"`
	FormName     string             `json:"form_name"`
	Pokemon      NamedResource      `json:"pokemon"`
	Types        []PokemonFormType  `json:"types"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup NamedResource      `json:"version_group"`
	Names        []Name             `json:"names"`
	FormNames    []Name             `json:"form_names"`
}

// Represents a set of sprites for a pokemon form
type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

// Represents a single pokemon habitat resource
type PokemonHabitat struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Names          []Name          `json:"names"`
	PokemonSpecies []NamedResource `json:"pokemon_species"`
}

// Represents a single pokemon shape resource
type PokemonShape struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	AwesomeNames   []AwesomeName   `json:"awesome_names"`
	Names          []Name          `json:"names"`
	PokemonSpecies []NamedResource `json:"pokemon_species"`
}

// Represents scientific name of a pokemon shape
type AwesomeName struct {
	AwesomeName string        `json:"awesome_name"`
	Language    NamedResource `json:"language"`
}

// Represents a single pokemon species resource
type PokemonSpecies struct {
	ID                   int                      `json:"id"`
	Name                 string                   `json:"name"`
	Order                int                      `json:"order"`
	GenderRate           int                      `json:"gender_rate"`
	CaptureRate          int                      `json:"capture_rate"`
	BaseHappiness        int                      `json:"base_happiness"`
	IsBaby               bool                     `json:"is_baby"`
	IsLegendary          bool                     `json:"is_legendary"`
	IsMythical           bool                     `json:"is_mythical"`
	HatchCounter         int                      `json:"hatch_counter"`
	HasGenderDifferences bool                     `json:"has_gender_differences"`
	FormsSwitchable      bool                     `json:"forms_switchable"`
	GrowthRate           NamedResource            `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups            []NamedResource          `json:"egg_groups"`
	Color                NamedResource            `json:"color"`
	Shape                NamedResource            `json:"shape"`
	EvolvesFromSpecies   NamedResource            `json:"evolves_from_species"`
	EvolutionChain       Resource                 `json:"evolution_chain"`
	Habitat              NamedResource            `json:"habitat"`
	Generation           NamedResource            `json:"generation"`
	Names                []Name                   `json:"names"`
	PalParkEncounters    []PalParkEncounterArea   `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText             `json:"flavor_text_entries"`
	FormDescriptions     []Description            `json:"form_descriptions"`
	Genera               []Genus                  `json:"genera"`
	Varieties            []PokemonSpeciesVariety  `json:"varieties"`
}

// Respresents the localized genus of a pokemon species
type Genus struct {
	Genus    string        `json:"genus"`
	Language NamedResource `json:"language"`
}

// Represents a single pokemon species dex entry
type PokemonSpeciesDexEntry struct {
	EntryNumber int           `json:"entry_number"`
	Pokedex     NamedResource `json:"pokedex"`
}

// Represents an encounter for a species in pal park
type PalParkEncounterArea struct {
	BaseScore int           `json:"base_score"`
	Rate      int           `json:"rate"`
	Area      NamedResource `json:"area"`
}

// Respresents a pokemon that exists within a species
type PokemonSpeciesVariety struct {
	IsDefault bool          `json:"is_default"`
	Pokemon   NamedResource `json:"pokemon"`
}

// Represents a single stat resource
type Stat struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	GameIndex        int                  `json:"game_index"`
	IsBattleOnly     bool                 `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics  []Resource           `json:"characteristics"`
	MoveDamageClass  NamedResource        `json:"move_damage_class"`
	Names            []Name               `json:"names"`
}

// Represents a deatil of moves that affect a stat
type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}

// Represents a move that affects a stat
type MoveStatAffect struct {
	Change int           `json:"change"`
	Move   NamedResource `json:"move"`
}

// Represents a detail of natures that affect a stat
type NatureStatAffectSets struct {
	Increase []NamedResource `json:"increase"`
	Decrease []NamedResource `json:"decrease"`
}

// Represents a single type resource
type Type struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	DamageRelations     TypeRelations         `json:"damage_relations"`
	PastDamageRelations []TypeRelationsPast   `json:"past_damage_relations"`
	GameIndices         []GenerationGameIndex `json:"game_indices"`
	Generation          NamedResource         `json:"generation"`
	MoveDamageClass     NamedResource         `json:"move_damage_class"`
	Names               []Name                `json:"names"`
	Pokemon             []TypePokemon         `json:"pokemon"`
}

// Respresents details of a pokemon that has a particular type
type TypePokemon struct {
	Slot    int           `json:"slot"`
	Pokemon NamedResource `json:"pokemon"`
}

// Represents the damage relations of a type
type TypeRelations struct {
	NoDamageTo       []NamedResource `json:"no_damage_to"`
	HalfDamageTo     []NamedResource `json:"half_damage_to"`
	DoubleDamageTo   []NamedResource `json:"double_damage_to"`
	NoDamageFrom     []NamedResource `json:"no_damage_from"`
	HalfDamageFrom   []NamedResource `json:"half_damage_from"`
	DoubleDamageFrom []NamedResource `json:"double_damage_from"`
}

// Represents the damage relations of a type in a previous generation
type TypeRelationsPast struct {
	Generation      NamedResource `json:"generation"`
	DamageRelations TypeRelations `json:"damage_relations"`
}
