package encounters

// Encounters group interface
type IEncounters interface {
}

// Encounters group struct
type Encounters struct{}

// Return an instance of Encounters resource group struct
func NewEncountersGroup() Encounters {
	return Encounters{}
}
