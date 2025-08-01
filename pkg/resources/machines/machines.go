package machines

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Machines group interface
type IMachines interface {
	GetMachine(id string) (*models.Machine, error)
	GetMachineList(limit int, offset int) (*models.ResourceList, error)
	GetMachineURL() string
}

// Machines group struct
type Machines struct {
	MachineURL string
	Cache      *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Machines resource group initialized")
}

// Return an instance of Items resource group struct
func NewMachinesGroup() Machines {
	url := endpoints.BaseURL
	return Machines{
		MachineURL: url + endpoints.Machine,
		Cache:      cache.C,
	}
}

// Return a single Machine resource by  ID
func (m Machines) GetMachine(id string) (*models.Machine, error) {
	machine, err := request.GetResource[models.Machine](m.MachineURL + id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

// Return a list of Machine resource
func (m Machines) GetMachineList(limit int, offset int) (*models.ResourceList, error) {
	machineList, err := request.GetResourceList[models.ResourceList](m.MachineURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return machineList, nil
}

// Return the Machine resource url
func (m Machines) GetMachineURL() string {
	return m.MachineURL
}
