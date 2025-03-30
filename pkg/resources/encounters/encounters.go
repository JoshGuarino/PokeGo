package encounters

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Encounters group interface
type IEncounters interface {
	GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error)
	GetEncounterMethodList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error)
	GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error)
	GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error)
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
	encounterMethod, err := request.GetResource[models.EncounterMethod](endpoints.EncounterMethod + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterMethod, nil
}

// Return a list of EncounterMethod resource
func (e Encounters) GetEncounterMethodList(limit int, offest int) (*models.NamedResourceList, error) {
	encounterMethodList, err := request.GetResourceList[models.NamedResourceList](endpoints.EncounterMethod, limit, offest)
	if err != nil {
		return nil, err
	}
	return encounterMethodList, nil
}

// Return a single EncounterCondition resource by name or ID
func (e Encounters) GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error) {
	encounterCondition, err := request.GetResource[models.EncounterCondition](endpoints.EncounterCondition + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterCondition, nil
}

// Return a list of EncounterCondition resource
func (e Encounters) GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionList, err := request.GetResourceList[models.NamedResourceList](endpoints.EncounterCondition, limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionList, nil
}

// Return a single EncounterConditionValue resource by name or ID
func (e Encounters) GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error) {
	encounterConditionValue, err := request.GetResource[models.EncounterConditionValue](endpoints.EncounterConditionValue + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterConditionValue, nil
}

// Return a list of EncounterConditionValue resource
func (e Encounters) GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionValueList, err := request.GetResourceList[models.NamedResourceList](endpoints.EncounterConditionValue, limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionValueList, nil
}
