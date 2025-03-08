package request

import (
	"fmt"
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
	url := constants.PokemonEndpoint
	key := fmt.Sprintf("%s?offset=%d&limit=%d", url, 0, 20)
	list, err := GetResourceList(url, models.PaginationOptions{Offest: 0, Limit: 20})
	data, _ := c.Get(key)
	assert.Equal(t, list, data, "Expected resource to be cached")
	assert.IsType(t, &models.ResourceList{}, list, "Expected ResourceList instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetSpecificResource(t *testing.T) {
	url := constants.PokemonEndpoint + "1"
	key := url
	resource, err := GetSpecificResource[models.Pokemon](url)
	data, _ := c.Get(key)
	assert.Equal(t, resource, data, "Expected resource to be cached")
	assert.IsType(t, &models.Pokemon{}, resource, "Unexpected type parameter returned")
	assert.NoError(t, err, "Expected error to nil")
}
