package berries

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var berries IBerries = NewBerriesGroup()

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
}

func TestGetBerryList(t *testing.T) {
	rList, _ := berries.GetBerryList(models.PaginationOptions{})
	rPage, _ := berries.GetBerryList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "cheri", rList.Results[0].Name, "Unexpected Name for Berry resource")
	assert.Equal(t, "chesto", rPage.Results[0].Name, "Unexpected Name for Berry resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := berries.GetBerryFirmness("1")
	rByName, _ := berries.GetBerryFirmness("very-soft")
	_, err := berries.GetBerryFirmness("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for BerryFirmness resource")
	assert.Equal(t, "very-soft", rByName.Name, "Unexpected Name for BerryFirmness resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	rList, _ := berries.GetBerryFirmnessList(models.PaginationOptions{})
	rPage, _ := berries.GetBerryFirmnessList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "very-soft", rList.Results[0].Name, "Unexpected Name for BerryFirmness resource")
	assert.Equal(t, "soft", rPage.Results[0].Name, "Unexpected Name for BerryFirmness resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := berries.GetBerryFlavor("1")
	rByName, _ := berries.GetBerryFlavor("spicy")
	_, err := berries.GetBerryFlavor("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for BerryFlavor resource")
	assert.Equal(t, "spicy", rByName.Name, "Unexpected Name for BerryFlavor resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFlavorList(t *testing.T) {
	rList, _ := berries.GetBerryFlavorList(models.PaginationOptions{})
	rPage, _ := berries.GetBerryFlavorList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "spicy", rList.Results[0].Name, "Unexpected Name for BerryFlavor resource")
	assert.Equal(t, "dry", rPage.Results[0].Name, "Unexpected Name for BerryFlavor resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
