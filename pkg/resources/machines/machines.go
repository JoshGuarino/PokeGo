package machines

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Machines group resource endpoints
const (
	MachineEndpoint = "/machine/"
)

// Machines group interface
type IMachines interface {
	GetMachine(id string) (*models.Machine, error)
	GetMachineList(limit int, offset int) (*models.ResourceList, error)
	GetMachineURL() string
}

// Machines group struct
type Machines struct {
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Machines resource group initialized")
}

// Return an instance of Items resource group struct
func NewMachinesGroup() Machines {
	return Machines{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single Machine resource by  ID
func (m Machines) GetMachine(id string) (*models.Machine, error) {
	machine, err := request.GetResource[models.Machine](m.GetMachineURL() + id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

// Return a list of Machine resource
func (m Machines) GetMachineList(limit int, offset int) (*models.ResourceList, error) {
	machineList, err := request.GetResourceList[models.ResourceList](m.GetMachineURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return machineList, nil
}

// Return the Machine resource url
func (m Machines) GetMachineURL() string {
	return m.Env.URL() + MachineEndpoint
}
