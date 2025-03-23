package pokemon

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Pokemon group interface
type IPokemon interface {
	GetAbility(nameOrId string) (*models.Ability, error)
	GetAbilityList(limit int, offset int) (*models.ResourceList, error)
	GetCharacteristic(id string) (*models.Characteristic, error)
	GetCharacteristicList(limit int, offset int) (*models.ResourceList, error)
	GetEggGroup(nameOrId string) (*models.EggGroup, error)
	GetEggGroupList(limit int, offset int) (*models.ResourceList, error)
	GetGender(nameOrId string) (*models.Gender, error)
	GetGenderList(limit int, offset int) (*models.ResourceList, error)
	GetGrowthRate(nameOrId string) (*models.GrowthRate, error)
	GetGrowthRateList(limit int, offset int) (*models.ResourceList, error)
	GetNature(nameOrId string) (*models.Nature, error)
	GetNatureList(limit int, offset int) (*models.ResourceList, error)
	GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error)
	GetPokeathlonStatList(limit int, offset int) (*models.ResourceList, error)
	GetPokemon(nameOrId string) (*models.Pokemon, error)
	GetPokemonList(limit int, offset int) (*models.ResourceList, error)
	GetPokemonColor(nameOrId string) (*models.PokemonColor, error)
	GetPokemonColorList(limit int, offset int) (*models.ResourceList, error)
	GetPokemonForm(nameOrId string) (*models.PokemonForm, error)
	GetPokemonFormList(limit int, offset int) (*models.ResourceList, error)
	GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error)
	GetPokemonHabitatList(limit int, offset int) (*models.ResourceList, error)
	GetPokemonShape(nameOrId string) (*models.PokemonShape, error)
	GetPokemonShapeList(limit int, offset int) (*models.ResourceList, error)
	GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error)
	GetPokemonSpeciesList(limit int, offset int) (*models.ResourceList, error)
	GetStat(nameOrId string) (*models.Stat, error)
	GetStatList(limit int, offset int) (*models.ResourceList, error)
	GetType(nameOrId string) (*models.Type, error)
	GetTypeList(limit int, offset int) (*models.ResourceList, error)
}

// Pokemon group struct
type Pokemon struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Pokemon resource group initialized")
}

