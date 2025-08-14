package evolution

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/internal/logger"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Evolution group resource endpoints
const (
	EvolutionChainEndpoint   = "/evolution-chain/"
	EvolutionTriggerEndpoint = "/evolution-trigger/"
)

// Evolution group interface
type IEvolution interface {
	GetEvolutionChain(id string) (*models.EvolutionChain, error)
	GetEvolutionChainList(limit int, offest int) (*models.ResourceList, error)
	GetEvolutionChainURL() string
	GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error)
	GetEvolutionTriggerList(limit int, offset int) (*models.NamedResourceList, error)
	GetEvolutionTriggerURL() string
}

// Evolution group struct
type Evolution struct {
	Cache *cache.Cache
	Env   *environment.Environment
	Log   *logger.Logger
}

// Initialize function
func init() {
	logger.LOG.Info("Evolution resource group initialized")
}

// Return an instance of Evolution resource group struct
func NewEvolutionGroup() Evolution {
	return Evolution{
		Cache: cache.CACHE,
		Env:   environment.ENV,
		Log:   logger.LOG,
	}
}

// Return a single EvolutionChain resource by ID
func (e Evolution) GetEvolutionChain(id string) (*models.EvolutionChain, error) {
	evolutionChain, err := request.GetResource[models.EvolutionChain](e.GetEvolutionChainURL() + id)
	if err != nil {
		return nil, err
	}
	return evolutionChain, nil
}

// Return a list of EvolutionChain resource
func (e Evolution) GetEvolutionChainList(limit int, offset int) (*models.ResourceList, error) {
	evolutionChainList, err := request.GetResourceList[models.ResourceList](e.GetEvolutionChainURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionChainList, nil
}

// Return the EvolutionChain resource URL
func (e Evolution) GetEvolutionChainURL() string {
	return e.Env.URL() + EvolutionChainEndpoint
}

// Return a single EvolutionTrigger resource by name or ID
func (e Evolution) GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error) {
	evolutionTrigger, err := request.GetResource[models.EvolutionTrigger](e.GetEvolutionTriggerURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return evolutionTrigger, nil
}

// Return a list of EvolutionTrigger resource
func (e Evolution) GetEvolutionTriggerList(limit int, offset int) (*models.NamedResourceList, error) {
	evolutionTriggerList, err := request.GetResourceList[models.NamedResourceList](e.GetEvolutionTriggerURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionTriggerList, nil
}

// Return the EvolutionTrigger resource URL
func (e Evolution) GetEvolutionTriggerURL() string {
	return e.Env.URL() + EvolutionTriggerEndpoint
}
