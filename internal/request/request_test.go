package request

import (
	"fmt"
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var url string = endpoints.BaseURL

func TestGet(t *testing.T) {
	body, err := Get(url)
	assert.IsType(t, []byte{}, body, "Expected slice of bytes to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetData(t *testing.T) {
	data, err := GetData[models.Root](url)
	_, found := cache.C.Get(url)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, models.Root{}, data, "Expected Root instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
	cachedData, err := GetData[models.Root](url)
	assert.NoError(t, err, "Expected error to nil")
	assert.Equal(t, data, cachedData, "Expected cached data to be returned")
}

func TestGetDataWithoutCache(t *testing.T) {
	cache.C.Clear()
	cache.C.SetActive(false)
	data, err := GetData[models.Root](url)
	_, found := cache.C.Get(url)
	assert.Equal(t, false, found, "Expected resource to not be cached")
	assert.IsType(t, models.Root{}, data, "Expected Root instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
	cache.C.SetActive(true)
}

func TestGetResourceList(t *testing.T) {
	url := endpoints.Pokemon
	key := fmt.Sprintf("%s?limit=%d&offset=%d", url, 20, 0)
	list, err := GetResourceList[models.NamedResourceList](url, 20, 0)
	_, found := cache.C.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, &models.NamedResourceList{}, list, "Expected ResourceList instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetSpecificResource(t *testing.T) {
	url := endpoints.Pokemon + "1"
	key := url
	resource, err := GetResource[models.Pokemon](url)
	_, found := cache.C.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, &models.Pokemon{}, resource, "Unexpected type parameter returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetResourceSlice(t *testing.T) {
	url := endpoints.Pokemon + "1" + "/encounters"
	key := url
	resourceSlice, err := GetResourceSlice[models.LocationAreaEncounter](url)
	_, found := cache.C.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, []*models.LocationAreaEncounter{}, resourceSlice, "Unexpected type parameter returned")
	assert.NoError(t, err, "Expected error to nil")
}
