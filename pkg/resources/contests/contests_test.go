package contests

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var contests IContests = NewContestsGroup()
var url string = environment.ENV.URL()

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
	assert.IsType(t, &models.ContestType{}, rById, "Expected ContestType instance to be returned")
}

func TestGetContestTypeList(t *testing.T) {
	rList, _ := contests.GetContestTypeList(20, 0)
	rPage, _ := contests.GetContestTypeList(1, 1)
	assert.Equal(t, "cool", rList.Results[0].Name, "Unexpected Name for ContestType resource")
	assert.Equal(t, "beauty", rPage.Results[0].Name, "Unexpected Name for ContestType resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetContestTypeURL(t *testing.T) {
	contestTypeURL := contests.GetContestTypeURL()
	assert.Equal(t, url+ContestTypeEndpoint, contestTypeURL, "Unexpected ContestType resource URL")
	assert.IsType(t, "", contestTypeURL, "Expected ContestType resource URL to be a string")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := contests.GetContestEffect("1")
	_, err := contests.GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ContestEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.ContestEffect{}, rById, "Expected ContestEffect instance to be returned")
}

func TestGetContestEffectList(t *testing.T) {
	rList, _ := contests.GetContestEffectList(20, 0)
	rPage, _ := contests.GetContestEffectList(1, 1)
	assert.Equal(t, contests.GetContestEffectURL()+"1/", rList.Results[0].URL, "Unexpected URL for ContestEffect resource")
	assert.Equal(t, contests.GetContestEffectURL()+"2/", rPage.Results[0].URL, "Unexpected URL for ContestEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.ResourceList{}, rList, "Expected ResourceList instance to be returned")
}

func TestGetContestEffectURL(t *testing.T) {
	contestEffectURL := contests.GetContestEffectURL()
	assert.Equal(t, url+ContestEffectEndpoint, contestEffectURL, "Unexpected ContestEffect resource URL")
	assert.IsType(t, "", contestEffectURL, "Expected ContestEffect resource URL to be a string")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := contests.GetSuperContestEffect("1")
	_, err := contests.GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for SuperContestEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.SuperContestEffect{}, rById, "Expected SuperContestEffect instance to be returned")
}

func TestGetSuperContestEffectList(t *testing.T) {
	rList, _ := contests.GetSuperContestEffectList(20, 0)
	rPage, _ := contests.GetSuperContestEffectList(1, 1)
	assert.Equal(t, contests.GetSuperContestEffectURL()+"1/", rList.Results[0].URL, "Unexpected URL for SuperContestEffect resource")
	assert.Equal(t, contests.GetSuperContestEffectURL()+"2/", rPage.Results[0].URL, "Unexpected URL for SuperContestEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.ResourceList{}, rList, "Expected ResourceList instance to be returned")
}

func TestGetSuperContestEffectURL(t *testing.T) {
	superContestEffectURL := contests.GetSuperContestEffectURL()
	assert.Equal(t, url+SuperContestEffectEndpoint, superContestEffectURL, "Unexpected SuperContestEffect resource URL")
	assert.IsType(t, "", superContestEffectURL, "Expected SuperContestEffect resource URL to be a string")
}