// Return an instance of Pokmon resource group struct
func NewPokemonGroup() Pokemon {
	return Pokemon{
		Cache: cache.C,
	}
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
func (p Pokemon) GetAbilityList(limit int, offset int) (*models.ResourceList, error) {
	abilityList, err := request.GetResourceList(constants.AbilityEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return abilityList, nil
}

// Return a single Characteristic resource by ID
func (p Pokemon) GetCharacteristic(id string) (*models.Characteristic, error) {
	characteristic, err := request.GetSpecificResource[models.Characteristic](constants.CharacteristicEndpoint + id)
	if err != nil {
		return nil, err
	}
	return characteristic, nil
}

// Return a list of characteristic resource
func (p Pokemon) GetCharacteristicList(limit int, offset int) (*models.ResourceList, error) {
	characteristicList, err := request.GetResourceList(constants.CharacteristicEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return characteristicList, nil
}

// Return a single EggGroup resource by name or ID
func (p Pokemon) GetEggGroup(nameOrId string) (*models.EggGroup, error) {
	eggGroup, err := request.GetSpecificResource[models.EggGroup](constants.EggGroupEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return eggGroup, nil
}

// Return a list of EggGroup resource
func (p Pokemon) GetEggGroupList(limit int, offset int) (*models.ResourceList, error) {
	eggGroupList, err := request.GetResourceList(constants.EggGroupEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return eggGroupList, nil
}

// Return a single Gender by name or ID
func (p Pokemon) GetGender(nameOrId string) (*models.Gender, error) {
	gender, err := request.GetSpecificResource[models.Gender](constants.GenderEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return gender, nil
}

// Return a list of Gender resource
func (p Pokemon) GetGenderList(limit int, offset int) (*models.ResourceList, error) {
	genderList, err := request.GetResourceList(constants.GenderEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return genderList, nil
}

// Return a single GrowthRate by name or ID
func (p Pokemon) GetGrowthRate(nameOrId string) (*models.GrowthRate, error) {
	growthRate, err := request.GetSpecificResource[models.GrowthRate](constants.GrowthRateEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return growthRate, nil
}

// Return a list of GrowthRate resource
func (p Pokemon) GetGrowthRateList(limit int, offset int) (*models.ResourceList, error) {
	growthRateList, err := request.GetResourceList(constants.GrowthRateEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return growthRateList, nil
}

// Return a single Nature by name or ID
func (p Pokemon) GetNature(nameOrId string) (*models.Nature, error) {
	nature, err := request.GetSpecificResource[models.Nature](constants.NatureEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return nature, nil
}

// Return a list of Nature resource
func (p Pokemon) GetNatureList(limit int, offset int) (*models.ResourceList, error) {
	natureList, err := request.GetResourceList(constants.NatureEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return natureList, nil
}

// Return a single PokeathlonStat by name or ID
func (p Pokemon) GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error) {
	pokeathlonStat, err := request.GetSpecificResource[models.PokeathlonStat](constants.PokeathlonStatEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokeathlonStat, nil
}

// Return a list of PokeathlonStat resource
func (p Pokemon) GetPokeathlonStatList(limit int, offset int) (*models.ResourceList, error) {
	pokeathlonStatList, err := request.GetResourceList(constants.PokeathlonStatEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokeathlonStatList, nil
}

// Return a single Pokemon by name or ID
func (p Pokemon) GetPokemon(nameOrId string) (*models.Pokemon, error) {
	pokemon, err := request.GetSpecificResource[models.Pokemon](constants.PokemonEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Return a list of Pokemon resource
func (p Pokemon) GetPokemonList(limit int, offset int) (*models.ResourceList, error) {
	pokemonList, err := request.GetResourceList(constants.PokemonEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonList, nil
}

// Return a single PokemonColor by name or ID
func (p Pokemon) GetPokemonColor(nameOrId string) (*models.PokemonColor, error) {
	pokemonColor, err := request.GetSpecificResource[models.PokemonColor](constants.PokemonColorEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonColor, nil
}

// Return a list of PokemonColor resource
func (p Pokemon) GetPokemonColorList(limit int, offset int) (*models.ResourceList, error) {
	pokemonColorList, err := request.GetResourceList(constants.PokemonColorEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonColorList, nil
}

// Return a single PokemonForm by name or ID
func (p Pokemon) GetPokemonForm(nameOrId string) (*models.PokemonForm, error) {
	pokemonForm, err := request.GetSpecificResource[models.PokemonForm](constants.PokemonFormEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonForm, nil
}

// Return a list of PokemonForm resource
func (p Pokemon) GetPokemonFormList(limit int, offset int) (*models.ResourceList, error) {
	pokemonFormList, err := request.GetResourceList(constants.PokemonFormEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonFormList, nil
}

// Return a single PokemonHabitat by name or ID
func (p Pokemon) GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error) {
	pokemonHabitat, err := request.GetSpecificResource[models.PokemonHabitat](constants.PokemonHabitatEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonHabitat, nil
}

// Return a list of PokemonHabitat resource
func (p Pokemon) GetPokemonHabitatList(limit int, offset int) (*models.ResourceList, error) {
	pokemonHabitatList, err := request.GetResourceList(constants.PokemonHabitatEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonHabitatList, nil
}

// Return a single PokemonShape by name or ID
func (p Pokemon) GetPokemonShape(nameOrId string) (*models.PokemonShape, error) {
	pokemonShape, err := request.GetSpecificResource[models.PokemonShape](constants.PokemonShapeEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonShape, nil
}

// Return a list of PokemonShape resource
func (p Pokemon) GetPokemonShapeList(limit int, offset int) (*models.ResourceList, error) {
	pokemonShapeList, err := request.GetResourceList(constants.PokemonShapeEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonShapeList, nil
}

// Return a single PokemonSpecies by name or ID
func (p Pokemon) GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error) {
	pokemonSpecies, err := request.GetSpecificResource[models.PokemonSpecies](constants.PokemonSpeciesEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonSpecies, nil
}

// Return a list of PokemonSpecies resource
func (p Pokemon) GetPokemonSpeciesList(limit int, offset int) (*models.ResourceList, error) {
	pokemonSpeciesList, err := request.GetResourceList(constants.PokemonSpeciesEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonSpeciesList, nil
}

// Return a single Stat by name or ID
func (p Pokemon) GetStat(nameOrId string) (*models.Stat, error) {
	stat, err := request.GetSpecificResource[models.Stat](constants.StatEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// Return a list of Stat resource
func (p Pokemon) GetStatList(limit int, offset int) (*models.ResourceList, error) {
	statList, err := request.GetResourceList(constants.StatEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return statList, nil
}

// Return a single Type by name or ID
func (p Pokemon) GetType(nameOrId string) (*models.Type, error) {
	typeResource, err := request.GetSpecificResource[models.Type](constants.TypeEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return typeResource, nil
}

// Return a list of Type resource
func (p Pokemon) GetTypeList(limit int, offset int) (*models.ResourceList, error) {
	typeList, err := request.GetResourceList(constants.TypeEndpoint, limit, offset)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}
