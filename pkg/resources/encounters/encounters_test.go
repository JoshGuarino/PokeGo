package encounters

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var encounters IEncounters = Encounters{}

func TestGetEncounterMethod(t *testing.T) {
	rById, _ := encounters.GetEncounterMethod("1")
	rByName, _ := encounters.GetEncounterMethod("walk")
	_, err := encounters.GetEncounterMethod("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "walk", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterMethodList(t *testing.T) {
	rList, _ := encounters.GetEncounterMethodList(models.PaginationOptions{})
	// _, err := encounters.GetEncounterMethodList()
	assert.Equal(t, "walk", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterCondition(t *testing.T) {
	rById, _ := encounters.GetEncounterCondition("1")
	rByName, _ := encounters.GetEncounterCondition("swarm")
	_, err := encounters.GetEncounterCondition("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "swarm", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionList(t *testing.T) {
	rList, _ := encounters.GetEncounterConditionList(models.PaginationOptions{})
	// _, err := encounters.GetEncounterConditionList()
	assert.Equal(t, "swarm", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionValue(t *testing.T) {
	rById, _ := encounters.GetEncounterConditionValue("1")
	rByName, _ := encounters.GetEncounterConditionValue("swarm-yes")
	_, err := encounters.GetEncounterConditionValue("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "swarm-yes", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionValueList(t *testing.T) {
	rList, _ := encounters.GetEncounterConditionValueList(models.PaginationOptions{})
	// _, err := encounters.GetEncounterConditionList()
	assert.Equal(t, "swarm-yes", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}
