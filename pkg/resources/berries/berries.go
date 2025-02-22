package berries

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Berries group interface
type IBerries interface {
	GetBerry(nameOrId string) (*models.Berry, error)
	GetBerryList() (*models.ResourceList, error)
	GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error)
	GetBerryFirmnessList() (*models.ResourceList, error)
	GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error)
	GetBerryFlavorList() (*models.ResourceList, error)
}

// Berries group struct
type Berries struct{}

// Return an instance of Berry resource group struct
func NewBerriesGroup() Berries {
	return Berries{}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (*models.Berry, error) {
	berry, err := request.GetSpecificResource[models.Berry](constants.BerryEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

// Return a list of Berry resource
func (b Berries) GetBerryList() (*models.ResourceList, error) {
	berryList, err := request.GetResourceList(constants.BerryEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryList, nil
}

// Return a single BerryFirmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error) {
	berryFirmness, err := request.GetSpecificResource[models.BerryFirmness](constants.BerryFirmnessEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

// Return a list of BerryFirmness resource
func (b Berries) GetBerryFirmnessList() (*models.ResourceList, error) {
	berryFirmnessList, err := request.GetResourceList(constants.BerryFirmnessEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryFirmnessList, nil

}

// Return a single BerryFlavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error) {
	berryFlavor, err := request.GetSpecificResource[models.BerryFlavor](constants.BerryFlavorEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

// Return a list of BerryFlavor resource
func (b Berries) GetBerryFlavorList() (*models.ResourceList, error) {
	berryFlavorList, err := request.GetResourceList(constants.BerryFlavorEndpoint)
	if err != nil {
		return nil, err
	}
	return &berryFlavorList, nil

}
