package models

// Represents a single Item resource
type Item struct {
	Attributes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"attributes"`
	BabyTriggerFor interface{} `json:"baby_trigger_for"`
	Category       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"category"`
	Cost          int `json:"cost"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Text         string `json:"text"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	FlingEffect interface{} `json:"fling_effect"`
	FlingPower  interface{} `json:"fling_power"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	HeldByPokemon []interface{} `json:"held_by_pokemon"`
	ID            int           `json:"id"`
	Machines      []interface{} `json:"machines"`
	Name          string        `json:"name"`
	Names         []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Sprites struct {
		Default string `json:"default"`
	} `json:"sprites"`
}

// Represents a single ItemAttribute resource
type ItemAttribute struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID    int `json:"id"`
	Items []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"items"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

// Represents a single ItemCategory resource
type ItemCategory struct {
	ID    int `json:"id"`
	Items []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"items"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pocket struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pocket"`
}

// Represents a single ItemFlingEffect resource
type ItemFlingEffect struct {
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	ID    int `json:"id"`
	Items []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"items"`
	Name string `json:"name"`
}

// Represents a single ItemPocket resource
type ItemPocket struct {
	Categories []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"categories"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}
