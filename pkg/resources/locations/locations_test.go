package locations

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var locations ILocations = NewLocationsGroup()
var url string = endpoints.BaseURL

func TestNewLocationsGroup(t *testing.T) {
	locations := NewLocationsGroup()
	assert.IsType(t, Locations{}, locations, "Expected Locations instance to be returned")
}

func TestGetLocation(t *testing.T) {
	rById, _ := locations.GetLocation("1")
	rByName, _ := locations.GetLocation("canalave-city")
	_, err := locations.GetLocation("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Location resource")
	assert.Equal(t, "canalave-city", rByName.Name, "Unexpected Name for Location resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Location{}, rById, "Expected Location resource to be returned")
}

func TestGetLocationList(t *testing.T) {
	rList, _ := locations.GetLocationList(20, 0)
	rPage, _ := locations.GetLocationList(1, 1)
	assert.Equal(t, "canalave-city", rList.Results[0].Name, "Unexpected Name for Location resource")
	assert.Equal(t, "eterna-city", rPage.Results[0].Name, "Unexpected Name for Location resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetLocationURL(t *testing.T) {
	locationURL := locations.GetLocationURL()
	assert.Equal(t, url+endpoints.Location, locationURL, "Unexpected Location resource URL")
	assert.IsType(t, "", locationURL, "Expected Location resource URL to be a string")
}

func TestGetLocationArea(t *testing.T) {
	rById, _ := locations.GetLocationArea("1")
	rByName, _ := locations.GetLocationArea("canalave-city-area")
	_, err := locations.GetLocationArea("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for LocationArea resource")
	assert.Equal(t, "canalave-city-area", rByName.Name, "Unexpected Name for LocationArea resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.LocationArea{}, rById, "Expected LocationArea resource to be returned")
}

func TestGetLocationAreaList(t *testing.T) {
	rList, _ := locations.GetLocationAreaList(20, 0)
	rPage, _ := locations.GetLocationAreaList(1, 1)
	assert.Equal(t, "canalave-city-area", rList.Results[0].Name, "Unexpected Name for LocationArea resource")
	assert.Equal(t, "eterna-city-area", rPage.Results[0].Name, "Unexpected Name for LocationArea resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetLocationAreaURL(t *testing.T) {
	locationAreaURL := locations.GetLocationAreaURL()
	assert.Equal(t, url+endpoints.LocationArea, locationAreaURL, "Unexpected LocationArea resource URL")
	assert.IsType(t, "", locationAreaURL, "Expected LocationArea resource URL to be a string")
}

func TestGetPalParkArea(t *testing.T) {
	rById, _ := locations.GetPalParkArea("1")
	rByName, _ := locations.GetPalParkArea("forest")
	_, err := locations.GetPalParkArea("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for PalParkArea resource")
	assert.Equal(t, "forest", rByName.Name, "Unexpected Name for PalParkArea resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.PalParkArea{}, rById, "Expected PalParkArea resource to be returned")
}

func TestGetPalParkAreaList(t *testing.T) {
	rList, _ := locations.GetPalParkAreaList(20, 0)
	rPage, _ := locations.GetPalParkAreaList(1, 1)
	assert.Equal(t, "forest", rList.Results[0].Name, "Unexpected Name for PalParkArea resource")
	assert.Equal(t, "field", rPage.Results[0].Name, "Unexpected Name for PalParkArea resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetPalParkAreaURL(t *testing.T) {
	palParkAreaURL := locations.GetPalParkAreaURL()
	assert.Equal(t, url+endpoints.PalParkArea, palParkAreaURL, "Unexpected PalParkArea resource URL")
	assert.IsType(t, "", palParkAreaURL, "Expected PalParkArea resource URL to be a string")
}

func TestGetRegion(t *testing.T) {
	rById, _ := locations.GetRegion("1")
	rByName, _ := locations.GetRegion("kanto")
	_, err := locations.GetRegion("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Region resource")
	assert.Equal(t, "kanto", rByName.Name, "Unexpected Name for Region resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Region{}, rById, "Expected Region resource to be returned")
}

func TestGetRegionList(t *testing.T) {
	rList, _ := locations.GetRegionList(20, 0)
	rPage, _ := locations.GetRegionList(1, 1)
	assert.Equal(t, "kanto", rList.Results[0].Name, "Unexpected Name for Region resource")
	assert.Equal(t, "johto", rPage.Results[0].Name, "Unexpected Name for Region resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetRegionURL(t *testing.T) {
	regionURL := locations.GetRegionURL()
	assert.Equal(t, url+endpoints.Region, regionURL, "Unexpected Region resource URL")
	assert.IsType(t, "", regionURL, "Expected Region resource URL to be a string")
}
