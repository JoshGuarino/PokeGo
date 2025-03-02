package utility

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Utility group interface
type IUtility interface {
	GetLanguage(nameOrId string) (*models.Language, error)
	GetLanguageList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Utility group struct
type Utility struct{}

// Return an instance of Utility resource group struct
func NewUtilityGroup() Utility {
	return Utility{}
}

// Return a single Language resource by name or ID
func (u Utility) GetLanguage(nameOrId string) (*models.Language, error) {
	language, err := request.GetSpecificResource[models.Language](constants.LanguageEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return language, nil
}

// Return a list of Language resource
func (u Utility) GetLanguageList(options models.PaginationOptions) (*models.ResourceList, error) {
	languageList, err := request.GetResourceList(constants.LanguageEndpoint, options)
	if err != nil {
		return nil, err
	}
	return languageList, nil
}
