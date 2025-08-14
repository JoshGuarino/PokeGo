package request

import (
	"fmt"
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var url string = environment.ENV.URL()
var endpoint string = "/pokemon/"

func TestGet(t *testing.T) {
	body, err := Get(url)
	assert.IsType(t, []byte{}, body, "Expected slice of bytes to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetData(t *testing.T) {
	data, err := GetData[models.Root](url)
	_, found := cache.CACHE.Get(url)
	assert.Equal(t, true, found, "Expected data to be cached")
	assert.IsType(t, models.Root{}, data, "Expected Root instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
	cachedData, err := GetData[models.Root](url)
	assert.NoError(t, err, "Expected error to nil")
	assert.Equal(t, data, cachedData, "Expected cached data to be returned")
}

func TestGetDataWithoutCache(t *testing.T) {
	cache.CACHE.Clear()
	cache.CACHE.SetActive(false)
	data, err := GetData[models.Root](url)
	_, found := cache.CACHE.Get(url)
	assert.Equal(t, false, found, "Expected data to not be cached")
	assert.IsType(t, models.Root{}, data, "Expected Pokemon instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
	cache.CACHE.SetActive(true)
}

func TestGetResourceList(t *testing.T) {
	key := fmt.Sprintf("%s?limit=%d&offset=%d", url+endpoint, 20, 0)
	list, err := GetResourceList[models.NamedResourceList](url+endpoint, 20, 0)
	_, found := cache.CACHE.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, &models.NamedResourceList{}, list, "Expected NamedResourceList instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetSpecificResource(t *testing.T) {
	key := url + endpoint + "1"
	resource, err := GetResource[models.Pokemon](url + endpoint + "1")
	_, found := cache.CACHE.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, &models.Pokemon{}, resource, "Expected Pokemon instance to be returned")
	assert.NoError(t, err, "Expected error to nil")
}

func TestGetResourceSlice(t *testing.T) {
	key := url + endpoint + "1" + "/encounters"
	resourceSlice, err := GetResourceSlice[models.LocationAreaEncounter](key)
	_, found := cache.CACHE.Get(key)
	assert.Equal(t, true, found, "Expected resource to be cached")
	assert.IsType(t, []*models.LocationAreaEncounter{}, resourceSlice, "Expected LocationAreaEncounter slice to be returned")
	assert.NoError(t, err, "Expected error to nil")
}
