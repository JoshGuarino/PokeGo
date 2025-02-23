package contests

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/stretchr/testify/assert"
)

// Contests struct instance for testing
var contests IContests = Contests{}

func TestGetContestType(t *testing.T) {
	rById, _ := contests.GetContestType("1")
	rByName, _ := contests.GetContestType("cool")
	_, err := contests.GetContestType("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "cool", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestTypeList(t *testing.T) {
	rList, _ := contests.GetContestTypeList()
	// _, err := contests.GetContestTypeList()
	assert.Equal(t, "cool", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := contests.GetContestEffect("1")
	_, err := contests.GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestEffectList(t *testing.T) {
	rList, _ := contests.GetContestEffectList()
	// _, err := contests.GetContestEffectList()
	assert.Equal(t, constants.ContestEffectEndpoint+"1/", rList.Results[0].URL, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := contests.GetSuperContestEffect("1")
	_, err := contests.GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSuperContestEffectList(t *testing.T) {
	rList, _ := contests.GetSuperContestEffectList()
	// _, err := contests.GetContestEffectList()
	assert.Equal(t, constants.SuperContestEffectEndpoint+"1/", rList.Results[0].URL, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}
