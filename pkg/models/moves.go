package models

// Represents a single move resource
type Move struct {
	ID                 int                    `json:"id"`
	Name               string                 `json:"name"`
	Accuracy           int                    `json:"accuracy"`
	EffectChance       int                    `json:"effect_chance"`
	PP                 int                    `json:"pp"`
	Priority           int                    `json:"priority"`
	Power              int                    `json:"power"`
	ContestCombos      ContestComboSets       `json:"contest_combos"`
	ContestType        NamedResource          `json:"contest_type"`
	ContestEffect      Resource               `json:"contest_effect"`
	DamageClass        NamedResource          `json:"damage_class"`
	EffectEntries      []VerboseEffect        `json:"effect_entries"`
	EffectChanges      []AbilityEffectChange  `json:"effect_changes"`
	LearnedByPokemon   []NamedResource        `json:"learned_by_pokemon"`
	FlavorTextEntries  []MoveFlavorText       `json:"flavor_text_entries"`
	Generation         NamedResource          `json:"generation"`
	Machines           []MachineVersionDetail `json:"machines"`
	Meta               MoveMetaData           `json:"meta"`
	Names              []Name                 `json:"names"`
	PastValues         []PastMoveStatValues   `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect Resource               `json:"super_contest_effect"`
	Target             NamedResource          `json:"target"`
	Type               NamedResource          `json:"type"`
}

// Respresents a set of contest combos
type ContestComboSets struct {
	Normal ContestComboDetail `json:"normal"`
	Super  ContestComboDetail `json:"super"`
}

// Represents the details of a contest combo
type ContestComboDetail struct {
	UseBefore []NamedResource `json:"use_before"`
	UseAfter  []NamedResource `json:"use_after"`
}

// Represents the localized flavor text of a move
type MoveFlavorText struct {
	FlavorText   string        `json:"flavor_text"`
	Language     NamedResource `json:"language"`
	VersionGroup NamedResource `json:"version_group"`
}

// Represents meta data about a move
type MoveMetaData struct {
	Ailment       NamedResource `json:"ailment"`
	Category      NamedResource `json:"category"`
	MinHits       int           `json:"min_hits"`
	MaxHits       int           `json:"max_hits"`
	MinTurns      int           `json:"min_turns"`
	MaxTurns      int           `json:"max_turns"`
	Drain         int           `json:"drain"`
	Healing       int           `json:"healing"`
	CritRate      int           `json:"crit_rate"`
	AilmentChance int           `json:"ailment_chance"`
	FlinchChance  int           `json:"flinch_chance"`
	StatChance    int           `json:"stat_chance"`
}

// Represents a move stat change
type MoveStatChange struct {
	Change int           `json:"change"`
	Stat   NamedResource `json:"stat"`
}

// Represents past move stat values
type PastMoveStatValues struct {
	Accuracy      int             `json:"accuracy"`
	EffectChance  int             `json:"effect_chance"`
	Power         int             `json:"power"`
	PP            int             `json:"pp"`
	EffectEntries []VerboseEffect `json:"effect_entries"`
	Type          NamedResource   `json:"type"`
	VersionGroup  NamedResource   `json:"version_group"`
}

// Represents a single move ailment resource
type MoveAilment struct {
	ID    int             `json:"id"`
	Name  string          `json:"name"`
	Moves []NamedResource `json:"moves"`
	Names []Name          `json:"names"`
}

// Represents a single move battle style resource
type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

// Represents a single move category resource
type MoveCategory struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Moves        []NamedResource `json:"moves"`
	Descriptions []Description   `json:"descriptions"`
}

// Represents a single move damage class resource
type MoveDamageClass struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Descriptions []Description   `json:"descriptions"`
	Moves        []NamedResource `json:"moves"`
	Names        []Name          `json:"names"`
}

// Represents a single move learn method resource
type MoveLearnMethod struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Descriptions  []Description   `json:"descriptions"`
	Names         []Name          `json:"names"`
	VersionGroups []NamedResource `json:"version_groups"`
}

// Represents a single move target resource
type MoveTarget struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Descriptions []Description   `json:"descriptions"`
	Moves        []NamedResource `json:"moves"`
	Names        []Name          `json:"names"`
}
