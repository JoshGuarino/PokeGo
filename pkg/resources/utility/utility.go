package utility

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Utility group resource endpoints
const (
	LanguageEndpoint = "/language/"
)

// Utility group interface
type IUtility interface {
	GetLanguage(nameOrId string) (*models.Language, error)
	GetLanguageList(limit int, offset int) (*models.NamedResourceList, error)
	GetLanguageURL() string
}

// Utility group struct
type Utility struct {
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Utility resource group initialized")
}

// Return an instance of Utility resource group struct
func NewUtilityGroup() Utility {
	return Utility{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single Language resource by name or ID
func (u Utility) GetLanguage(nameOrId string) (*models.Language, error) {
	language, err := request.GetResource[models.Language](u.GetLanguageURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return language, nil
}

// Return a list of Language resource
func (u Utility) GetLanguageList(limit int, offset int) (*models.NamedResourceList, error) {
	languageList, err := request.GetResourceList[models.NamedResourceList](u.GetLanguageURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return languageList, nil
}

// Return the Language resource url
func (u Utility) GetLanguageURL() string {
	return u.Env.URL() + LanguageEndpoint
}
