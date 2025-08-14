package locations

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/internal/logger"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Locations group resource endpoints
const (
	LocationEndpoint     = "/location/"
	LocationAreaEndpoint = "/location-area/"
	PalParkAreaEndpoint  = "/pal-park-area/"
	RegionEndpoint       = "/region/"
)

// Locations group interface
type ILocations interface {
	GetLocation(nameOrId string) (*models.Location, error)
	GetLocationList(limit int, offset int) (*models.NamedResourceList, error)
	GetLocationURL() string
	GetLocationArea(nameOrId string) (*models.LocationArea, error)
	GetLocationAreaList(limit int, offset int) (*models.NamedResourceList, error)
	GetLocationAreaURL() string
	GetPalParkArea(nameOrId string) (*models.PalParkArea, error)
	GetPalParkAreaList(limit int, offset int) (*models.NamedResourceList, error)
	GetPalParkAreaURL() string
	GetRegion(nameOrId string) (*models.Region, error)
	GetRegionList(limit int, offset int) (*models.NamedResourceList, error)
	GetRegionURL() string
}

// Locations group struct
type Locations struct {
	Cache *cache.Cache
	Env   *environment.Environment
	Log   *logger.Logger
}

// Initialize function
func init() {
	logger.LOG.Info("Locations resource group initialized")
}

// Return an instance of Locations resource group struct
func NewLocationsGroup() Locations {
	return Locations{
		Cache: cache.CACHE,
		Env:   environment.ENV,
		Log:   logger.LOG,
	}
}

// Return a single Location resource by name or ID
func (l Locations) GetLocation(nameOrId string) (*models.Location, error) {
	location, err := request.GetResource[models.Location](l.GetLocationURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return location, nil
}

// Return a list of Location resource
func (l Locations) GetLocationList(limit int, offset int) (*models.NamedResourceList, error) {
	locationsList, err := request.GetResourceList[models.NamedResourceList](l.GetLocationURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return locationsList, nil
}

// Return the Location resource URL
func (l Locations) GetLocationURL() string {
	return l.Env.URL() + LocationEndpoint
}

// Return a single LocationArea resource by name or ID
func (l Locations) GetLocationArea(nameOrId string) (*models.LocationArea, error) {
	locationArea, err := request.GetResource[models.LocationArea](l.GetLocationAreaURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return locationArea, nil
}

// Return a list of LocationArea resource
func (l Locations) GetLocationAreaList(limit int, offset int) (*models.NamedResourceList, error) {
	locationsAreaList, err := request.GetResourceList[models.NamedResourceList](l.GetLocationAreaURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return locationsAreaList, nil
}

// Return the LocationArea resource URL
func (l Locations) GetLocationAreaURL() string {
	return l.Env.URL() + LocationAreaEndpoint
}

// Return a single PalParkArea resource by name or ID
func (l Locations) GetPalParkArea(nameOrId string) (*models.PalParkArea, error) {
	PalParkArea, err := request.GetResource[models.PalParkArea](l.GetPalParkAreaURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return PalParkArea, nil
}

// Return a list of PalParkArea resource
func (l Locations) GetPalParkAreaList(limit int, offset int) (*models.NamedResourceList, error) {
	palParkAreaList, err := request.GetResourceList[models.NamedResourceList](l.GetPalParkAreaURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return palParkAreaList, nil
}

// Return the PalParkArea resource URL
func (l Locations) GetPalParkAreaURL() string {
	return l.Env.URL() + PalParkAreaEndpoint
}

// Return a single Region resource by name or ID
func (l Locations) GetRegion(nameOrId string) (*models.Region, error) {
	Region, err := request.GetResource[models.Region](l.GetRegionURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return Region, nil
}

// Return a list of Region resource
func (l Locations) GetRegionList(limit int, offset int) (*models.NamedResourceList, error) {
	regionList, err := request.GetResourceList[models.NamedResourceList](l.GetRegionURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return regionList, nil
}

// Return the Region resource URL
func (l Locations) GetRegionURL() string {
	return l.Env.URL() + RegionEndpoint
}
