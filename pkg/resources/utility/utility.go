package utility

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Utility group interface
type IUtility interface {
	GetLanguage(nameOrId string) (*models.Language, error)
	GetLanguageList(limit int, offset int) (*models.NamedResourceList, error)
}

// Utility group struct
type Utility struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Utility resource group initialized")
}

// Return an instance of Utility resource group struct
func NewUtilityGroup() Utility {
	return Utility{
		Cache: cache.C,
	}
}

// Return a single Language resource by name or ID
func (u Utility) GetLanguage(nameOrId string) (*models.Language, error) {
	language, err := request.GetResource[models.Language](endpoints.Language + nameOrId)
	if err != nil {
		return nil, err
	}
	return language, nil
}

// Return a list of Language resource
func (u Utility) GetLanguageList(limit int, offset int) (*models.NamedResourceList, error) {
	languageList, err := request.GetResourceList[models.NamedResourceList](endpoints.Language, limit, offset)
	if err != nil {
		return nil, err
	}
	return languageList, nil
}
