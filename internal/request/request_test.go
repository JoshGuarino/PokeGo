package request

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	body, err := Get(constants.BaseUrl)
	assert.IsType(t, []byte{}, body, "Expected slice of bytes to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetResourceList(t *testing.T) {
	list, err := GetResourceList(constants.PokemonEndpoint, models.PaginationOptions{})
	assert.IsType(t, &models.ResourceList{}, list, "Expected ResourceList instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetSpecificResource(t *testing.T) {
	resource, err := GetSpecificResource[models.Pokemon](constants.PokemonEndpoint + "1")
	assert.IsType(t, &models.Pokemon{}, resource, "Unexpected type parameter returned")
	assert.NoError(t, err, "Expected error to nil")
}
