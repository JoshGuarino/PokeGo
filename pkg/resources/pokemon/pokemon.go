package pokemon

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Pokemon group interface
type IPokemon interface {
	GetAbility(nameOrId string) (*models.Ability, error)
	GetAbilityList(limit int, offset int) (*models.NamedResourceList, error)
	GetAbilityURL() string
	GetCharacteristic(id string) (*models.Characteristic, error)
	GetCharacteristicList(limit int, offset int) (*models.ResourceList, error)
	GetCharacteristicURL() string
	GetEggGroup(nameOrId string) (*models.EggGroup, error)
	GetEggGroupList(limit int, offset int) (*models.NamedResourceList, error)
	GetEggGroupURL() string
	GetGender(nameOrId string) (*models.Gender, error)
	GetGenderList(limit int, offset int) (*models.NamedResourceList, error)
	GetGenderURL() string
	GetGrowthRate(nameOrId string) (*models.GrowthRate, error)
	GetGrowthRateList(limit int, offset int) (*models.NamedResourceList, error)
	GetGrowthRateURL() string
	GetNature(nameOrId string) (*models.Nature, error)
	GetNatureList(limit int, offset int) (*models.NamedResourceList, error)
	GetNatureURL() string
	GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error)
	GetPokeathlonStatList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokeathlonStatURL() string
	GetPokemon(nameOrId string) (*models.Pokemon, error)
	GetPokemonList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonURL() string
	GetPokemonLocationAreas(nameOrId string) ([]*models.LocationAreaEncounter, error)
	GetPokemonColor(nameOrId string) (*models.PokemonColor, error)
	GetPokemonColorList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonColorURL() string
	GetPokemonForm(nameOrId string) (*models.PokemonForm, error)
	GetPokemonFormList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonFormURL() string
	GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error)
	GetPokemonHabitatList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonHabitatURL() string
	GetPokemonShape(nameOrId string) (*models.PokemonShape, error)
	GetPokemonShapeList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonShapeURL() string
	GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error)
	GetPokemonSpeciesList(limit int, offset int) (*models.NamedResourceList, error)
	GetPokemonSpeciesURL() string
	GetStat(nameOrId string) (*models.Stat, error)
	GetStatList(limit int, offset int) (*models.NamedResourceList, error)
	GetStatURL() string
	GetType(nameOrId string) (*models.Type, error)
	GetTypeList(limit int, offset int) (*models.NamedResourceList, error)
	GetTypeURL() string
}

// Pokemon group struct
type Pokemon struct {
	AbilityURL        string
	CharacteristicURL string
	EggGroupURL       string
	GenderURL         string
	GrowthRateURL     string
	NatureURL         string
	PokeathlonStatURL string
	PokemonURL        string
	PokemonColorURL   string
	PokemonFormURL    string
	PokemonHabitatURL string
	PokemonShapeURL   string
	PokemonSpeciesURL string
	StatURL           string
	TypeURL           string
	Cache             *cache.Cache
}

// Initialize function
func init() {
	log.Info("Pokemon resource group initialized")
}

// Return an instance of Pokmon resource group struct
func NewPokemonGroup() Pokemon {
	url := endpoints.BaseURL
	return Pokemon{
		AbilityURL:        url + endpoints.Ability,
		CharacteristicURL: url + endpoints.Characteristic,
		EggGroupURL:       url + endpoints.EggGroup,
		GenderURL:         url + endpoints.Gender,
		GrowthRateURL:     url + endpoints.GrowthRate,
		NatureURL:         url + endpoints.Nature,
		PokeathlonStatURL: url + endpoints.PokeathlonStat,
		PokemonURL:        url + endpoints.Pokemon,
		PokemonColorURL:   url + endpoints.PokemonColor,
		PokemonFormURL:    url + endpoints.PokemonForm,
		PokemonHabitatURL: url + endpoints.PokemonHabitat,
		PokemonShapeURL:   url + endpoints.PokemonShape,
		PokemonSpeciesURL: url + endpoints.PokemonSpecies,
		StatURL:           url + endpoints.Stat,
		TypeURL:           url + endpoints.Type,
		Cache:             cache.C,
	}
}

