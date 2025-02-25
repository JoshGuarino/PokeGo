package evolution

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var evolution IEvolution = Evolution{}

func TestNewEvolutionGroup(t *testing.T) {
	evolution := NewEvolutionGroup()
	assert.IsType(t, Evolution{}, evolution, "Expected Evolution instance to be returned")
}

func TestGetEvolutionChain(t *testing.T) {
	rById, _ := evolution.GetEvolutionChain("1")
	_, err := evolution.GetEvolutionChain("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EvolutionChain resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionChainList(t *testing.T) {
	rList, _ := evolution.GetEvolutionChainList(models.PaginationOptions{})
	rPage, _ := evolution.GetEvolutionChainList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, constants.EvolutionChainEndpoint+"1/", rList.Results[0].URL, "Unexpected URL for EvolutionChain resource")
	assert.Equal(t, constants.EvolutionChainEndpoint+"2/", rPage.Results[0].URL, "Unexpected URL for EvolutionChain resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := evolution.GetEvolutionTrigger("1")
	rByName, _ := evolution.GetEvolutionTrigger("level-up")
	_, err := evolution.GetEvolutionTrigger("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EvolutionTrigger resource")
	assert.Equal(t, "level-up", rByName.Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	rList, _ := evolution.GetEvolutionTriggerList(models.PaginationOptions{})
	rPage, _ := evolution.GetEvolutionTriggerList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "level-up", rList.Results[0].Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Equal(t, "trade", rPage.Results[0].Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
