package contests

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

var contests IContests = NewContestsGroup()

func TestNewContestsGroup(t *testing.T) {
	contests := NewContestsGroup()
	assert.IsType(t, Contests{}, contests, "Expected Contests instance to be returned")
}

func TestGetContestType(t *testing.T) {
	rById, _ := contests.GetContestType("1")
	rByName, _ := contests.GetContestType("cool")
	_, err := contests.GetContestType("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ContestType resource")
	assert.Equal(t, "cool", rByName.Name, "Unexpected Name for ContestType resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestTypeList(t *testing.T) {
	rList, _ := contests.GetContestTypeList(20, 0)
	rPage, _ := contests.GetContestTypeList(1, 1)
	assert.Equal(t, "cool", rList.Results[0].Name, "Unexpected Name for ContestType resource")
	assert.Equal(t, "beauty", rPage.Results[0].Name, "Unexpected Name for ContestType resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := contests.GetContestEffect("1")
	_, err := contests.GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ContestEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestEffectList(t *testing.T) {
	rList, _ := contests.GetContestEffectList(20, 0)
	rPage, _ := contests.GetContestEffectList(1, 1)
	assert.Equal(t, endpoints.ContestEffect+"1/", rList.Results[0].URL, "Unexpected URL for ContestEffect resource")
	assert.Equal(t, endpoints.ContestEffect+"2/", rPage.Results[0].URL, "Unexpected URL for ContestEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := contests.GetSuperContestEffect("1")
	_, err := contests.GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for SuperContestEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSuperContestEffectList(t *testing.T) {
	rList, _ := contests.GetSuperContestEffectList(20, 0)
	rPage, _ := contests.GetSuperContestEffectList(1, 1)
	assert.Equal(t, endpoints.SuperContestEffect+"1/", rList.Results[0].URL, "Unexpected URL for SuperContestEffect resource")
	assert.Equal(t, endpoints.SuperContestEffect+"2/", rPage.Results[0].URL, "Unexpected URL for SuperContestEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
