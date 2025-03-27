package pokemon

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Pokemon group interface
type IPokemon interface {
	GetAbility(nameOrId string) (*models.Ability, error)
	GetAbilityList(limit int, offset int) (*models.NamedResourceList, error)
	GetCharacteristic(id string) (*models.Characteristic, error)
	GetCharacteristicList(limit int, offset int) (*models.ResourceList, error)
	GetEggGroup(nameOrId string) (*models.EggGroup, error)
	GetEggGroupList(limit int, offset int) (*models.NamedResourceList, error)
	GetGender(nameOrId string) (*models.Gender, error)
	GetGenderList(limit int, offset int) (*models.NamedResourceList, error)
	GetGrowthRate(nameOrId string) (*models.GrowthRate, error)
	GetGrowthRateList(limit int, offset int) (*models.NamedResourceList, error)
	GetNature(nameOrId string) (*models.Nature, error)
	GetNatureList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error)
	GetPokeathlonStatList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemon(nameOrId string) (*models.Pokemon, error)
	GetPokemonList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonLocationAreas(nameOrId string) ([]*models.LocationAreaEncounter, error)
	GetPokemonColor(nameOrId string) (*models.PokemonColor, error)
	GetPokemonColorList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonForm(nameOrId string) (*models.PokemonForm, error)
	GetPokemonFormList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error)
	GetPokemonHabitatList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonShape(nameOrId string) (*models.PokemonShape, error)
	GetPokemonShapeList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error)
	GetPokemonSpeciesList(limit int, offset int) (*models.NamedResourceList, error)
	GetStat(nameOrId string) (*models.Stat, error)
	GetStatList(limit int, offset int) (*models.NamedResourceList, error)
	GetType(nameOrId string) (*models.Type, error)
	GetTypeList(limit int, offset int) (*models.NamedResourceList, error)
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
	ability, err := request.GetResource[models.Ability](endpoints.Ability + nameOrId)
	if err != nil {
		return nil, err
	}
	return ability, nil
}

// Return a list of Ability resource
func (p Pokemon) GetAbilityList(limit int, offset int) (*models.NamedResourceList, error) {
	abilityList, err := request.GetResourceList[models.NamedResourceList](endpoints.Ability, limit, offset)
	if err != nil {
		return nil, err
	}
	return abilityList, nil
}

// Return a single Characteristic resource by ID
func (p Pokemon) GetCharacteristic(id string) (*models.Characteristic, error) {
	characteristic, err := request.GetResource[models.Characteristic](endpoints.Characteristic + id)
	if err != nil {
		return nil, err
	}
	return characteristic, nil
}

// Return a list of characteristic resource
func (p Pokemon) GetCharacteristicList(limit int, offset int) (*models.ResourceList, error) {
	characteristicList, err := request.GetResourceList[models.ResourceList](endpoints.Characteristic, limit, offset)
	if err != nil {
		return nil, err
	}
	return characteristicList, nil
}

// Return a single EggGroup resource by name or ID
func (p Pokemon) GetEggGroup(nameOrId string) (*models.EggGroup, error) {
	eggGroup, err := request.GetResource[models.EggGroup](endpoints.EggGroup + nameOrId)
	if err != nil {
		return nil, err
	}
	return eggGroup, nil
}

// Return a list of EggGroup resource
func (p Pokemon) GetEggGroupList(limit int, offset int) (*models.NamedResourceList, error) {
	eggGroupList, err := request.GetResourceList[models.NamedResourceList](endpoints.EggGroup, limit, offset)
	if err != nil {
		return nil, err
	}
	return eggGroupList, nil
}

// Return a single Gender by name or ID
func (p Pokemon) GetGender(nameOrId string) (*models.Gender, error) {
	gender, err := request.GetResource[models.Gender](endpoints.Gender + nameOrId)
	if err != nil {
		return nil, err
	}
	return gender, nil
}

// Return a list of Gender resource
func (p Pokemon) GetGenderList(limit int, offset int) (*models.NamedResourceList, error) {
	genderList, err := request.GetResourceList[models.NamedResourceList](endpoints.Gender, limit, offset)
	if err != nil {
		return nil, err
	}
	return genderList, nil
}

// Return a single GrowthRate by name or ID
func (p Pokemon) GetGrowthRate(nameOrId string) (*models.GrowthRate, error) {
	growthRate, err := request.GetResource[models.GrowthRate](endpoints.GrowthRate + nameOrId)
	if err != nil {
		return nil, err
	}
	return growthRate, nil
}

// Return a list of GrowthRate resource
func (p Pokemon) GetGrowthRateList(limit int, offset int) (*models.NamedResourceList, error) {
	growthRateList, err := request.GetResourceList[models.NamedResourceList](endpoints.GrowthRate, limit, offset)
	if err != nil {
		return nil, err
	}
	return growthRateList, nil
}

