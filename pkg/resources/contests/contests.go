package contests

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Contests group interface
type IContests interface {
	GetContestType(nameOrId string) (*models.ContestType, error)
	GetContestTypeList(options models.PaginationOptions) (*models.ResourceList, error)
	GetContestEffect(id string) (*models.ContestEffect, error)
	GetContestEffectList(options models.PaginationOptions) (*models.ResourceList, error)
	GetSuperContestEffect(id string) (*models.SuperContestEffect, error)
	GetSuperContestEffectList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Contests group struct
type Contests struct{}

// Return an instance of Contests resource group struct
func NewContestsGroup() Contests {
	return Contests{}
}

// Return a single ContestsType resource by name or ID
func (c Contests) GetContestType(nameOrId string) (*models.ContestType, error) {
	contestType, err := request.GetSpecificResource[models.ContestType](constants.ContestTypeEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return contestType, nil
}

// Return a list of Berry resource
func (c Contests) GetContestTypeList(options models.PaginationOptions) (*models.ResourceList, error) {
	contestTypeList, err := request.GetResourceList(constants.ContestTypeEndpoint, options)
	if err != nil {
		return nil, err
	}
	return &contestTypeList, nil
}

// Return a single ContestsEffect resource by ID
func (c Contests) GetContestEffect(id string) (*models.ContestEffect, error) {
	contestEffect, err := request.GetSpecificResource[models.ContestEffect](constants.ContestEffectEndpoint + id)
	if err != nil {
		return nil, err
	}
	return contestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetContestEffectList(options models.PaginationOptions) (*models.ResourceList, error) {
	contestEffectList, err := request.GetResourceList(constants.ContestEffectEndpoint, options)
	if err != nil {
		return nil, err
	}
	return &contestEffectList, nil
}

// Return a single SuperContestsEffect resource by ID
func (c Contests) GetSuperContestEffect(id string) (*models.SuperContestEffect, error) {
	superContestEffect, err := request.GetSpecificResource[models.SuperContestEffect](constants.SuperContestEffectEndpoint + id)
	if err != nil {
		return nil, err
	}
	return superContestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetSuperContestEffectList(options models.PaginationOptions) (*models.ResourceList, error) {
	superContestEffectList, err := request.GetResourceList(constants.SuperContestEffectEndpoint, options)
	if err != nil {
		return nil, err
	}
	return &superContestEffectList, nil
}
