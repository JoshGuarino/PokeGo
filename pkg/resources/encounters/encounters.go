package encounters

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Encounters group interface
type IEncounters interface {
	GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error)
	GetEncounterMethodList(options models.PaginationOptions) (*models.ResourceList, error)
	GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error)
	GetEncounterConditionList(options models.PaginationOptions) (*models.ResourceList, error)
	GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error)
	GetEncounterConditionValueList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Encounters group struct
type Encounters struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Encounters resource group initialized")
}

// Return an instance of Encounters resource group struct
func NewEncountersGroup() Encounters {
	return Encounters{
		Cache: cache.C,
	}
}

// Return a single EncounterMethod resource by name or ID
func (e Encounters) GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error) {
	encounterMethod, err := request.GetSpecificResource[models.EncounterMethod](constants.EncounterMethodEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterMethod, nil
}

// Return a list of EncounterMethod resource
func (e Encounters) GetEncounterMethodList(options models.PaginationOptions) (*models.ResourceList, error) {
	encounterMethodList, err := request.GetResourceList(constants.EncounterMethodEndpoint, options)
	if err != nil {
		return nil, err
	}
	return encounterMethodList, nil
}

// Return a single EncounterCondition resource by name or ID
func (e Encounters) GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error) {
	encounterCondition, err := request.GetSpecificResource[models.EncounterCondition](constants.EncounterConditionEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterCondition, nil
}

// Return a list of EncounterCondition resource
func (e Encounters) GetEncounterConditionList(options models.PaginationOptions) (*models.ResourceList, error) {
	encounterConditionList, err := request.GetResourceList(constants.EncounterConditionEndpoint, options)
	if err != nil {
		return nil, err
	}
	return encounterConditionList, nil
}

// Return a single EncounterCondition resource by name or ID
func (e Encounters) GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error) {
	encounterConditionValue, err := request.GetSpecificResource[models.EncounterConditionValue](constants.EncounterConditionValueEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterConditionValue, nil
}

// Return a list of EncounterCondition resource
func (e Encounters) GetEncounterConditionValueList(options models.PaginationOptions) (*models.ResourceList, error) {
	encounterConditionValueList, err := request.GetResourceList(constants.EncounterConditionValueEndpoint, options)
	if err != nil {
		return nil, err
	}
	return encounterConditionValueList, nil
}
