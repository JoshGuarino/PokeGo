package berries

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
)

// Berries group interface
type IBerries interface {
	GetBerry(nameOrId string) (*Berry, error)
	GetBerryList() (*request.ResourceList, error)
	GetBerryFirmness(nameOrId string) (*BerryFirmness, error)
	GetBerryFirmnessList() (*request.ResourceList, error)
	GetBerryFlavor(nameOrId string) (*BerryFlavor, error)
	GetBerryFlavorList() (*request.ResourceList, error)
}

// Berries group struct
type Berries struct{}

// Return an instance of Berry resource group struct
func NewBerriesGroup() Berries {
	return Berries{}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (*Berry, error) {
	berry, err := request.GetSpecificResource[Berry](constants.BerryEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

// Return a list of Berry resource
func (b Berries) GetBerryList() (*request.ResourceList, error) {
	berryList, err := request.GetResourceList(constants.BerryEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryList, nil
}

// Return a single BerryFirmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (*BerryFirmness, error) {
	berryFirmness, err := request.GetSpecificResource[BerryFirmness](constants.BerryFirmnessEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

// Return a list of BerryFirmness resource
func (b Berries) GetBerryFirmnessList() (*request.ResourceList, error) {
	berryFirmnessList, err := request.GetResourceList(constants.BerryFirmnessEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryFirmnessList, nil

}

// Return a single BerryFlavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (*BerryFlavor, error) {
	berryFlavor, err := request.GetSpecificResource[BerryFlavor](constants.BerryFlavorEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

// Return a list of BerryFlavor resource
func (b Berries) GetBerryFlavorList() (*request.ResourceList, error) {
	berryFlavorList, err := request.GetResourceList(constants.BerryFlavorEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryFlavorList, nil

}
