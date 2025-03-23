package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var games IGames = NewGamesGroup()

func TestNewGamesGroup(t *testing.T) {
	games := NewGamesGroup()
	assert.IsType(t, Games{}, games, "Expected Games instance to be returned")
}

func TestGetGeneration(t *testing.T) {
	rById, _ := games.GetGeneration("1")
	rByName, _ := games.GetGeneration("generation-i")
	_, err := games.GetGeneration("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Generation resource")
	assert.Equal(t, "generation-i", rByName.Name, "Unexpected Name for Generation resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetGenerationList(t *testing.T) {
	rList, _ := games.GetGenerationList(20, 0)
	rPage, _ := games.GetGenerationList(1, 1)
	assert.Equal(t, "generation-i", rList.Results[0].Name, "Unexpected Name for Generation resource")
	assert.Equal(t, "generation-ii", rPage.Results[0].Name, "Unexpected Name for Generation resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPokedex(t *testing.T) {
	rById, _ := games.GetPokedex("1")
	rByName, _ := games.GetPokedex("national")
	_, err := games.GetPokedex("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Pokedex resource")
	assert.Equal(t, "national", rByName.Name, "Unexpected Name for Pokedex resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokedexList(t *testing.T) {
	rList, _ := games.GetPokedexList(20, 0)
	rPage, _ := games.GetPokedexList(1, 1)
	assert.Equal(t, "national", rList.Results[0].Name, "Unexpected Name for Pokedex resource")
	assert.Equal(t, "kanto", rPage.Results[0].Name, "Unexpected Name for Pokedex resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetVersion(t *testing.T) {
	rById, _ := games.GetVersion("1")
	rByName, _ := games.GetVersion("red")
	_, err := games.GetVersion("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Version resource")
	assert.Equal(t, "red", rByName.Name, "Unexpected Name for Version resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetVersionList(t *testing.T) {
	rList, _ := games.GetVersionList(20, 0)
	rPage, _ := games.GetVersionList(1, 1)
	assert.Equal(t, "red", rList.Results[0].Name, "Unexpected Name for Version resource")
	assert.Equal(t, "blue", rPage.Results[0].Name, "Unexpected Name for Version resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetVersionGroup(t *testing.T) {
	rById, _ := games.GetVersionGroup("1")
	rByName, _ := games.GetVersionGroup("red-blue")
	_, err := games.GetVersionGroup("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for VersionGroup resource")
	assert.Equal(t, "red-blue", rByName.Name, "Unexpected Name for VersionGroup resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetVersionListGroup(t *testing.T) {
	rList, _ := games.GetVersionGroupList(20, 0)
	rPage, _ := games.GetVersionGroupList(1, 1)
	assert.Equal(t, "red-blue", rList.Results[0].Name, "Unexpected Name for VersionGroup resource")
	assert.Equal(t, "yellow", rPage.Results[0].Name, "Unexpected Name for VersionGroup resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
