package pokemon

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Pokemon group interface
type IPokemon interface {
	GetAbility(nameOrId string) (*models.Ability, error)
	GetAbilityList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Pokemon group struct
type Pokemon struct{}

// Return an instance of Pokmon resource group struct
func NewPokemonGroup() Pokemon {
	return Pokemon{}
}

// Return a single Ability resource by name or ID
func (p Pokemon) GetAbility(nameOrId string) (*models.Ability, error) {
	ability, err := request.GetSpecificResource[models.Ability](constants.AbilityEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return ability, nil
}

// Return a list of Ability resource
func (p Pokemon) GetAbilityList(options models.PaginationOptions) (*models.ResourceList, error) {
	abilityList, err := request.GetResourceList(constants.AbilityEndpoint, options)
	if err != nil {
		return nil, err
	}
	return abilityList, nil
}
