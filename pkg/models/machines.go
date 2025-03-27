package models

// Represents a single machine resource
type Machine struct {
	ID           int           `json:"id"`
	Item         NamedResource `json:"item"`
	Move         NamedResource `json:"move"`
	VersionGroup NamedResource `json:"version_group"`
}
