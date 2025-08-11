package berries

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Berries group endpoints
const (
	BerryEndpoint         = "/berry/"
	BerryFirmnessEndpoint = "/berry-firmness/"
	BerryFlavorEndpoint   = "/berry-flavor/"
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
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Berries resource group initialized")
}

// Return an instance of Berry resource group struct
func NewBerriesGroup() Berries {
	return Berries{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single Berry resource by name or ID
func (b Berries) GetBerry(nameOrId string) (*models.Berry, error) {
	berry, err := request.GetResource[models.Berry](b.GetBerryURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

// Return a list of Berry resource
func (b Berries) GetBerryList(limit int, offset int) (*models.NamedResourceList, error) {
	berryList, err := request.GetResourceList[models.NamedResourceList](b.GetBerryURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return berryList, nil
}

// Return the Berry resource URL
func (b Berries) GetBerryURL() string {
	return b.Env.URL() + BerryEndpoint
}

// Return a single BerryFirmness resource by name or ID
func (b Berries) GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error) {
	berryFirmness, err := request.GetResource[models.BerryFirmness](b.GetBerryFirmnessURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

// Return a list of BerryFirmness resource
func (b Berries) GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFirmnessList, err := request.GetResourceList[models.NamedResourceList](b.GetBerryFirmnessURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFirmnessList, nil

}

// Return the BerryFirmness resource URL
func (b Berries) GetBerryFirmnessURL() string {
	return b.Env.URL() + BerryFirmnessEndpoint
}

// Return a single BerryFlavor resource by name or ID
func (b Berries) GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error) {
	berryFlavor, err := request.GetResource[models.BerryFlavor](b.GetBerryFlavorURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

// Return a list of BerryFlavor resource
func (b Berries) GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFlavorList, err := request.GetResourceList[models.NamedResourceList](b.GetBerryFlavorURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFlavorList, nil

}

// Return the BerryFlavor resource URL
func (b Berries) GetBerryFlavorURL() string {
	return b.Env.URL() + BerryFlavorEndpoint
}
