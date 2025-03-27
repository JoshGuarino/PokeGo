package models

// Represents a single language resource
type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Iso639   string `json:"iso639"`
	Iso3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}
