package evolution

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Evolution group interface
type IEvolution interface {
	GetEvolutionChain(id string) (*models.EvolutionChain, error)
	GetEvolutionChainList(options models.PaginationOptions) (*models.ResourceList, error)
	GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error)
	GetEvolutionTriggerList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Evolution group struct
type Evolution struct{}

// Return an instance of Evolution resource group struct
func NewEvolutionGroup() Evolution {
	return Evolution{}
}

// Return a single EvolutionChain resource by ID
func (e Evolution) GetEvolutionChain(id string) (*models.EvolutionChain, error) {
	evolutionChain, err := request.GetSpecificResource[models.EvolutionChain](constants.EvolutionChainEndpoint + id)
	if err != nil {
		return nil, err
	}
	return evolutionChain, nil
}

// Return a list of EvolutionChain resource
func (e Evolution) GetEvolutionChainList(options models.PaginationOptions) (*models.ResourceList, error) {
	evolutionChainList, err := request.GetResourceList(constants.EvolutionChainEndpoint, options)
	if err != nil {
		return nil, err
	}
	return evolutionChainList, nil
}

// Return a single EvolutionTrigger resource by name or ID
func (e Evolution) GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error) {
	evolutionTrigger, err := request.GetSpecificResource[models.EvolutionTrigger](constants.EvolutionTriggerEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return evolutionTrigger, nil
}

// Return a list of EvolutionTrigger resource
func (e Evolution) GetEvolutionTriggerList(options models.PaginationOptions) (*models.ResourceList, error) {
	evolutionTriggerList, err := request.GetResourceList(constants.EvolutionTriggerEndpoint, options)
	if err != nil {
		return nil, err
	}
	return evolutionTriggerList, nil
}
