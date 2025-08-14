package cache

import (
	"sync"
	"time"

	"github.com/JoshGuarino/PokeGo/internal/logger"
)

// Cache interface
type ICache interface {
	Get(key string) (any, bool)
	Set(key string, value any, expiration time.Duration)
	Delete(key string)
	Clear()
	SetExpiration(expiration time.Duration)
	Expiration() time.Duration
	SetActive(active bool)
	Active() bool
}

// Cache struct
type Cache struct {
	data     map[string]Value
	settings Settings
	lock     sync.Mutex
}

// Value struct
type Value struct {
	value      any
	expiration time.Time
}

// Cache settings
type Settings struct {
	expiration time.Duration
	active     bool
}

// Cache global variable
var CACHE *Cache

// Initialize cache
func init() {
	CACHE = NewCache()
	logger.LOG.Info("Cache initialized")
}

// Return an instance of Cache struct
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Value),
		settings: Settings{
			expiration: 24 * time.Hour,
			active:     true,
		},
	}
}

// Set a key-value pair in the cache
func (c *Cache) Set(key string, value any) {
	c.lock.Lock()
	defer c.lock.Unlock()

	expirationTime := time.Now().Add(c.settings.expiration)
	c.data[key] = Value{
		value:      value,
		expiration: expirationTime,
	}
}

// Get a value from the cache by key
func (c *Cache) Get(key string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.data[key]
	if !ok || time.Now().After(value.expiration) {
		delete(c.data, key)
		return nil, false
	}

	return value.value, true
}

// Delete a key-value pair from the Cache
func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, key)
}

// Clear the cache
func (c *Cache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data = make(map[string]Value)
}

// Set default expiration time for Cache
func (c *Cache) SetExpiration(expiration time.Duration) {
	c.settings.expiration = expiration
}

// Get expiration time of Cache
func (c *Cache) Expiration() time.Duration {
	return c.settings.expiration
}

// Set active status of cache
func (c *Cache) SetActive(active bool) {
	c.settings.active = active
}

// Get active status of Cache
func (c *Cache) Active() bool {
	return c.settings.active
}
