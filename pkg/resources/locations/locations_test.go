package locations

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var locations ILocations = Locations{}

func TestGetLocation(t *testing.T) {
	rById, _ := locations.GetLocation("1")
	rByName, _ := locations.GetLocation("canalave-city")
	_, err := locations.GetLocation("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Location resource")
	assert.Equal(t, "canalave-city", rByName.Name, "Unexpected Name for Location resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMachineList(t *testing.T) {
	rList, _ := locations.GetLocationList(models.PaginationOptions{})
	rPage, _ := locations.GetLocationList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "canalave-city", rList.Results[0].Name, "Unexpected URL for Location resource")
	assert.Equal(t, "eterna-city", rPage.Results[0].Name, "Unexpected URL for Location resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

// func TestGetLocation(t *testing.T) {
// 	rById, _ := locations.GetLocation("1")
// 	rByName, _ := locations.GetLocation("canalave-city")
// 	_, err := locations.GetLocation("test")
// 	assert.Equal(t, 1, rById.ID, "Unexpected ID for Location resource")
// 	assert.Equal(t, "canalave-city", rByName.Name, "Unexpected Name for Location resource")
// 	assert.Error(t, err, "Expected an error to be thrown.")
// }
//
// func TestGetMachineList(t *testing.T) {
// 	rList, _ := locations.GetLocationList(models.PaginationOptions{})
// 	rPage, _ := locations.GetLocationList(models.PaginationOptions{Limit: 1, Offest: 1})
// 	assert.Equal(t, "canalave-city", rList.Results[0].Name, "Unexpected URL for Location resource")
// 	assert.Equal(t, "eterna-city", rPage.Results[0].Name, "Unexpected URL for Location resource")
// 	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
// }
