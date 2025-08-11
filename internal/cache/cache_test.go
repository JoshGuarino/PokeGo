package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var cache = NewCache()

func TestNewCache(t *testing.T) {
	assert.IsType(t, &Cache{}, cache)
}

func TestSet(t *testing.T) {
	cache.Set("key", "value")
	assert.Equal(t, "value", cache.data["key"].value, "Unexpected value in cache")
}

func TestGet(t *testing.T) {
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	assert.Equal(t, "value", value, "Unexpected value in cache")
	assert.True(t, ok)
}

func TestClear(t *testing.T) {
	cache.Set("key", "value")
	cache.Clear()
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.False(t, ok)
}

func TestDelete(t *testing.T) {
	cache.Set("key", "value")
	cache.Delete("key")
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.False(t, ok)
}

func TestSetExpiration(t *testing.T) {
	cache.SetExpiration(1 * time.Second)
	expiration := cache.Expiration()
	cache.Set("key", "value")
	time.Sleep(2 * time.Second)
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.Equal(t, 1*time.Second, expiration, "Unexpected expiration time")
	assert.False(t, ok)
}

func TestGetActive(t *testing.T) {
	assert.True(t, cache.Active(), "Expected initial active value to be true")
}

func TestSetActive(t *testing.T) {
	cache.SetActive(false)
	assert.False(t, cache.settings.active)
}
