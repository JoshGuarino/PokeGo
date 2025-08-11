package encounters

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Encounters group resource endpoints
const (
	EncounterMethodEndpoint         = "/encounter-method/"
	EncounterConditionEndpoint      = "/encounter-condition/"
	EncounterConditionValueEndpoint = "/encounter-condition-value/"
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
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Encounters resource group initialized")
}

// Return an instance of Encounters resource group struct
func NewEncountersGroup() Encounters {
	return Encounters{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single EncounterMethod resource by name or ID
func (e Encounters) GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error) {
	encounterMethod, err := request.GetResource[models.EncounterMethod](e.GetEncounterMethodURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterMethod, nil
}

// Return a list of EncounterMethod resource
func (e Encounters) GetEncounterMethodList(limit int, offest int) (*models.NamedResourceList, error) {
	encounterMethodList, err := request.GetResourceList[models.NamedResourceList](e.GetEncounterMethodURL(), limit, offest)
	if err != nil {
		return nil, err
	}
	return encounterMethodList, nil
}

// Return the EncounterMethod resource URL
func (e Encounters) GetEncounterMethodURL() string {
	return e.Env.URL() + EncounterMethodEndpoint
}

// Return a single EncounterCondition resource by name or ID
func (e Encounters) GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error) {
	encounterCondition, err := request.GetResource[models.EncounterCondition](e.GetEncounterConditionURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterCondition, nil
}

// Return a list of EncounterCondition resource
func (e Encounters) GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionList, err := request.GetResourceList[models.NamedResourceList](e.GetEncounterConditionURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionList, nil
}

// Return the EncounterCondition resource URL
func (e Encounters) GetEncounterConditionURL() string {
	return e.Env.URL() + EncounterConditionEndpoint
}

// Return a single EncounterConditionValue resource by name or ID
func (e Encounters) GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error) {
	encounterConditionValue, err := request.GetResource[models.EncounterConditionValue](e.GetEncounterConditionValueURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterConditionValue, nil
}

// Return a list of EncounterConditionValue resource
func (e Encounters) GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionValueList, err := request.GetResourceList[models.NamedResourceList](e.GetEncounterConditionValueURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionValueList, nil
}

// Return the EncounterConditionValue resource URL
func (e Encounters) GetEncounterConditionValueURL() string {
	return e.Env.URL() + EncounterConditionValueEndpoint
}
