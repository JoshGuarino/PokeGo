package evolution

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Evolution group interface
type IEvolution interface {
	GetEvolutionChain(id string) (*models.EvolutionChain, error)
	GetEvolutionChainList(limit int, offest int) (*models.ResourceList, error)
	GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error)
	GetEvolutionTriggerList(limit int, offset int) (*models.ResourceList, error)
}

// Evolution group struct
type Evolution struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Evolution resource group initialized")
}

// Return an instance of Evolution resource group struct
func NewEvolutionGroup() Evolution {
	return Evolution{
		Cache: cache.C,
	}
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
func (e Evolution) GetEvolutionChainList(limit int, offset int) (*models.ResourceList, error) {
	evolutionChainList, err := request.GetResourceList(constants.EvolutionChainEndpoint, limit, offset)
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
func (e Evolution) GetEvolutionTriggerList(limit int, offset int) (*models.ResourceList, error) {
	evolutionTriggerList, err := request.GetResourceList(constants.EvolutionTriggerEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionTriggerList, nil
}
