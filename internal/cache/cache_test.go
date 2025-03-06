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
	cache.Set("key", "value", 5*time.Second)
	assert.Equal(t, "value", cache.data["key"].value, "Unexpected value in cache")
}

func TestGet(t *testing.T) {
	cache.Set("key", "value", 5*time.Second)
	value, ok := cache.Get("key")
	assert.Equal(t, "value", value, "Unexpected value in cache")
	assert.True(t, ok)
}

func TestCear(t *testing.T) {
	cache.Set("key", "value", 100*time.Second)
	cache.Clear()
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.False(t, ok)
}
