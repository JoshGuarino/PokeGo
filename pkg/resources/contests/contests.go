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
	GetContestTypeList(limit int, offset int) (*models.ResourceList, error)
	GetContestEffect(id string) (*models.ContestEffect, error)
	GetContestEffectList(limit int, offset int) (*models.ResourceList, error)
	GetSuperContestEffect(id string) (*models.SuperContestEffect, error)
	GetSuperContestEffectList(limit int, offset int) (*models.ResourceList, error)
}

// Contests group struct
type Contests struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Contests resource group initialized")
}

// Return an instance of Contests resource group struct
func NewContestsGroup() Contests {
	return Contests{
		Cache: cache.C,
	}
}

// Return a single ContestsType resource by name or ID
func (c Contests) GetContestType(nameOrId string) (*models.ContestType, error) {
	contestType, err := request.GetSpecificResource[models.ContestType](endpoints.ContestType + nameOrId)
	if err != nil {
		return nil, err
	}
	return contestType, nil
}

// Return a list of Berry resource
func (c Contests) GetContestTypeList(limit int, offset int) (*models.ResourceList, error) {
	contestTypeList, err := request.GetResourceList(endpoints.ContestType, limit, offset)
	if err != nil {
		return nil, err
	}
	return contestTypeList, nil
}

// Return a single ContestsEffect resource by ID
func (c Contests) GetContestEffect(id string) (*models.ContestEffect, error) {
	contestEffect, err := request.GetSpecificResource[models.ContestEffect](endpoints.ContestEffect + id)
	if err != nil {
		return nil, err
	}
	return contestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetContestEffectList(limit int, offest int) (*models.ResourceList, error) {
	contestEffectList, err := request.GetResourceList(endpoints.ContestEffect, limit, offest)
	if err != nil {
		return nil, err
	}
	return contestEffectList, nil
}

// Return a single SuperContestsEffect resource by ID
func (c Contests) GetSuperContestEffect(id string) (*models.SuperContestEffect, error) {
	superContestEffect, err := request.GetSpecificResource[models.SuperContestEffect](endpoints.SuperContestEffect + id)
	if err != nil {
		return nil, err
	}
	return superContestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetSuperContestEffectList(limit int, offest int) (*models.ResourceList, error) {
	superContestEffectList, err := request.GetResourceList(endpoints.SuperContestEffect, limit, offest)
	if err != nil {
		return nil, err
	}
	return superContestEffectList, nil
}
