package request

import (
	"fmt"
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	body, err := Get(endpoints.BaseUrl)
	assert.IsType(t, []byte{}, body, "Expected slice of bytes to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetResourceList(t *testing.T) {
	url := endpoints.Pokemon
	key := fmt.Sprintf("%s?limit=%d&offset=%d", url, 20, 0)
	list, err := GetResourceList(url, 20, 0)
	data, _ := cache.C.Get(key)
	assert.Equal(t, list, data, "Expected resource to be cached")
	assert.IsType(t, &models.ResourceList{}, list, "Expected ResourceList instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetSpecificResource(t *testing.T) {
	url := endpoints.Pokemon + "1"
	key := url
	resource, err := GetSpecificResource[models.Pokemon](url)
	data, _ := cache.C.Get(key)
	assert.Equal(t, resource, data, "Expected resource to be cached")
	assert.IsType(t, &models.Pokemon{}, resource, "Unexpected type parameter returned")
	assert.NoError(t, err, "Expected error to nil")
}
