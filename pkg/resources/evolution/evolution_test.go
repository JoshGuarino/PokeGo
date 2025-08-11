package evolution

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var evolution IEvolution = NewEvolutionGroup()
var url string = env.ENV.URL()

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
	rList, _ := evolution.GetEvolutionChainList(20, 0)
	rPage, _ := evolution.GetEvolutionChainList(1, 1)
	assert.Equal(t, evolution.GetEvolutionChainURL()+"1/", rList.Results[0].URL, "Unexpected URL for EvolutionChain resource")
	assert.Equal(t, evolution.GetEvolutionChainURL()+"2/", rPage.Results[0].URL, "Unexpected URL for EvolutionChain resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.ResourceList{}, rList, "Expected ResourceList instance to be returned")
}

func TestGetEvolutionChainURL(t *testing.T) {
	evolutionChainURL := evolution.GetEvolutionChainURL()
	assert.Equal(t, url+EvolutionChainEndpoint, evolutionChainURL, "Unexpected EvolutionChain resource URL")
	assert.IsType(t, "", evolutionChainURL, "Expected EvolutionChain resource URL to be a string")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := evolution.GetEvolutionTrigger("1")
	rByName, _ := evolution.GetEvolutionTrigger("level-up")
	_, err := evolution.GetEvolutionTrigger("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EvolutionTrigger resource")
	assert.Equal(t, "level-up", rByName.Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.EvolutionTrigger{}, rById, "Expected EvolutionTrigger instance to be returned")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	rList, _ := evolution.GetEvolutionTriggerList(20, 0)
	rPage, _ := evolution.GetEvolutionTriggerList(1, 1)
	assert.Equal(t, "level-up", rList.Results[0].Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Equal(t, "trade", rPage.Results[0].Name, "Unexpected Name for EvolutionTrigger resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetEvolutionTriggerURL(t *testing.T) {
	evolutionTriggerURL := evolution.GetEvolutionTriggerURL()
	assert.Equal(t, url+EvolutionTriggerEndpoint, evolutionTriggerURL, "Unexpected EvolutionTrigger resource URL")
	assert.IsType(t, "", evolutionTriggerURL, "Expected EvolutionTrigger resource URL to be a string")
}
