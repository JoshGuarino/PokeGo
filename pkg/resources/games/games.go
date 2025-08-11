package games

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Games group resource endpoints
const (
	GenerationEndpoint   = "/generation/"
	PokedexEndpoint      = "/pokedex/"
	VersionEndpoint      = "/version/"
	VersionGroupEndpoint = "/version-group/"
)

// Games group interface
type IGames interface {
	GetGeneration(nameOrId string) (*models.Generation, error)
	GetGenerationList(limit int, offest int) (*models.NamedResourceList, error)
	GetGenerationURL() string
	GetPokedex(nameOrId string) (*models.Pokedex, error)
	GetPokedexList(limit int, offest int) (*models.NamedResourceList, error)
	GetPokedexURL() string
	GetVersion(nameOrId string) (*models.Version, error)
	GetVersionList(limit int, offest int) (*models.NamedResourceList, error)
	GetVersionURL() string
	GetVersionGroup(nameOrId string) (*models.VersionGroup, error)
	GetVersionGroupList(limit int, offest int) (*models.NamedResourceList, error)
	GetVersionGroupURL() string
}

// Games group struct
type Games struct {
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Games resource group initialized")
}

// Return an instance of Games resource group struct
func NewGamesGroup() Games {
	return Games{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single Generation resource by name or ID
func (g Games) GetGeneration(nameOrId string) (*models.Generation, error) {
	generation, err := request.GetResource[models.Generation](g.GetGenerationURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// Return a list of Generation resource
func (g Games) GetGenerationList(limit int, offest int) (*models.NamedResourceList, error) {
	generationList, err := request.GetResourceList[models.NamedResourceList](g.GetGenerationURL(), limit, offest)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}

// Return the Generation resource URL
func (g Games) GetGenerationURL() string {
	return g.Env.URL() + GenerationEndpoint
}

// Return a single Pokedex resource by name or ID
func (g Games) GetPokedex(nameOrId string) (*models.Pokedex, error) {
	pokedex, err := request.GetResource[models.Pokedex](g.GetPokedexURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokedex, nil
}

// Return a list of Pokedex resource
func (g Games) GetPokedexList(limit int, offest int) (*models.NamedResourceList, error) {
	pokedexList, err := request.GetResourceList[models.NamedResourceList](g.GetPokedexURL(), limit, offest)
	if err != nil {
		return nil, err
	}
	return pokedexList, nil
}

// Return the Pokedex resource URL
func (g Games) GetPokedexURL() string {
	return g.Env.URL() + PokedexEndpoint
}

// Return a single Version resource by name or ID
func (g Games) GetVersion(nameOrId string) (*models.Version, error) {
	version, err := request.GetResource[models.Version](g.GetVersionURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return version, nil
}

// Return a list of Version resource
func (g Games) GetVersionList(limit int, offest int) (*models.NamedResourceList, error) {
	versionList, err := request.GetResourceList[models.NamedResourceList](g.GetVersionURL(), limit, offest)
	if err != nil {
		return nil, err
	}
	return versionList, nil
}

// Return the Version resource URL
func (g Games) GetVersionURL() string {
	return g.Env.URL() + VersionEndpoint
}

// Return a single VersionGroup resource by name or ID
func (g Games) GetVersionGroup(nameOrId string) (*models.VersionGroup, error) {
	versionGroup, err := request.GetResource[models.VersionGroup](g.GetVersionGroupURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return versionGroup, nil
}

// Return a list of VersionGroup resource
func (g Games) GetVersionGroupList(limit int, offest int) (*models.NamedResourceList, error) {
	versionGroupList, err := request.GetResourceList[models.NamedResourceList](g.GetVersionGroupURL(), limit, offest)
	if err != nil {
		return nil, err
	}
	return versionGroupList, nil
}

// Return the VersionGroup resource URL
func (g Games) GetVersionGroupURL() string {
	return g.Env.URL() + VersionGroupEndpoint
}
