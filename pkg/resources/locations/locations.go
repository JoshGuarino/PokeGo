package locations

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Locations group interface
type ILocations interface {
	GetLocation(nameOrId string) (*models.Location, error)
	GetLocationList(limit int, offset int) (*models.NamedResourceList, error)
	GetLocationArea(nameOrId string) (*models.LocationArea, error)
	GetLocationAreaList(limit int, offset int) (*models.NamedResourceList, error)
	GetPalParkArea(nameOrId string) (*models.PalParkArea, error)
	GetPalParkAreaList(limit int, offset int) (*models.NamedResourceList, error)
	GetRegion(nameOrId string) (*models.Region, error)
	GetRegionList(limit int, offset int) (*models.NamedResourceList, error)
}

// Locations group struct
type Locations struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Locations resource group initialized")
}

// Return an instance of Locations resource group struct
func NewLocationsGroup() Locations {
	return Locations{
		Cache: cache.C,
	}
}

// Return a single Location resource by name or ID
func (l Locations) GetLocation(nameOrId string) (*models.Location, error) {
	location, err := request.GetResource[models.Location](endpoints.Location + nameOrId)
	if err != nil {
		return nil, err
	}
	return location, nil
}

// Return a list of Location resource
func (l Locations) GetLocationList(limit int, offset int) (*models.NamedResourceList, error) {
	locationsList, err := request.GetResourceList[models.NamedResourceList](endpoints.Location, limit, offset)
	if err != nil {
		return nil, err
	}
	return locationsList, nil
}

// Return a single LocationArea resource by name or ID
func (l Locations) GetLocationArea(nameOrId string) (*models.LocationArea, error) {
	locationArea, err := request.GetResource[models.LocationArea](endpoints.LocationArea + nameOrId)
	if err != nil {
		return nil, err
	}
	return locationArea, nil
}

// Return a list of LocationArea resource
func (l Locations) GetLocationAreaList(limit int, offset int) (*models.NamedResourceList, error) {
	locationsAreaList, err := request.GetResourceList[models.NamedResourceList](endpoints.LocationArea, limit, offset)
	if err != nil {
		return nil, err
	}
	return locationsAreaList, nil
}

// Return a single PalParkArea resource by name or ID
func (l Locations) GetPalParkArea(nameOrId string) (*models.PalParkArea, error) {
	PalParkArea, err := request.GetResource[models.PalParkArea](endpoints.PalParkArea + nameOrId)
	if err != nil {
		return nil, err
	}
	return PalParkArea, nil
}

// Return a list of PalParkArea resource
func (l Locations) GetPalParkAreaList(limit int, offset int) (*models.NamedResourceList, error) {
	palParkAreaList, err := request.GetResourceList[models.NamedResourceList](endpoints.PalParkArea, limit, offset)
	if err != nil {
		return nil, err
	}
	return palParkAreaList, nil
}

// Return a single Region resource by name or ID
func (l Locations) GetRegion(nameOrId string) (*models.Region, error) {
	Region, err := request.GetResource[models.Region](endpoints.Region + nameOrId)
	if err != nil {
		return nil, err
	}
	return Region, nil
}

// Return a list of PalParkArea resource
func (l Locations) GetRegionList(limit int, offset int) (*models.NamedResourceList, error) {
	regionList, err := request.GetResourceList[models.NamedResourceList](endpoints.Region, limit, offset)
	if err != nil {
		return nil, err
	}
	return regionList, nil
}
