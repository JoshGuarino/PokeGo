package pokemon

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/stretchr/testify/assert"
)

var pokemon IPokemon = NewPokemonGroup()

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
}

func TestGetAbilityList(t *testing.T) {
	rList, _ := pokemon.GetAbilityList(20, 0)
	rPage, _ := pokemon.GetAbilityList(1, 1)
	assert.Equal(t, "stench", rList.Results[0].Name, "Unexpected Name for Ability resource")
	assert.Equal(t, "drizzle", rPage.Results[0].Name, "Unexpected Name for Ability resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetCharacteristic(t *testing.T) {
	rById, _ := pokemon.GetCharacteristic("1")
	_, err := pokemon.GetCharacteristic("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Characteristic resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetCharacteristicList(t *testing.T) {
	rList, _ := pokemon.GetCharacteristicList(20, 0)
	rPage, _ := pokemon.GetCharacteristicList(1, 1)
	assert.Equal(t, constants.CharacteristicEndpoint+"1/", rList.Results[0].URL, "Unexpected URL for Characteristic resource")
	assert.Equal(t, constants.CharacteristicEndpoint+"2/", rPage.Results[0].URL, "Unexpected URL for Characteristic resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetEggGroup(t *testing.T) {
	rById, _ := pokemon.GetEggGroup("1")
	rByName, _ := pokemon.GetEggGroup("monster")
	_, err := pokemon.GetEggGroup("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EggGroup resource")
	assert.Equal(t, "monster", rByName.Name, "Unexpected Name for EggGroup resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEggGroupList(t *testing.T) {
	rList, _ := pokemon.GetEggGroupList(20, 0)
	rPage, _ := pokemon.GetEggGroupList(1, 1)
	assert.Equal(t, "monster", rList.Results[0].Name, "Unexpected Name for EggGroup resource")
	assert.Equal(t, "water1", rPage.Results[0].Name, "Unexpected Name for EggGroup resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetGender(t *testing.T) {
	rById, _ := pokemon.GetGender("1")
	rByName, _ := pokemon.GetGender("female")
	_, err := pokemon.GetGender("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Gender resource")
	assert.Equal(t, "female", rByName.Name, "Unexpected Name for Gender resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetGenderList(t *testing.T) {
	rList, _ := pokemon.GetGenderList(20, 0)
	rPage, _ := pokemon.GetGenderList(1, 1)
	assert.Equal(t, "female", rList.Results[0].Name, "Unexpected Name for Gender resource")
	assert.Equal(t, "male", rPage.Results[0].Name, "Unexpected Name for Gender resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetGrowthRate(t *testing.T) {
	rById, _ := pokemon.GetGrowthRate("1")
	rByName, _ := pokemon.GetGrowthRate("slow")
	_, err := pokemon.GetGrowthRate("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for GrowthRate resource")
	assert.Equal(t, "slow", rByName.Name, "Unexpected Name for GrowthRate resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetGrowthRateList(t *testing.T) {
	rList, _ := pokemon.GetGrowthRateList(20, 0)
	rPage, _ := pokemon.GetGrowthRateList(1, 1)
	assert.Equal(t, "slow", rList.Results[0].Name, "Unexpected Name for GrowthRate resource")
	assert.Equal(t, "medium", rPage.Results[0].Name, "Unexpected Name for GrowthRate resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetNature(t *testing.T) {
	rById, _ := pokemon.GetNature("1")
	rByName, _ := pokemon.GetNature("hardy")
	_, err := pokemon.GetNature("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Nature resource")
	assert.Equal(t, "hardy", rByName.Name, "Unexpected Name for Nature resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetNatureList(t *testing.T) {
	rList, _ := pokemon.GetNatureList(20, 0)
	rPage, _ := pokemon.GetNatureList(1, 1)
	assert.Equal(t, "hardy", rList.Results[0].Name, "Unexpected Name for Nature resource")
	assert.Equal(t, "bold", rPage.Results[0].Name, "Unexpected Name for Nature resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokeathlonStat(t *testing.T) {
	rById, _ := pokemon.GetPokeathlonStat("1")
	rByName, _ := pokemon.GetPokeathlonStat("speed")
	_, err := pokemon.GetPokeathlonStat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokeathlonStat resource")
	assert.Equal(t, "speed", rByName.Name, "Unexpected Name for PokeathlonStat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokeathlonStatList(t *testing.T) {
	rList, _ := pokemon.GetPokeathlonStatList(20, 0)
	rPage, _ := pokemon.GetPokeathlonStatList(1, 1)
	assert.Equal(t, "speed", rList.Results[0].Name, "Unexpected Name for PokeathlonStat resource")
	assert.Equal(t, "power", rPage.Results[0].Name, "Unexpected Name for PokeathlonStat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemon(t *testing.T) {
	rById, _ := pokemon.GetPokemon("1")
	rByName, _ := pokemon.GetPokemon("bulbasaur")
	_, err := pokemon.GetPokemon("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Pokemon resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for Pokemon resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonList(t *testing.T) {
	rList, _ := pokemon.GetPokemonList(20, 0)
	rPage, _ := pokemon.GetPokemonList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for Pokemon resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for Pokemon resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemonColor(t *testing.T) {
	rById, _ := pokemon.GetPokemonColor("1")
	rByName, _ := pokemon.GetPokemonColor("black")
	_, err := pokemon.GetPokemonColor("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonColor resource")
	assert.Equal(t, "black", rByName.Name, "Unexpected Name for PokemonColor resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonColorList(t *testing.T) {
	rList, _ := pokemon.GetPokemonColorList(20, 0)
	rPage, _ := pokemon.GetPokemonColorList(1, 1)
	assert.Equal(t, "black", rList.Results[0].Name, "Unexpected Name for PokemonColor resource")
	assert.Equal(t, "blue", rPage.Results[0].Name, "Unexpected Name for PokemonColor resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemonForm(t *testing.T) {
	rById, _ := pokemon.GetPokemonForm("1")
	rByName, _ := pokemon.GetPokemonForm("bulbasaur")
	_, err := pokemon.GetPokemonForm("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonForm resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for PokemonForm resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonFormList(t *testing.T) {
	rList, _ := pokemon.GetPokemonFormList(20, 0)
	rPage, _ := pokemon.GetPokemonFormList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for PokemonForm resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for PokemonForm resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemonHabitat(t *testing.T) {
	rById, _ := pokemon.GetPokemonHabitat("1")
	rByName, _ := pokemon.GetPokemonHabitat("cave")
	_, err := pokemon.GetPokemonHabitat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonHabitat resource")
	assert.Equal(t, "cave", rByName.Name, "Unexpected Name for PokemonHabitat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonHabitatList(t *testing.T) {
	rList, _ := pokemon.GetPokemonHabitatList(20, 0)
	rPage, _ := pokemon.GetPokemonHabitatList(1, 1)
	assert.Equal(t, "cave", rList.Results[0].Name, "Unexpected Name for PokemonHabitat resource")
	assert.Equal(t, "forest", rPage.Results[0].Name, "Unexpected Name for PokemonHabitat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemonShape(t *testing.T) {
	rById, _ := pokemon.GetPokemonShape("1")
	rByName, _ := pokemon.GetPokemonShape("ball")
	_, err := pokemon.GetPokemonShape("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonShape resource")
	assert.Equal(t, "ball", rByName.Name, "Unexpected Name for PokemonShape resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonShapeList(t *testing.T) {
	rList, _ := pokemon.GetPokemonShapeList(20, 0)
	rPage, _ := pokemon.GetPokemonShapeList(1, 1)
	assert.Equal(t, "ball", rList.Results[0].Name, "Unexpected Name for PokemonShape resource")
	assert.Equal(t, "squiggle", rPage.Results[0].Name, "Unexpected Name for PokemonShape resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokemonSpecies(t *testing.T) {
	rById, _ := pokemon.GetPokemonSpecies("1")
	rByName, _ := pokemon.GetPokemonSpecies("bulbasaur")
	_, err := pokemon.GetPokemonSpecies("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PokemonSpecies resource")
	assert.Equal(t, "bulbasaur", rByName.Name, "Unexpected Name for PokemonSpecies resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokemonSpeciesList(t *testing.T) {
	rList, _ := pokemon.GetPokemonSpeciesList(20, 0)
	rPage, _ := pokemon.GetPokemonSpeciesList(1, 1)
	assert.Equal(t, "bulbasaur", rList.Results[0].Name, "Unexpected Name for PokemonSpecies resource")
	assert.Equal(t, "ivysaur", rPage.Results[0].Name, "Unexpected Name for PokemonSpecies resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetStat(t *testing.T) {
	rById, _ := pokemon.GetStat("1")
	rByName, _ := pokemon.GetStat("hp")
	_, err := pokemon.GetStat("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Stat resource")
	assert.Equal(t, "hp", rByName.Name, "Unexpected Name for Stat resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetStatList(t *testing.T) {
	rList, _ := pokemon.GetStatList(20, 0)
	rPage, _ := pokemon.GetStatList(1, 1)
	assert.Equal(t, "hp", rList.Results[0].Name, "Unexpected Name for Stat resource")
	assert.Equal(t, "attack", rPage.Results[0].Name, "Unexpected Name for Stat resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetType(t *testing.T) {
	rById, _ := pokemon.GetType("1")
	rByName, _ := pokemon.GetType("normal")
	_, err := pokemon.GetType("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Type resource")
	assert.Equal(t, "normal", rByName.Name, "Unexpected Name for Type resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetTypeList(t *testing.T) {
	rList, _ := pokemon.GetTypeList(20, 0)
	rPage, _ := pokemon.GetTypeList(1, 1)
	assert.Equal(t, "normal", rList.Results[0].Name, "Unexpected Name for Type resource")
	assert.Equal(t, "fighting", rPage.Results[0].Name, "Unexpected Name for Type resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
