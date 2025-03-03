package locations

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var locations ILocations = NewLocationsGroup()

func TestGetLocation(t *testing.T) {
	rById, _ := locations.GetLocation("1")
	rByName, _ := locations.GetLocation("canalave-city")
	_, err := locations.GetLocation("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Location resource")
	assert.Equal(t, "canalave-city", rByName.Name, "Unexpected Name for Location resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetLocationList(t *testing.T) {
	rList, _ := locations.GetLocationList(models.PaginationOptions{})
	rPage, _ := locations.GetLocationList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "canalave-city", rList.Results[0].Name, "Unexpected Name for Location resource")
	assert.Equal(t, "eterna-city", rPage.Results[0].Name, "Unexpected Name for Location resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetLocationArea(t *testing.T) {
	rById, _ := locations.GetLocationArea("1")
	rByName, _ := locations.GetLocationArea("canalave-city-area")
	_, err := locations.GetLocationArea("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for LocationArea resource")
	assert.Equal(t, "canalave-city-area", rByName.Name, "Unexpected Name for LocationArea resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetLocationAreaList(t *testing.T) {
	rList, _ := locations.GetLocationAreaList(models.PaginationOptions{})
	rPage, _ := locations.GetLocationAreaList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "canalave-city-area", rList.Results[0].Name, "Unexpected Name for LocationArea resource")
	assert.Equal(t, "eterna-city-area", rPage.Results[0].Name, "Unexpected Name for LocationArea resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetPalParkArea(t *testing.T) {
	rById, _ := locations.GetPalParkArea("1")
	rByName, _ := locations.GetPalParkArea("forest")
	_, err := locations.GetPalParkArea("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PalParkArea resource")
	assert.Equal(t, "forest", rByName.Name, "Unexpected Name for PalParkArea resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPalParkAreaList(t *testing.T) {
	rList, _ := locations.GetPalParkAreaList(models.PaginationOptions{})
	rPage, _ := locations.GetPalParkAreaList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "forest", rList.Results[0].Name, "Unexpected Name for PalParkArea resource")
	assert.Equal(t, "field", rPage.Results[0].Name, "Unexpected Name for PalParkArea resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetRegion(t *testing.T) {
	rById, _ := locations.GetRegion("1")
	rByName, _ := locations.GetRegion("kanto")
	_, err := locations.GetRegion("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Region resource")
	assert.Equal(t, "kanto", rByName.Name, "Unexpected Name for Region resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetRegionList(t *testing.T) {
	rList, _ := locations.GetRegionList(models.PaginationOptions{})
	rPage, _ := locations.GetRegionList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "kanto", rList.Results[0].Name, "Unexpected Name for Region resource")
	assert.Equal(t, "johto", rPage.Results[0].Name, "Unexpected Name for Region resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
