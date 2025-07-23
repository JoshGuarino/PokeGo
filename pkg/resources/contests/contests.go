package contests

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Contests group interface
type IContests interface {
	GetContestType(nameOrId string) (*models.ContestType, error)
	GetContestTypeList(limit int, offset int) (*models.NamedResourceList, error)
	GetContestTypeURL() string
	GetContestEffect(id string) (*models.ContestEffect, error)
	GetContestEffectList(limit int, offset int) (*models.ResourceList, error)
	GetContestEffectURL() string
	GetSuperContestEffect(id string) (*models.SuperContestEffect, error)
	GetSuperContestEffectList(limit int, offset int) (*models.ResourceList, error)
	GetSuperContestEffectURL() string
}

// Contests group struct
type Contests struct {
	ContestTypeURL        string
	ContestEffectURL      string
	SuperContestEffectURL string
	Cache                 *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Contests resource group initialized")
}

// Return an instance of Contests resource group struct
func NewContestsGroup() Contests {
	url := endpoints.BaseURL
	return Contests{
		ContestTypeURL:        url + endpoints.ContestType,
		ContestEffectURL:      url + endpoints.ContestEffect,
		SuperContestEffectURL: url + endpoints.SuperContestEffect,
		Cache:                 cache.C,
	}
}

// Return a single ContestType resource by name or ID
func (c Contests) GetContestType(nameOrId string) (*models.ContestType, error) {
	contestType, err := request.GetResource[models.ContestType](c.ContestTypeURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return contestType, nil
}

// Return a list of ContestType resource
func (c Contests) GetContestTypeList(limit int, offset int) (*models.NamedResourceList, error) {
	contestTypeList, err := request.GetResourceList[models.NamedResourceList](c.ContestTypeURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return contestTypeList, nil
}

// Return the the ContestType resource URL
func (c Contests) GetContestTypeURL() string {
	return c.ContestTypeURL
}

// Return a single ContestEffect resource by ID
func (c Contests) GetContestEffect(id string) (*models.ContestEffect, error) {
	contestEffect, err := request.GetResource[models.ContestEffect](c.ContestEffectURL + id)
	if err != nil {
		return nil, err
	}
	return contestEffect, nil
}

// Return a list of ContestEffect resource
func (c Contests) GetContestEffectList(limit int, offest int) (*models.ResourceList, error) {
	contestEffectList, err := request.GetResourceList[models.ResourceList](c.ContestEffectURL, limit, offest)
	if err != nil {
		return nil, err
	}
	return contestEffectList, nil
}

// Return the the ContestEffect resource URL
func (c Contests) GetContestEffectURL() string {
	return c.ContestEffectURL
}

// Return a single SuperContestEffect resource by ID
func (c Contests) GetSuperContestEffect(id string) (*models.SuperContestEffect, error) {
	superContestEffect, err := request.GetResource[models.SuperContestEffect](c.SuperContestEffectURL + id)
	if err != nil {
		return nil, err
	}
	return superContestEffect, nil
}

// Return a list of SuperContestEffect resource
func (c Contests) GetSuperContestEffectList(limit int, offest int) (*models.ResourceList, error) {
	superContestEffectList, err := request.GetResourceList[models.ResourceList](c.SuperContestEffectURL, limit, offest)
	if err != nil {
		return nil, err
	}
	return superContestEffectList, nil
}

// Return the the SuperContestEffect resource URL
func (c Contests) GetSuperContestEffectURL() string {
	return c.SuperContestEffectURL
}
