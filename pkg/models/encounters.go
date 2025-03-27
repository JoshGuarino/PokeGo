package models

// Represents a single encounter method resource
type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []Name `json:"names"`
}

// Represents a single encounter condition resource
type EncounterCondition struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Names  []Name          `json:"names"`
	Values []NamedResource `json:"values"`
}

// Represents a single encounter condition value resource
type EncounterConditionValue struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Condition NamedResource `json:"condition"`
	Names     []Name        `json:"names"`
}
