package machines

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Machines group interface
type IMachines interface {
	GetMachine(id string) (*models.Machine, error)
	GetMachineList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Machines group struct
type Machines struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Machines resource group initialized")
}

// Return an instance of Items resource group struct
func NewMachinesGroup() Machines {
	return Machines{
		Cache: cache.C,
	}
}

// Return a single Machine resource by  ID
func (m Machines) GetMachine(id string) (*models.Machine, error) {
	machine, err := request.GetSpecificResource[models.Machine](constants.MachineEndpoint + id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

// Return a list of Machine resource
func (m Machines) GetMachineList(options models.PaginationOptions) (*models.ResourceList, error) {
	machineList, err := request.GetResourceList(constants.MachineEndpoint, options)
	if err != nil {
		return nil, err
	}
	return machineList, nil
}
