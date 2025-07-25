package berries

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Berries group interface
type IBerries interface {
	GetBerry(nameOrId string) (*models.Berry, error)
	GetBerryList(limit int, offset int) (*models.NamedResourceList, error)
	GetBerryURL() string
	GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error)
	GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error)
	GetBerryFirmnessURL() string
	GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error)
	GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error)
	GetBerryFlavorURL() string
}

// Berries group struct
type Berries struct {
	BerryURL         string
	BerryFirmnessURL string
	BerryFlavorURL   string
	Cache            *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Berries resource group initialized")
}

// Return an instance of Berry resource group struct
func NewBerriesGroup() Berries {
	url := endpoints.BaseURL
	return Berries{
		BerryURL:         url + endpoints.Berry,
		BerryFirmnessURL: url + endpoints.BerryFirmness,
		BerryFlavorURL:   url + endpoints.BerryFlavor,
		Cache:            cache.C,
	}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (*models.Berry, error) {
	berry, err := request.GetResource[models.Berry](b.BerryURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

// Return a list of Berry resource
func (b Berries) GetBerryList(limit int, offset int) (*models.NamedResourceList, error) {
	berryList, err := request.GetResourceList[models.NamedResourceList](b.BerryURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryList, nil
}

// Return the Berry resource URL
func (b Berries) GetBerryURL() string {
	return b.BerryURL
}

// Return a single BerryFirmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error) {
	berryFirmness, err := request.GetResource[models.BerryFirmness](b.BerryFirmnessURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

// Return a list of BerryFirmness resource
func (b Berries) GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFirmnessList, err := request.GetResourceList[models.NamedResourceList](b.BerryFirmnessURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFirmnessList, nil

}

// Return the BerryFirmness resource URL
func (b Berries) GetBerryFirmnessURL() string {
	return b.BerryFirmnessURL
}

// Return a single BerryFlavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error) {
	berryFlavor, err := request.GetResource[models.BerryFlavor](b.BerryFlavorURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

// Return a list of BerryFlavor resource
func (b Berries) GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFlavorList, err := request.GetResourceList[models.NamedResourceList](b.BerryFlavorURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFlavorList, nil

}

// Return the BerryFlavor resource URL
func (b Berries) GetBerryFlavorURL() string {
	return b.BerryFlavorURL
}
