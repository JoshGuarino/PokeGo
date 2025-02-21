package contests

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
)

// Contests group interface
type IContests interface {
	GetContestType(nameOrId string) (*ContestType, error)
	GetContestTypeList() (*request.ResourceList, error)
	GetContestEffect(id string) (*ContestEffect, error)
	GetContestEffectList() (*request.ResourceList, error)
	GetSuperContestEffect(id string) (*SuperContestEffect, error)
	GetSuperContestEffectList() (*request.ResourceList, error)
}

// Contests group struct
type Contests struct{}

// Return an instance of Berry resource group struct
func NewContestsGroup() Contests {
	return Contests{}
}

// Return a single ContestsType resource by name or ID
func (c Contests) GetContestType(nameOrId string) (*ContestType, error) {
	contestType, err := request.GetSpecificResource[ContestType](constants.ContestTypeEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return contestType, nil
}

// Return a list of Berry resource
func (c Contests) GetContestTypeList() (*request.ResourceList, error) {
	contestTypeList, err := request.GetResourceList(constants.ContestTypeEndpoint)
	if err != nil {
		return nil, err
	}
	return &contestTypeList, nil
}

// Return a single ContestsEffect resource by ID
func (c Contests) GetContestEffect(id string) (*ContestEffect, error) {
	contestEffect, err := request.GetSpecificResource[ContestEffect](constants.ContestEffectEndpoint + id)
	if err != nil {
		return nil, err
	}
	return contestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetContestEffectList() (*request.ResourceList, error) {
	contestEffectList, err := request.GetResourceList(constants.ContestEffectEndpoint)
	if err != nil {
		return nil, err
	}
	return &contestEffectList, nil
}

// Return a single SuperContestsEffect resource by ID
func (c Contests) GetSuperContestEffect(id string) (*SuperContestEffect, error) {
	superContestEffect, err := request.GetSpecificResource[SuperContestEffect](constants.SuperContestEffectEndpoint + id)
	if err != nil {
		return nil, err
	}
	return superContestEffect, nil
}

// Return a list of Berry resource
func (c Contests) GetSuperContestEffectList() (*request.ResourceList, error) {
	superContestEffectList, err := request.GetResourceList(constants.SuperContestEffectEndpoint)
	if err != nil {
		return nil, err
	}
	return &superContestEffectList, nil
}
