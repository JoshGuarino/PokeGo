package games

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Games group interface
type IGames interface {
	GetGeneration(nameOrId string) (*models.Generation, error)
	GetGenerationList(options models.PaginationOptions) (*models.ResourceList, error)
	GetPokedex(nameOrId string) (*models.Pokedex, error)
	GetPokedexList(options models.PaginationOptions) (*models.ResourceList, error)
	GetVersion(nameOrId string) (*models.Version, error)
	GetVersionList(options models.PaginationOptions) (*models.ResourceList, error)
	GetVersionGroup(nameOrId string) (*models.VersionGroup, error)
	GetVersionGroupList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Games group struct
type Games struct {
	Cache *cache.Cache
}

// Return an instance of Games resource group struct
func NewGamesGroup() Games {
	return Games{
		Cache: cache.C,
	}
}

// Return a single Generation resource by name or ID
func (g Games) GetGeneration(nameOrId string) (*models.Generation, error) {
	generation, err := request.GetSpecificResource[models.Generation](constants.GenerationEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// Return a list of Generation resource
func (g Games) GetGenerationList(options models.PaginationOptions) (*models.ResourceList, error) {
	generationList, err := request.GetResourceList(constants.GenerationEndpoint, options)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}

// Return a single Pokedex resource by name or ID
func (g Games) GetPokedex(nameOrId string) (*models.Pokedex, error) {
	pokedex, err := request.GetSpecificResource[models.Pokedex](constants.PokedexEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokedex, nil
}

// Return a list of Pokedex resource
func (g Games) GetPokedexList(options models.PaginationOptions) (*models.ResourceList, error) {
	pokedexList, err := request.GetResourceList(constants.PokedexEndpoint, options)
	if err != nil {
		return nil, err
	}
	return pokedexList, nil
}

// Return a single Version resource by name or ID
func (g Games) GetVersion(nameOrId string) (*models.Version, error) {
	version, err := request.GetSpecificResource[models.Version](constants.VersionEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return version, nil
}

// Return a list of Version resource
func (g Games) GetVersionList(options models.PaginationOptions) (*models.ResourceList, error) {
	versionList, err := request.GetResourceList(constants.VersionEndpoint, options)
	if err != nil {
		return nil, err
	}
	return versionList, nil
}

// Return a single VersionGroup resource by name or ID
func (g Games) GetVersionGroup(nameOrId string) (*models.VersionGroup, error) {
	versionGroup, err := request.GetSpecificResource[models.VersionGroup](constants.VersionGroupEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return versionGroup, nil
}

// Return a list of VersionGroup resource
func (g Games) GetVersionGroupList(options models.PaginationOptions) (*models.ResourceList, error) {
	versionGroupList, err := request.GetResourceList(constants.VersionGroupEndpoint, options)
	if err != nil {
		return nil, err
	}
	return versionGroupList, nil
}
