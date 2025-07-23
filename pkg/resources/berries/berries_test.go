package berries

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var berries IBerries = NewBerriesGroup()
var url string = endpoints.BaseURL

func TestNewBerriesGroup(t *testing.T) {
	berries := NewBerriesGroup()
	assert.IsType(t, Berries{}, berries, "Expected Berries instance to be returned")
}

func TestGetBerry(t *testing.T) {
	rById, _ := berries.GetBerry("1")
	rByName, _ := berries.GetBerry("cheri")
	_, err := berries.GetBerry("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Berry resource")
	assert.Equal(t, "cheri", rByName.Name, "Unexpected name for Berry resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Berry{}, rById, "Expected Berry instance to be returned")
}

func TestGetBerryList(t *testing.T) {
	rList, _ := berries.GetBerryList(20, 0)
	rPage, _ := berries.GetBerryList(1, 1)
	assert.Equal(t, "cheri", rList.Results[0].Name, "Unexpected Name for Berry resource")
	assert.Equal(t, "chesto", rPage.Results[0].Name, "Unexpected Name for Berry resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetBerryURL(t *testing.T) {
	berryURL := berries.GetBerryURL()
	assert.Equal(t, url+endpoints.Berry, berryURL, "Unexpected Berry resource URL")
	assert.IsType(t, "", berryURL, "Expected Berry resource URL to be a string")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := berries.GetBerryFirmness("1")
	rByName, _ := berries.GetBerryFirmness("very-soft")
	_, err := berries.GetBerryFirmness("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for BerryFirmness resource")
	assert.Equal(t, "very-soft", rByName.Name, "Unexpected Name for BerryFirmness resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.BerryFirmness{}, rById, "Expected BerryFirmness instance to be returned")
}

func TestGetBerryFirmnessList(t *testing.T) {
	rList, _ := berries.GetBerryFirmnessList(20, 0)
	rPage, _ := berries.GetBerryFirmnessList(1, 1)
	assert.Equal(t, "very-soft", rList.Results[0].Name, "Unexpected Name for BerryFirmness resource")
	assert.Equal(t, "soft", rPage.Results[0].Name, "Unexpected Name for BerryFirmness resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetBerryFirmnessURL(t *testing.T) {
	berryFirmnessURL := berries.GetBerryFirmnessURL()
	assert.Equal(t, url+endpoints.BerryFirmness, berryFirmnessURL, "Unexpected BerryFirmness resource URL")
	assert.IsType(t, "", berryFirmnessURL, "Expected BerryFirmness resource URL to be a string")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := berries.GetBerryFlavor("1")
	rByName, _ := berries.GetBerryFlavor("spicy")
	_, err := berries.GetBerryFlavor("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for BerryFlavor resource")
	assert.Equal(t, "spicy", rByName.Name, "Unexpected Name for BerryFlavor resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.BerryFlavor{}, rById, "Expected BerryFlavor instance to be returned")
}

func TestGetBerryFlavorList(t *testing.T) {
	rList, _ := berries.GetBerryFlavorList(20, 0)
	rPage, _ := berries.GetBerryFlavorList(1, 1)
	assert.Equal(t, "spicy", rList.Results[0].Name, "Unexpected Name for BerryFlavor resource")
	assert.Equal(t, "dry", rPage.Results[0].Name, "Unexpected Name for BerryFlavor resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetBerryFlavorURL(t *testing.T) {
	berryFlavorURL := berries.GetBerryFlavorURL()
	assert.Equal(t, url+endpoints.BerryFlavor, berryFlavorURL, "Unexpected BerryFlavor resource URL")
	assert.IsType(t, "", berryFlavorURL, "Expected BerryFlavor resource URL to be a string")
}
