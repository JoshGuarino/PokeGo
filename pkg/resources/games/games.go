package games

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Games group interface
type IGames interface {
	GetGeneration(nameOrId string) (*models.Generation, error)
	GetGenerationList(limit int, offest int) (*models.ResourceList, error)
	GetPokedex(nameOrId string) (*models.Pokedex, error)
	GetPokedexList(limit int, offest int) (*models.ResourceList, error)
	GetVersion(nameOrId string) (*models.Version, error)
	GetVersionList(limit int, offest int) (*models.ResourceList, error)
	GetVersionGroup(nameOrId string) (*models.VersionGroup, error)
	GetVersionGroupList(limit int, offest int) (*models.ResourceList, error)
}

// Games group struct
type Games struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Games resource group initialized")
}

// Return an instance of Games resource group struct
func NewGamesGroup() Games {
	return Games{
		Cache: cache.C,
	}
}

// Return a single Generation resource by name or ID
func (g Games) GetGeneration(nameOrId string) (*models.Generation, error) {
	generation, err := request.GetSpecificResource[models.Generation](endpoints.Generation + nameOrId)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// Return a list of Generation resource
func (g Games) GetGenerationList(limit int, offest int) (*models.ResourceList, error) {
	generationList, err := request.GetResourceList(endpoints.Generation, limit, offest)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}

// Return a single Pokedex resource by name or ID
func (g Games) GetPokedex(nameOrId string) (*models.Pokedex, error) {
	pokedex, err := request.GetSpecificResource[models.Pokedex](endpoints.Pokedex + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokedex, nil
}

// Return a list of Pokedex resource
func (g Games) GetPokedexList(limit int, offest int) (*models.ResourceList, error) {
	pokedexList, err := request.GetResourceList(endpoints.Pokedex, limit, offest)
	if err != nil {
		return nil, err
	}
	return pokedexList, nil
}

// Return a single Version resource by name or ID
func (g Games) GetVersion(nameOrId string) (*models.Version, error) {
	version, err := request.GetSpecificResource[models.Version](endpoints.Version + nameOrId)
	if err != nil {
		return nil, err
	}
	return version, nil
}

// Return a list of Version resource
func (g Games) GetVersionList(limit int, offest int) (*models.ResourceList, error) {
	versionList, err := request.GetResourceList(endpoints.Version, limit, offest)
	if err != nil {
		return nil, err
	}
	return versionList, nil
}

// Return a single VersionGroup resource by name or ID
func (g Games) GetVersionGroup(nameOrId string) (*models.VersionGroup, error) {
	versionGroup, err := request.GetSpecificResource[models.VersionGroup](endpoints.VersionGroup + nameOrId)
	if err != nil {
		return nil, err
	}
	return versionGroup, nil
}

// Return a list of VersionGroup resource
func (g Games) GetVersionGroupList(limit int, offest int) (*models.ResourceList, error) {
	versionGroupList, err := request.GetResourceList(endpoints.VersionGroup, limit, offest)
	if err != nil {
		return nil, err
	}
	return versionGroupList, nil
}
