package locations

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Locations group interface
type ILocations interface {
	GetLocation(nameOrId string) (*models.Location, error)
	GetLocationList(options models.PaginationOptions) (*models.ResourceList, error)
	GetLocationArea(nameOrId string) (*models.LocationArea, error)
	GetLocationAreaList(options models.PaginationOptions) (*models.ResourceList, error)
	GetPalParkArea(nameOrId string) (*models.PalParkArea, error)
	GetPalParkAreaList(options models.PaginationOptions) (*models.ResourceList, error)
	GetRegion(nameOrId string) (*models.Region, error)
	GetRegionList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Locations group struct
type Locations struct{}

// Return an instance of Locations resource group struct
func NewLocationsGroup() Locations {
	return Locations{}
}

// Return a single Location resource by name or ID
func (l Locations) GetLocation(nameOrId string) (*models.Location, error) {
	location, err := request.GetSpecificResource[models.Location](constants.LocationEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return location, nil
}

// Return a list of Location resource
func (l Locations) GetLocationList(options models.PaginationOptions) (*models.ResourceList, error) {
	locationsList, err := request.GetResourceList(constants.LocationEndpoint, options)
	if err != nil {
		return nil, err
	}
	return locationsList, nil
}

// Return a single LocationArea resource by name or ID
func (l Locations) GetLocationArea(nameOrId string) (*models.LocationArea, error) {
	locationArea, err := request.GetSpecificResource[models.LocationArea](constants.LocationAreaEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return locationArea, nil
}

// Return a list of LocationArea resource
func (l Locations) GetLocationAreaList(options models.PaginationOptions) (*models.ResourceList, error) {
	locationsAreaList, err := request.GetResourceList(constants.LocationAreaEndpoint, options)
	if err != nil {
		return nil, err
	}
	return locationsAreaList, nil
}

// Return a single PalParkArea resource by name or ID
func (l Locations) GetPalParkArea(nameOrId string) (*models.PalParkArea, error) {
	PalParkArea, err := request.GetSpecificResource[models.PalParkArea](constants.PalParkAreaEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return PalParkArea, nil
}

// Return a list of PalParkArea resource
func (l Locations) GetPalParkAreaList(options models.PaginationOptions) (*models.ResourceList, error) {
	palParkAreaList, err := request.GetResourceList(constants.PalParkAreaEndpoint, options)
	if err != nil {
		return nil, err
	}
	return palParkAreaList, nil
}

// Return a single Region resource by name or ID
func (l Locations) GetRegion(nameOrId string) (*models.Region, error) {
	Region, err := request.GetSpecificResource[models.Region](constants.RegionEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return Region, nil
}

// Return a list of PalParkArea resource
func (l Locations) GetRegionList(options models.PaginationOptions) (*models.ResourceList, error) {
	regionList, err := request.GetResourceList(constants.RegionEndpoint, options)
	if err != nil {
		return nil, err
	}
	return regionList, nil
}