// Return a single Nature by name or ID
func (p Pokemon) GetNature(nameOrId string) (*models.Nature, error) {
	nature, err := request.GetResource[models.Nature](endpoints.Nature + nameOrId)
	if err != nil {
		return nil, err
	}
	return nature, nil
}

// Return a list of Nature resource
func (p Pokemon) GetNatureList(limit int, offset int) (*models.NamedResourceList, error) {
	natureList, err := request.GetResourceList[models.NamedResourceList](endpoints.Nature, limit, offset)
	if err != nil {
		return nil, err
	}
	return natureList, nil
}

// Return a single PokeathlonStat by name or ID
func (p Pokemon) GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error) {
	pokeathlonStat, err := request.GetResource[models.PokeathlonStat](endpoints.PokeathlonStat + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokeathlonStat, nil
}

// Return a list of PokeathlonStat resource
func (p Pokemon) GetPokeathlonStatList(limit int, offset int) (*models.NamedResourceList, error) {
	pokeathlonStatList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokeathlonStat, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokeathlonStatList, nil
}

// Return a single Pokemon by name or ID
func (p Pokemon) GetPokemon(nameOrId string) (*models.Pokemon, error) {
	pokemon, err := request.GetResource[models.Pokemon](endpoints.Pokemon + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Return a list of Pokemon resource
func (p Pokemon) GetPokemonList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonList, err := request.GetResourceList[models.NamedResourceList](endpoints.Pokemon, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonList, nil
}

// Return a single LocationAreaEncounter by name or ID
func (p Pokemon) GetPokemonLocationAreas(nameOrId string) ([]*models.LocationAreaEncounter, error) {
	pokemonLocationAreas, err := request.GetResourceSlice[models.LocationAreaEncounter](endpoints.Pokemon + nameOrId + "/encounters")
	if err != nil {
		return nil, err
	}
	return pokemonLocationAreas, nil
}

// Return a single PokemonColor by name or ID
func (p Pokemon) GetPokemonColor(nameOrId string) (*models.PokemonColor, error) {
	pokemonColor, err := request.GetResource[models.PokemonColor](endpoints.PokemonColor + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonColor, nil
}

// Return a list of PokemonColor resource
func (p Pokemon) GetPokemonColorList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonColorList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokemonColor, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonColorList, nil
}

// Return a single PokemonForm by name or ID
func (p Pokemon) GetPokemonForm(nameOrId string) (*models.PokemonForm, error) {
	pokemonForm, err := request.GetResource[models.PokemonForm](endpoints.PokemonForm + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonForm, nil
}

// Return a list of PokemonForm resource
func (p Pokemon) GetPokemonFormList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonFormList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokemonForm, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonFormList, nil
}

// Return a single PokemonHabitat by name or ID
func (p Pokemon) GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error) {
	pokemonHabitat, err := request.GetResource[models.PokemonHabitat](endpoints.PokemonHabitat + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonHabitat, nil
}

// Return a list of PokemonHabitat resource
func (p Pokemon) GetPokemonHabitatList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonHabitatList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokemonHabitat, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonHabitatList, nil
}

// Return a single PokemonShape by name or ID
func (p Pokemon) GetPokemonShape(nameOrId string) (*models.PokemonShape, error) {
	pokemonShape, err := request.GetResource[models.PokemonShape](endpoints.PokemonShape + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonShape, nil
}

// Return a list of PokemonShape resource
func (p Pokemon) GetPokemonShapeList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonShapeList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokemonShape, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonShapeList, nil
}

// Return a single PokemonSpecies by name or ID
func (p Pokemon) GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error) {
	pokemonSpecies, err := request.GetResource[models.PokemonSpecies](endpoints.PokemonSpecies + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonSpecies, nil
}

// Return a list of PokemonSpecies resource
func (p Pokemon) GetPokemonSpeciesList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonSpeciesList, err := request.GetResourceList[models.NamedResourceList](endpoints.PokemonSpecies, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonSpeciesList, nil
}

// Return a single Stat by name or ID
func (p Pokemon) GetStat(nameOrId string) (*models.Stat, error) {
	stat, err := request.GetResource[models.Stat](endpoints.Stat + nameOrId)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// Return a list of Stat resource
func (p Pokemon) GetStatList(limit int, offset int) (*models.NamedResourceList, error) {
	statList, err := request.GetResourceList[models.NamedResourceList](endpoints.Stat, limit, offset)
	if err != nil {
		return nil, err
	}
	return statList, nil
}

// Return a single Type by name or ID
func (p Pokemon) GetType(nameOrId string) (*models.Type, error) {
	typeResource, err := request.GetResource[models.Type](endpoints.Type + nameOrId)
	if err != nil {
		return nil, err
	}
	return typeResource, nil
}

// Return a list of Type resource
func (p Pokemon) GetTypeList(limit int, offset int) (*models.NamedResourceList, error) {
	typeList, err := request.GetResourceList[models.NamedResourceList](endpoints.Type, limit, offset)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}
