package utility

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Utility group interface
type IUtility interface {
	GetLanguage(nameOrId string) (*models.Language, error)
	GetLanguageList(limit int, offset int) (*models.NamedResourceList, error)
	GetLanguageURL() string
}

// Utility group struct
type Utility struct {
	LanguageURL string
	Cache       *cache.Cache
}

// Initialize function
func init() {
	log.Info("Utility resource group initialized")
}

// Return an instance of Utility resource group struct
func NewUtilityGroup() Utility {
	url := endpoints.BaseURL
	return Utility{
		LanguageURL: url + endpoints.Language,
		Cache:       cache.C,
	}
}

// Return a single Language resource by name or ID
func (u Utility) GetLanguage(nameOrId string) (*models.Language, error) {
	language, err := request.GetResource[models.Language](u.LanguageURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return language, nil
}

// Return a list of Language resource
func (u Utility) GetLanguageList(limit int, offset int) (*models.NamedResourceList, error) {
	languageList, err := request.GetResourceList[models.NamedResourceList](u.LanguageURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return languageList, nil
}

// Return the Language resource url
func (u Utility) GetLanguageURL() string {
	return u.LanguageURL
}
