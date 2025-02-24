package models

// Represents a single EncounterMethod resource
type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Order int `json:"order"`
}

// Represents a single EncounterCondition resource
type EncounterCondition struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"values"`
}

// Represents a single EncounterConditionValue resource
type EncounterConditionValue struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"values"`
}