// Return a single Ability resource by name or ID
func (p Pokemon) GetAbility(nameOrId string) (*models.Ability, error) {
	ability, err := request.GetResource[models.Ability](p.AbilityURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return ability, nil
}

// Return a list of Ability resource
func (p Pokemon) GetAbilityList(limit int, offset int) (*models.NamedResourceList, error) {
	abilityList, err := request.GetResourceList[models.NamedResourceList](p.AbilityURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return abilityList, nil
}

// Return the Ability resource url
func (p Pokemon) GetAbilityURL() string {
	return p.AbilityURL
}

// Return a single Characteristic resource by ID
func (p Pokemon) GetCharacteristic(id string) (*models.Characteristic, error) {
	characteristic, err := request.GetResource[models.Characteristic](p.CharacteristicURL + id)
	if err != nil {
		return nil, err
	}
	return characteristic, nil
}

// Return a list of characteristic resource
func (p Pokemon) GetCharacteristicList(limit int, offset int) (*models.ResourceList, error) {
	characteristicList, err := request.GetResourceList[models.ResourceList](p.CharacteristicURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return characteristicList, nil
}

// Return the Characteristic resource url
func (p Pokemon) GetCharacteristicURL() string {
	return p.CharacteristicURL
}

// Return a single EggGroup resource by name or ID
func (p Pokemon) GetEggGroup(nameOrId string) (*models.EggGroup, error) {
	eggGroup, err := request.GetResource[models.EggGroup](p.EggGroupURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return eggGroup, nil
}

// Return a list of EggGroup resource
func (p Pokemon) GetEggGroupList(limit int, offset int) (*models.NamedResourceList, error) {
	eggGroupList, err := request.GetResourceList[models.NamedResourceList](p.EggGroupURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return eggGroupList, nil
}

// Return the EggGroup resource url
func (p Pokemon) GetEggGroupURL() string {
	return p.EggGroupURL
}

// Return a single Gender by name or ID
func (p Pokemon) GetGender(nameOrId string) (*models.Gender, error) {
	gender, err := request.GetResource[models.Gender](p.GenderURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return gender, nil
}

// Return a list of Gender resource
func (p Pokemon) GetGenderList(limit int, offset int) (*models.NamedResourceList, error) {
	genderList, err := request.GetResourceList[models.NamedResourceList](p.GenderURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return genderList, nil
}

// Return the Gender resource url
func (p Pokemon) GetGenderURL() string {
	return p.GenderURL
}

// Return a single GrowthRate by name or ID
func (p Pokemon) GetGrowthRate(nameOrId string) (*models.GrowthRate, error) {
	growthRate, err := request.GetResource[models.GrowthRate](p.GrowthRateURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return growthRate, nil
}

// Return a list of GrowthRate resource
func (p Pokemon) GetGrowthRateList(limit int, offset int) (*models.NamedResourceList, error) {
	growthRateList, err := request.GetResourceList[models.NamedResourceList](p.GrowthRateURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return growthRateList, nil
}

// Return the GrowthRate resource url
func (p Pokemon) GetGrowthRateURL() string {
	return p.GrowthRateURL
}

// Return a single Nature by name or ID
func (p Pokemon) GetNature(nameOrId string) (*models.Nature, error) {
	nature, err := request.GetResource[models.Nature](p.NatureURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return nature, nil
}

// Return a list of Nature resource
func (p Pokemon) GetNatureList(limit int, offset int) (*models.NamedResourceList, error) {
	natureList, err := request.GetResourceList[models.NamedResourceList](p.NatureURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return natureList, nil
}

// Return the Nature resource url
func (p Pokemon) GetNatureURL() string {
	return p.NatureURL
}

// Return a single PokeathlonStat by name or ID
func (p Pokemon) GetPokeathlonStat(nameOrId string) (*models.PokeathlonStat, error) {
	pokeathlonStat, err := request.GetResource[models.PokeathlonStat](p.PokeathlonStatURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokeathlonStat, nil
}

// Return a list of PokeathlonStat resource
func (p Pokemon) GetPokeathlonStatList(limit int, offset int) (*models.NamedResourceList, error) {
	pokeathlonStatList, err := request.GetResourceList[models.NamedResourceList](p.PokeathlonStatURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokeathlonStatList, nil
}

// Return the PokeathlonStat resource url
func (p Pokemon) GetPokeathlonStatURL() string {
	return p.PokeathlonStatURL
}

// Return a single Pokemon by name or ID
func (p Pokemon) GetPokemon(nameOrId string) (*models.Pokemon, error) {
	pokemon, err := request.GetResource[models.Pokemon](p.PokemonURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Return a list of Pokemon resource
func (p Pokemon) GetPokemonList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonList, err := request.GetResourceList[models.NamedResourceList](p.PokemonURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonList, nil
}

// Return the Pokemon resource url
func (p Pokemon) GetPokemonURL() string {
	return p.PokemonURL
}

// Return a single LocationAreaEncounter by name or ID
func (p Pokemon) GetPokemonLocationAreas(nameOrId string) ([]*models.LocationAreaEncounter, error) {
	pokemonLocationAreas, err := request.GetResourceSlice[models.LocationAreaEncounter](p.PokemonURL + nameOrId + "/encounters")
	if err != nil {
		return nil, err
	}
	return pokemonLocationAreas, nil
}

// Return a single PokemonColor by name or ID
func (p Pokemon) GetPokemonColor(nameOrId string) (*models.PokemonColor, error) {
	pokemonColor, err := request.GetResource[models.PokemonColor](p.PokemonColorURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonColor, nil
}

// Return a list of PokemonColor resource
func (p Pokemon) GetPokemonColorList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonColorList, err := request.GetResourceList[models.NamedResourceList](p.PokemonColorURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonColorList, nil
}

// Return the PokemonColor resource url
func (p Pokemon) GetPokemonColorURL() string {
	return p.PokemonColorURL
}

// Return a single PokemonForm by name or ID
func (p Pokemon) GetPokemonForm(nameOrId string) (*models.PokemonForm, error) {
	pokemonForm, err := request.GetResource[models.PokemonForm](p.PokemonFormURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonForm, nil
}

// Return a list of PokemonForm resource
func (p Pokemon) GetPokemonFormList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonFormList, err := request.GetResourceList[models.NamedResourceList](p.PokemonFormURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonFormList, nil
}

// Return the PokemonForm resource url
func (p Pokemon) GetPokemonFormURL() string {
	return p.PokemonFormURL
}

// Return a single PokemonHabitat by name or ID
func (p Pokemon) GetPokemonHabitat(nameOrId string) (*models.PokemonHabitat, error) {
	pokemonHabitat, err := request.GetResource[models.PokemonHabitat](p.PokemonHabitatURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonHabitat, nil
}

// Return a list of PokemonHabitat resource
func (p Pokemon) GetPokemonHabitatList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonHabitatList, err := request.GetResourceList[models.NamedResourceList](p.PokemonHabitatURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonHabitatList, nil
}

// Return the PokemonHabitat resource url
func (p Pokemon) GetPokemonHabitatURL() string {
	return p.PokemonHabitatURL
}

// Return a single PokemonShape by name or ID
func (p Pokemon) GetPokemonShape(nameOrId string) (*models.PokemonShape, error) {
	pokemonShape, err := request.GetResource[models.PokemonShape](p.PokemonShapeURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonShape, nil
}

// Return a list of PokemonShape resource
func (p Pokemon) GetPokemonShapeList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonShapeList, err := request.GetResourceList[models.NamedResourceList](p.PokemonShapeURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonShapeList, nil
}

// Return the PokemonShape resource url
func (p Pokemon) GetPokemonShapeURL() string {
	return p.PokemonShapeURL
}

// Return a single PokemonSpecies by name or ID
func (p Pokemon) GetPokemonSpecies(nameOrId string) (*models.PokemonSpecies, error) {
	pokemonSpecies, err := request.GetResource[models.PokemonSpecies](p.PokemonSpeciesURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return pokemonSpecies, nil
}

// Return a list of PokemonSpecies resource
func (p Pokemon) GetPokemonSpeciesList(limit int, offset int) (*models.NamedResourceList, error) {
	pokemonSpeciesList, err := request.GetResourceList[models.NamedResourceList](p.PokemonSpeciesURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return pokemonSpeciesList, nil
}

// Return the PokemonSpecies resource url
func (p Pokemon) GetPokemonSpeciesURL() string {
	return p.PokemonSpeciesURL
}

// Return a single Stat by name or ID
func (p Pokemon) GetStat(nameOrId string) (*models.Stat, error) {
	stat, err := request.GetResource[models.Stat](p.StatURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// Return a list of Stat resource
func (p Pokemon) GetStatList(limit int, offset int) (*models.NamedResourceList, error) {
	statList, err := request.GetResourceList[models.NamedResourceList](p.StatURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return statList, nil
}

// Return the Stat resource url
func (p Pokemon) GetStatURL() string {
	return p.StatURL
}

// Return a single Type by name or ID
func (p Pokemon) GetType(nameOrId string) (*models.Type, error) {
	typeResource, err := request.GetResource[models.Type](p.TypeURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return typeResource, nil
}

// Return a list of Type resource
func (p Pokemon) GetTypeList(limit int, offset int) (*models.NamedResourceList, error) {
	typeList, err := request.GetResourceList[models.NamedResourceList](p.TypeURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}

// Return the Type resource url
func (p Pokemon) GetTypeURL() string {
	return p.TypeURL
}
