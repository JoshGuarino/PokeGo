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
	GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error)
	GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error)
	GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error)
	GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error)
}

// Berries group struct
type Berries struct {
	BerryUrl         string
	BerryFirmnessUrl string
	BerryFlavorUrl   string
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
		BerryUrl:         url + endpoints.Berry,
		BerryFirmnessUrl: url + endpoints.BerryFirmness,
		BerryFlavorUrl:   url + endpoints.BerryFlavor,
		Cache:            cache.C,
	}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (*models.Berry, error) {
	berry, err := request.GetResource[models.Berry](b.BerryUrl + nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

// Return a list of Berry resource
func (b Berries) GetBerryList(limit int, offset int) (*models.NamedResourceList, error) {
	berryList, err := request.GetResourceList[models.NamedResourceList](b.BerryUrl, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryList, nil
}

// Return the the Berry resource URL
func (b Berries) GetBerryUrl() string {
	return b.BerryUrl
}

// Return a single BerryFirmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error) {
	berryFirmness, err := request.GetResource[models.BerryFirmness](b.BerryFirmnessUrl + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

// Return a list of BerryFirmness resource
func (b Berries) GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFirmnessList, err := request.GetResourceList[models.NamedResourceList](b.BerryFirmnessUrl, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFirmnessList, nil

}

// Return the the BerryFirmness resource URL
func (b Berries) GetBerryFirmnessUrl() string {
	return b.BerryFirmnessUrl
}

// Return a single BerryFlavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error) {
	berryFlavor, err := request.GetResource[models.BerryFlavor](b.BerryFlavorUrl + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

// Return a list of BerryFlavor resource
func (b Berries) GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFlavorList, err := request.GetResourceList[models.NamedResourceList](b.BerryFlavorUrl, limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFlavorList, nil

}

// Return the the BerryFlavor resource URL
func (b Berries) GetBerryFlavorUrl() string {
	return b.BerryFlavorUrl
}
