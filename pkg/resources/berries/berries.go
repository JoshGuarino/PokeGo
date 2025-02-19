package berries

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
)

// Berry group interface
type IBerries interface {
	GetBerry(nameOrId string) (Berry, error)
}

// Berry group struct
type Berries struct{}

// Return an instance of Berry resource group struct
func NewBerriesGroup() Berries {
	return Berries{}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (Berry, error) {
	berry, err := request.GetSpecificResource[Berry](constants.BerryEndpoint + nameOrId)
	if err != nil {
		return Berry{}, err
	}
	return *berry, nil
}

// Return a list of Berry resource
func GetBerryList() {
}

// Return a single Berry Firmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (BerryFirmness, error) {
	berryFirmness, err := request.GetSpecificResource[BerryFirmness](constants.BerryFirmnessEndpoint + nameOrId)
	if err != nil {
		return BerryFirmness{}, err
	}
	return *berryFirmness, nil
}

// Return a list of Berry Firmness resource
func GetBerryFirmnessList() {
}

// Return a single Berry Flavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (BerryFlavor, error) {
	berryFlavor, err := request.GetSpecificResource[BerryFlavor](constants.BerryFlavorEndpoint + nameOrId)
	if err != nil {
		return BerryFlavor{}, err
	}
	return *berryFlavor, nil

}

// Return a list of Berry Flavor resource
func GetBerryFlavorList() {
}
