package evolution

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
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
	EvolutionChainURL   string
	EvolutionTriggerURL string
	Cache               *cache.Cache
}

// Initialize function
func init() {
	log.Info("Evolution resource group initialized")
}

// Return an instance of Evolution resource group struct
func NewEvolutionGroup() Evolution {
	url := endpoints.BaseURL
	return Evolution{
		EvolutionChainURL:   url + endpoints.EvolutionChain,
		EvolutionTriggerURL: url + endpoints.EvolutionTrigger,
		Cache:               cache.C,
	}
}

// Return a single EvolutionChain resource by ID
func (e Evolution) GetEvolutionChain(id string) (*models.EvolutionChain, error) {
	evolutionChain, err := request.GetResource[models.EvolutionChain](e.EvolutionChainURL + id)
	if err != nil {
		return nil, err
	}
	return evolutionChain, nil
}

// Return a list of EvolutionChain resource
func (e Evolution) GetEvolutionChainList(limit int, offset int) (*models.ResourceList, error) {
	evolutionChainList, err := request.GetResourceList[models.ResourceList](e.EvolutionChainURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionChainList, nil
}

// Return the EvolutionChain resource URL
func (e Evolution) GetEvolutionChainURL() string {
	return e.EvolutionChainURL
}

// Return a single EvolutionTrigger resource by name or ID
func (e Evolution) GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error) {
	evolutionTrigger, err := request.GetResource[models.EvolutionTrigger](e.EvolutionTriggerURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return evolutionTrigger, nil
}

// Return a list of EvolutionTrigger resource
func (e Evolution) GetEvolutionTriggerList(limit int, offset int) (*models.NamedResourceList, error) {
	evolutionTriggerList, err := request.GetResourceList[models.NamedResourceList](e.EvolutionTriggerURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionTriggerList, nil
}

// Return the EvolutionTrigger resource URL
func (e Evolution) GetEvolutionTriggerURL() string {
	return e.EvolutionTriggerURL
}
