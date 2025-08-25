package pokemon

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var pokemon IPokemon = NewPokemonGroup()
var url string = environment.ENV.URL()

func TestNewPokemonGroup(t *testing.T) {
	pokemon := NewPokemonGroup()
	assert.IsType(t, Pokemon{}, pokemon, "Expected Pokemon instance to be returned")
}

func TestGetAbility(t *testing.T) {
	rById, _ := pokemon.GetAbility("1")
	rByName, _ := pokemon.GetAbility("stench")
	_, err := pokemon.GetAbility("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Ability resource")
	assert.Equal(t, "stench", rByName.Name, "Unexpected Name for Ability resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Ability{}, rById, "Expected Ability instance to be returned")
}

func TestGetAbilityList(t *testing.T) {
	rList, _ := pokemon.GetAbilityList(20, 0)
	rPage, _ := pokemon.GetAbilityList(1, 1)
	assert.Equal(t, "stench", rList.Results[0].Name, "Unexpected Name for Ability resource")
	assert.Equal(t, "drizzle", rPage.Results[0].Name, "Unexpected Name for Ability resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetAbilityURL(t *testing.T) {
	abilityURL := pokemon.GetAbilityURL()
	assert.Equal(t, url+AbilityEndpoint, abilityURL, "Unexpected Ability resource URL")
	assert.IsType(t, "", abilityURL, "Expected Ability resource URL to be a string")
}

func TestGetCharacteristic(t *testing.T) {
	rById, _ := pokemon.GetCharacteristic("1")
	_, err := pokemon.GetCharacteristic("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Characteristic resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Characteristic{}, rById, "Expected Characteristic instance to be returned")
}

func TestGetCharacteristicList(t *testing.T) {
	rList, _ := pokemon.GetCharacteristicList(20, 0)
	rPage, _ := pokemon.GetCharacteristicList(1, 1)
	assert.Equal(t, pokemon.GetCharacteristicURL()+"1/", rList.Results[0].URL, "Unexpected URL for Characteristic resource")
	assert.Equal(t, pokemon.GetCharacteristicURL()+"2/", rPage.Results[0].URL, "Unexpected URL for Characteristic resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.ResourceList{}, rList, "Expected ResourceList instance to be returned")
}

func TestGetCharacteristicURL(t *testing.T) {
	characteristicURL := pokemon.GetCharacteristicURL()
	assert.Equal(t, url+CharacteristicEndpoint, characteristicURL, "Unexpected Characteristic resource URL")
	assert.IsType(t, "", characteristicURL, "Expected Characteristic resource URL to be a string")
}

func TestGetEggGroup(t *testing.T) {
	rById, _ := pokemon.GetEggGroup("1")
	rByName, _ := pokemon.GetEggGroup("monster")
	_, err := pokemon.GetEggGroup("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EggGroup resource")
	assert.Equal(t, "monster", rByName.Name, "Unexpected Name for EggGroup resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.EggGroup{}, rById, "Expected EggGroup instance to be returned")
}

func TestGetEggGroupList(t *testing.T) {
	rList, _ := pokemon.GetEggGroupList(20, 0)
	rPage, _ := pokemon.GetEggGroupList(1, 1)
	assert.Equal(t, "monster", rList.Results[0].Name, "Unexpected Name for EggGroup resource")
	assert.Equal(t, "water1", rPage.Results[0].Name, "Unexpected Name for EggGroup resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetEggGroupURL(t *testing.T) {
	eggGroupURL := pokemon.GetEggGroupURL()
	assert.Equal(t, url+EggGroupEndpoint, eggGroupURL, "Unexpected EggGroup resource URL")
	assert.IsType(t, "", eggGroupURL, "Expected EggGroup resource URL to be a string")
}

func TestGetGender(t *testing.T) {
	rById, _ := pokemon.GetGender("1")
	rByName, _ := pokemon.GetGender("female")
	_, err := pokemon.GetGender("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Gender resource")
	assert.Equal(t, "female", rByName.Name, "Unexpected Name for Gender resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Gender{}, rById, "Expected Gender instance to be returned")
}

func TestGetGenderList(t *testing.T) {
	rList, _ := pokemon.GetGenderList(20, 0)
	rPage, _ := pokemon.GetGenderList(1, 1)
	assert.Equal(t, "female", rList.Results[0].Name, "Unexpected Name for Gender resource")
	assert.Equal(t, "male", rPage.Results[0].Name, "Unexpected Name for Gender resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetGenderURL(t *testing.T) {
	genderURL := pokemon.GetGenderURL()
	assert.Equal(t, url+GenderEndpoint, genderURL, "Unexpected Gender resource URL")
	assert.IsType(t, "", genderURL, "Expected Gender resource URL to be a string")
}

func TestGetGrowthRate(t *testing.T) {
	rById, _ := pokemon.GetGrowthRate("1")
	rByName, _ := pokemon.GetGrowthRate("slow")
	_, err := pokemon.GetGrowthRate("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for GrowthRate resource")
	assert.Equal(t, "slow", rByName.Name, "Unexpected Name for GrowthRate resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.GrowthRate{}, rById, "Expected GrowthRate instance to be returned")
}

func TestGetGrowthRateList(t *testing.T) {
	rList, _ := pokemon.GetGrowthRateList(20, 0)
	rPage, _ := pokemon.GetGrowthRateList(1, 1)
	assert.Equal(t, "slow", rList.Results[0].Name, "Unexpected Name for GrowthRate resource")
	assert.Equal(t, "medium", rPage.Results[0].Name, "Unexpected Name for GrowthRate resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetGrowthRateURL(t *testing.T) {
	growthRateURL := pokemon.GetGrowthRateURL()
	assert.Equal(t, url+GrowthRateEndpoint, growthRateURL, "Unexpected GrowthRate resource URL")
	assert.IsType(t, "", growthRateURL, "Expected GrowthRate resource URL to be a string")
}

func TestGetNature(t *testing.T) {
	rById, _ := pokemon.GetNature("1")
	rByName, _ := pokemon.GetNature("hardy")
	_, err := pokemon.GetNature("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Nature resource")
	assert.Equal(t, "hardy", rByName.Name, "Unexpected Name for Nature resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Nature{}, rById, "Expected Nature instance to be returned")
}

func TestGetNatureList(t *testing.T) {
	rList, _ := pokemon.GetNatureList(20, 0)
	rPage, _ := pokemon.GetNatureList(1, 1)
	assert.Equal(t, "hardy", rList.Results[0].Name, "Unexpected Name for Nature resource")
	assert.Equal(t, "bold", rPage.Results[0].Name, "Unexpected Name for Nature resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetNatureURL(t *testing.T) {
	natureURL := pokemon.GetNatureURL()
	assert.Equal(t, url+NatureEndpoint, natureURL, "Unexpected Nature resource URL")
	assert.IsType(t, "", natureURL, "Expected Nature resource URL to be a string")
}

func TestGetPokeathlonStat(t *testing.T) {
	rById, _ := pokemon.GetPokeathlonStat("1")
	rByName, _ := pokemon.GetPokeathlonStat("speed")
	_, err := pokemon.GetPokeathlonStat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokeathlonStat resource")
	assert.Equal(t, "speed", rByName.Name, "Unexpected Name for PokeathlonStat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokeathlonStat{}, rById, "Expected PokeathlonStat instance to be returned")
}

func TestGetPokeathlonStatList(t *testing.T) {
	rList, _ := pokemon.GetPokeathlonStatList(20, 0)
	rPage, _ := pokemon.GetPokeathlonStatList(1, 1)
	assert.Equal(t, "speed", rList.Results[0].Name, "Unexpected Name for PokeathlonStat resource")
	assert.Equal(t, "power", rPage.Results[0].Name, "Unexpected Name for PokeathlonStat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokeathlonStatURL(t *testing.T) {
	pokeathlonStatURL := pokemon.GetPokeathlonStatURL()
	assert.Equal(t, url+PokeathlonStatEndpoint, pokeathlonStatURL, "Unexpected PokeathlonStat resource URL")
	assert.IsType(t, "", pokeathlonStatURL, "Expected PokeathlonStat resource URL to be a string")
}

func TestGetPokemon(t *testing.T) {
	rById, _ := pokemon.GetPokemon("1")
	rByName, _ := pokemon.GetPokemon("bulbasaur")
	_, err := pokemon.GetPokemon("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Pokemon resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for Pokemon resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Pokemon{}, rById, "Expected Pokemon instance to be returned")
}

func TestGetPokemonList(t *testing.T) {
	rList, _ := pokemon.GetPokemonList(20, 0)
	rPage, _ := pokemon.GetPokemonList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for Pokemon resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for Pokemon resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonURL(t *testing.T) {
	pokemonURL := pokemon.GetPokemonURL()
	assert.Equal(t, url+PokemonEndpoint, pokemonURL, "Unexpected Pokemon resource URL")
	assert.IsType(t, "", pokemonURL, "Expected Pokemon resource URL to be a string")
}

func TestGetPokemonLocationAreas(t *testing.T) {
	rById, _ := pokemon.GetPokemonLocationAreas("1")
	rByName, _ := pokemon.GetPokemonLocationAreas("bulbasaur")
	_, err := pokemon.GetPokemonLocationAreas("test")
	assert.Equal(t, "cerulean-city-area", rById[0].LocationArea.Name, "Unexpected ID for PokemonLocationArea resource")
	assert.Equal(t, "cerulean-city-area", rByName[0].LocationArea.Name, "Unexpected Name for PokemonLocationArea resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, []*models.LocationAreaEncounter{}, rById, "Expected LocationAreaEncounter instance to be returned")
}

func TestGetPokemonColor(t *testing.T) {
	rById, _ := pokemon.GetPokemonColor("1")
	rByName, _ := pokemon.GetPokemonColor("black")
	_, err := pokemon.GetPokemonColor("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonColor resource")
	assert.Equal(t, "black", rByName.Name, "Unexpected Name for PokemonColor resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokemonColor{}, rById, "Expected PokemonColor instance to be returned")
}

func TestGetPokemonColorList(t *testing.T) {
	rList, _ := pokemon.GetPokemonColorList(20, 0)
	rPage, _ := pokemon.GetPokemonColorList(1, 1)
	assert.Equal(t, "black", rList.Results[0].Name, "Unexpected Name for PokemonColor resource")
	assert.Equal(t, "blue", rPage.Results[0].Name, "Unexpected Name for PokemonColor resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonColorURL(t *testing.T) {
	pokemonColorURL := pokemon.GetPokemonColorURL()
	assert.Equal(t, url+PokemonColorEndpoint, pokemonColorURL, "Unexpected PokemonColor resource URL")
	assert.IsType(t, "", pokemonColorURL, "Expected PokemonColor resource URL to be a string")
}

func TestGetPokemonForm(t *testing.T) {
	rById, _ := pokemon.GetPokemonForm("1")
	rByName, _ := pokemon.GetPokemonForm("bulbasaur")
	_, err := pokemon.GetPokemonForm("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonForm resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for PokemonForm resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokemonForm{}, rById, "Expected PokemonForm instance to be returned")
}

func TestGetPokemonFormList(t *testing.T) {
	rList, _ := pokemon.GetPokemonFormList(20, 0)
	rPage, _ := pokemon.GetPokemonFormList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for PokemonForm resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for PokemonForm resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonFormURL(t *testing.T) {
	pokemonFormURL := pokemon.GetPokemonFormURL()
	assert.Equal(t, url+PokemonFormEndpoint, pokemonFormURL, "Unexpected PokemonForm resource URL")
	assert.IsType(t, "", pokemonFormURL, "Expected PokemonForm resource URL to be a string")
}

func TestGetPokemonHabitat(t *testing.T) {
	rById, _ := pokemon.GetPokemonHabitat("1")
	rByName, _ := pokemon.GetPokemonHabitat("cave")
	_, err := pokemon.GetPokemonHabitat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonHabitat resource")
	assert.Equal(t, "cave", rByName.Name, "Unexpected Name for PokemonHabitat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokemonHabitat{}, rById, "Expected PokemonHabitat instance to be returned")
}

func TestGetPokemonHabitatList(t *testing.T) {
	rList, _ := pokemon.GetPokemonHabitatList(20, 0)
	rPage, _ := pokemon.GetPokemonHabitatList(1, 1)
	assert.Equal(t, "cave", rList.Results[0].Name, "Unexpected Name for PokemonHabitat resource")
	assert.Equal(t, "forest", rPage.Results[0].Name, "Unexpected Name for PokemonHabitat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonHabitatURL(t *testing.T) {
	pokemonHabitatURL := pokemon.GetPokemonHabitatURL()
	assert.Equal(t, url+PokemonHabitatEndpoint, pokemonHabitatURL, "Unexpected PokemonHabitat resource URL")
	assert.IsType(t, "", pokemonHabitatURL, "Expected PokemonHabitat resource URL to be a string")
}

func TestGetPokemonShape(t *testing.T) {
	rById, _ := pokemon.GetPokemonShape("1")
	rByName, _ := pokemon.GetPokemonShape("ball")
	_, err := pokemon.GetPokemonShape("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonShape resource")
	assert.Equal(t, "ball", rByName.Name, "Unexpected Name for PokemonShape resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokemonShape{}, rById, "Expected PokemonShape instance to be returned")
}

func TestGetPokemonShapeList(t *testing.T) {
	rList, _ := pokemon.GetPokemonShapeList(20, 0)
	rPage, _ := pokemon.GetPokemonShapeList(1, 1)
	assert.Equal(t, "ball", rList.Results[0].Name, "Unexpected Name for PokemonShape resource")
	assert.Equal(t, "squiggle", rPage.Results[0].Name, "Unexpected Name for PokemonShape resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonShapeURL(t *testing.T) {
	pokemonShapeURL := pokemon.GetPokemonShapeURL()
	assert.Equal(t, url+PokemonShapeEndpoint, pokemonShapeURL, "Unexpected PokemonShape resource URL")
	assert.IsType(t, "", pokemonShapeURL, "Expected PokemonShape resource URL to be a string")
}

func TestGetPokemonSpecies(t *testing.T) {
	rById, _ := pokemon.GetPokemonSpecies("1")
	rByName, _ := pokemon.GetPokemonSpecies("bulbasaur")
	_, err := pokemon.GetPokemonSpecies("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonSpecies resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for PokemonSpecies resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PokemonSpecies{}, rById, "Expected PokemonSpecies instance to be returned")
}

func TestGetPokemonSpeciesList(t *testing.T) {
	rList, _ := pokemon.GetPokemonSpeciesList(20, 0)
	rPage, _ := pokemon.GetPokemonSpeciesList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for PokemonSpecies resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for PokemonSpecies resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPokemonSpeciesURL(t *testing.T) {
	pokemonSpeciesURL := pokemon.GetPokemonSpeciesURL()
	assert.Equal(t, url+PokemonSpeciesEndpoint, pokemonSpeciesURL, "Unexpected PokemonSpecies resource URL")
	assert.IsType(t, "", pokemonSpeciesURL, "Expected PokemonSpecies resource URL to be a string")
}

func TestGetStat(t *testing.T) {
	rById, _ := pokemon.GetStat("1")
	rByName, _ := pokemon.GetStat("hp")
	_, err := pokemon.GetStat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Stat resource")
	assert.Equal(t, "hp", rByName.Name, "Unexpected Name for Stat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Stat{}, rById, "Expected Stat instance to be returned")
}

func TestGetStatList(t *testing.T) {
	rList, _ := pokemon.GetStatList(20, 0)
	rPage, _ := pokemon.GetStatList(1, 1)
	assert.Equal(t, "hp", rList.Results[0].Name, "Unexpected Name for Stat resource")
	assert.Equal(t, "attack", rPage.Results[0].Name, "Unexpected Name for Stat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetStatURL(t *testing.T) {
	statURL := pokemon.GetStatURL()
	assert.Equal(t, url+StatEndpoint, statURL, "Unexpected Stat resource URL")
	assert.IsType(t, "", statURL, "Expected Stat resource URL to be a string")
}

func TestGetType(t *testing.T) {
	rById, _ := pokemon.GetType("1")
	rByName, _ := pokemon.GetType("normal")
	_, err := pokemon.GetType("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Type resource")
	assert.Equal(t, "normal", rByName.Name, "Unexpected Name for Type resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Type{}, rById, "Expected Type instance to be returned")
}

func TestGetTypeList(t *testing.T) {
	rList, _ := pokemon.GetTypeList(20, 0)
	rPage, _ := pokemon.GetTypeList(1, 1)
	assert.Equal(t, "normal", rList.Results[0].Name, "Unexpected Name for Type resource")
	assert.Equal(t, "fighting", rPage.Results[0].Name, "Unexpected Name for Type resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetTypeURL(t *testing.T) {
	typeURL := pokemon.GetTypeURL()
	assert.Equal(t, url+TypeEndpoint, typeURL, "Unexpected Type resource URL")
	assert.IsType(t, "", typeURL, "Expected Type resource URL to be a string")
}
