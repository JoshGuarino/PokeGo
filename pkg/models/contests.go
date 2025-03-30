package models

// Represents a single contest type resource
type ContestType struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	BerryFlavor NamedResource `json:"berry_flavor"`
	Names       []ContestName `json:"names"`
}

// Represents name of a contest in a specific language
type ContestName struct {
	Name     string        `json:"name"`
	Color    string        `json:"color"`
	Language NamedResource `json:"language"`
}

// Represents a single contest effect resource
type ContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

// Represents a single super contest effect resource
type SuperContestEffect struct {
	ID                int             `json:"id"`
	Appeal            int             `json:"appeal"`
	FlavorTextEntries []FlavorText    `json:"flavor_text_entries"`
	Moves             []NamedResource `json:"moves"`
}
