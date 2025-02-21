package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var berries IBerries = Berries{}

func TestGetBerry(t *testing.T) {
	rById, _ := berries.GetBerry("1")
	rByName, _ := berries.GetBerry("cheri")
	_, err := berries.GetBerry("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "cheri", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryList(t *testing.T) {
	rList, _ := berries.GetBerryList()
	// _, err := berries.GetBerryList()
	assert.Equal(t, "cheri", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := berries.GetBerryFirmness("1")
	rByName, _ := berries.GetBerryFirmness("very-soft")
	_, err := berries.GetBerryFirmness("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "very-soft", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	rList, _ := berries.GetBerryFirmnessList()
	// _, err := berries.GetBerryFirmnessList()
	assert.Equal(t, "very-soft", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := berries.GetBerryFlavor("1")
	rByName, _ := berries.GetBerryFlavor("spicy")
	_, err := berries.GetBerryFlavor("test")
	assert.Equal(t, 1, rById.ID, "")
	assert.Equal(t, "spicy", rByName.Name, "")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFlavorList(t *testing.T) {
	rList, _ := berries.GetBerryFlavorList()
	// _, err := berries.GetBerryFlavorList()
	assert.Equal(t, "spicy", rList.Results[0].Name, "")
	// assert.Error(t, err, "Expected an error to be thrown.")
}
