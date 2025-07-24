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
	GetEncounterMethodURL() string
	GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error)
	GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterConditionURL() string
	GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error)
	GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterConditionValueURL() string
}

// Encounters group struct
type Encounters struct {
	EncounterMethodURL         string
	EncounterConditionURL      string
	EncounterConditionValueURL string
	Cache                      *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Encounters resource group initialized")
}

// Return an instance of Encounters resource group struct
func NewEncountersGroup() Encounters {
	url := endpoints.BaseURL
	return Encounters{
		EncounterMethodURL:         url + endpoints.EncounterMethod,
		EncounterConditionURL:      url + endpoints.EncounterCondition,
		EncounterConditionValueURL: url + endpoints.EncounterConditionValue,
		Cache:                      cache.C,
	}
}

// Return a single EncounterMethod resource by name or ID
func (e Encounters) GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error) {
	encounterMethod, err := request.GetResource[models.EncounterMethod](e.EncounterMethodURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterMethod, nil
}

// Return a list of EncounterMethod resource
func (e Encounters) GetEncounterMethodList(limit int, offest int) (*models.NamedResourceList, error) {
	encounterMethodList, err := request.GetResourceList[models.NamedResourceList](e.EncounterMethodURL, limit, offest)
	if err != nil {
		return nil, err
	}
	return encounterMethodList, nil
}

// Return the the EncounterMethod resource URL
func (e Encounters) GetEncounterMethodURL() string {
	return e.EncounterMethodURL
}

// Return a single EncounterCondition resource by name or ID
func (e Encounters) GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error) {
	encounterCondition, err := request.GetResource[models.EncounterCondition](e.EncounterConditionURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterCondition, nil
}

// Return a list of EncounterCondition resource
func (e Encounters) GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionList, err := request.GetResourceList[models.NamedResourceList](e.EncounterConditionURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionList, nil
}

// Return the the EncounterCondition resource URL
func (e Encounters) GetEncounterConditionURL() string {
	return e.EncounterConditionURL
}

// Return a single EncounterConditionValue resource by name or ID
func (e Encounters) GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error) {
	encounterConditionValue, err := request.GetResource[models.EncounterConditionValue](e.EncounterConditionValueURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterConditionValue, nil
}

// Return a list of EncounterConditionValue resource
func (e Encounters) GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionValueList, err := request.GetResourceList[models.NamedResourceList](e.EncounterConditionValueURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionValueList, nil
}

// Return the the EncounterConditionValue resource URL
func (e Encounters) GetEncounterConditionValueURL() string {
	return e.EncounterConditionValueURL
}
