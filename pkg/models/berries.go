package models

// Represents a single berry resource
type Berry struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	GrowthTime       int              `json:"growth_time"`
	MaxHarvest       int              `json:"max_harvest"`
	NaturalGiftPower int              `json:"natural_gift_power"`
	Size             int              `json:"size"`
	Smoothness       int              `json:"smoothness"`
	SoilDryness      int              `json:"soil_dryness"`
	Firmness         NamedResource    `json:"firmness"`
	Flavors          []BerryFlavorMap `json:"flavors"`
	Item             NamedResource    `json:"item"`
	NaturalGiftType  NamedResource    `json:"natural_gift_type"`
}

// Represents a flavor that can be found in berries
type BerryFlavorMap struct {
	Potency int           `json:"potency"`
	Flavor  NamedResource `json:"flavor"`
}

// Represents a single berry firmness resource
type BerryFirmness struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Berries []NamedResource `json:"berries"`
	Names   []Name          `json:"names"`
}

// Represents a single berry flavor resource
type BerryFlavor struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Berries []FlavorBerryMap `json:"berries"`
	Names   []Name           `json:"names"`
}

// Represents a berry that has a particular flavor
type FlavorBerryMap struct {
	Potency int           `json:"potency"`
	Berry   NamedResource `json:"berry"`
}
