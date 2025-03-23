package encounters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var encounters IEncounters = NewEncountersGroup()

func TestNewEncountersGroup(t *testing.T) {
	encounters := NewEncountersGroup()
	assert.IsType(t, Encounters{}, encounters, "Expected Encounters instance to be returned")
}

func TestGetEncounterMethod(t *testing.T) {
	rById, _ := encounters.GetEncounterMethod("1")
	rByName, _ := encounters.GetEncounterMethod("walk")
	_, err := encounters.GetEncounterMethod("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EncounterMethod resource")
	assert.Equal(t, "walk", rByName.Name, "Unexpected Name for EncounterMethod resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterMethodList(t *testing.T) {
	rList, _ := encounters.GetEncounterMethodList(20, 0)
	rPage, _ := encounters.GetEncounterMethodList(1, 1)
	assert.Equal(t, "walk", rList.Results[0].Name, "Unexpected Name for EncounterMethod resource")
	assert.Equal(t, "old-rod", rPage.Results[0].Name, "Unexpected Name for EncounterMethod resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetEncounterCondition(t *testing.T) {
	rById, _ := encounters.GetEncounterCondition("1")
	rByName, _ := encounters.GetEncounterCondition("swarm")
	_, err := encounters.GetEncounterCondition("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EncounterCondition resource")
	assert.Equal(t, "swarm", rByName.Name, "Unexpected Name for EncounterCondition resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionList(t *testing.T) {
	rList, _ := encounters.GetEncounterConditionList(20, 0)
	rPage, _ := encounters.GetEncounterConditionList(1, 1)
	assert.Equal(t, "swarm", rList.Results[0].Name, "Unexpected Name for EncounterCondition resource")
	assert.Equal(t, "time", rPage.Results[0].Name, "Unexpected Name for EncounterCondition resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetEncounterConditionValue(t *testing.T) {
	rById, _ := encounters.GetEncounterConditionValue("1")
	rByName, _ := encounters.GetEncounterConditionValue("swarm-yes")
	_, err := encounters.GetEncounterConditionValue("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for EncounterConditionValue resource")
	assert.Equal(t, "swarm-yes", rByName.Name, "Unexpected Name for EncounterConditionValue resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionValueList(t *testing.T) {
	rList, _ := encounters.GetEncounterConditionValueList(20, 0)
	rPage, _ := encounters.GetEncounterConditionValueList(1, 1)
	assert.Equal(t, "swarm-yes", rList.Results[0].Name, "Unexpected Name for EncounterConditionValue resource")
	assert.Equal(t, "swarm-no", rPage.Results[0].Name, "Unexpected Name for EncounterConditionValue resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
