package encounters

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Encounters group interface
type IEncounters interface {
	GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error)
	GetEncounterMethodList() (*models.ResourceList, error)
	GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error)
	GetEncounterConditionList() (*models.ResourceList, error)
	GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error)
	GetEncounterConditionValueList() (*models.ResourceList, error)
}

// Encounters group struct
type Encounters struct{}

// Return an instance of Encounters resource group struct
func NewEncountersGroup() Encounters {
	return Encounters{}
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
func (e Encounters) GetEncounterMethodList() (*models.ResourceList, error) {
	encounterMethodList, err := request.GetResourceList(constants.EncounterMethodEndpoint)
	if err != nil {
		return nil, err
	}
	return &encounterMethodList, nil
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
func (e Encounters) GetEncounterConditionList() (*models.ResourceList, error) {
	encounterConditionList, err := request.GetResourceList(constants.EncounterConditionEndpoint)
	if err != nil {
		return nil, err
	}
	return &encounterConditionList, nil
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
func (e Encounters) GetEncounterConditionValueList() (*models.ResourceList, error) {
	encounterConditionValueList, err := request.GetResourceList(constants.EncounterConditionValueEndpoint)
	if err != nil {
		return nil, err
	}
	return &encounterConditionValueList, nil
}
