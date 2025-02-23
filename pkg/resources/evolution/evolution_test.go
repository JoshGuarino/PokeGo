package evolution

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/stretchr/testify/assert"
)

var evolution IEvolution = Evolution{}

func TestGetEvolutionChain(t *testing.T) {
	rById, _ := evolution.GetEvolutionChain("1")
	_, err := evolution.GetEvolutionChain("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionChainList(t *testing.T) {
	rList, _ := evolution.GetEvolutionChainList()
	// _, err := evolution.GetEvolutionChainList()
	assert.Equal(t, constants.EvolutionChainEndpoint+"1/", rList.Results[0].URL, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := evolution.GetEvolutionTrigger("1")
	rByName, _ := evolution.GetEvolutionTrigger("level-up")
	_, err := evolution.GetEvolutionTrigger("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "level-up", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	rList, _ := evolution.GetEvolutionTriggerList()
	// _, err := evolution.GetEvolutionTriggerList()
	assert.Equal(t, "level-up", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}
