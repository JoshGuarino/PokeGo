package cache

import (
	"sync"
	"time"

	"github.com/JoshGuarino/PokeGo/internal/logger"
)

// Cache interface
type ICache interface {
	Get(key string) (any, bool)
	Set(key string, value any)
	Delete(key string)
	Clear()
	SetExpiration(expiration time.Duration)
	Expiration() time.Duration
	SetActive(active bool)
	Active() bool
}

// Cache struct
type Cache struct {
	store    map[string]data
	settings settings
	lock     sync.Mutex
}

// Data struct
type data struct {
	value      any
	expiration time.Time
}

// Cache settings
type settings struct {
	expiration time.Duration
	active     bool
}

// Cache global variable
var CACHE *Cache

// Logger reference variable
var log logger.ILogger = logger.LOG

// Initialize cache
func init() {
	CACHE = newCache()
	log.Info("Cache initialized")
}

// Return an instance of Cache struct
func newCache() *Cache {
	return &Cache{
		store: make(map[string]data),
		settings: settings{
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
	c.store[key] = data{
		value:      value,
		expiration: expirationTime,
	}
}

// Get a value from the cache by key
func (c *Cache) Get(key string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.store[key]
	if !ok || time.Now().After(value.expiration) {
		delete(c.store, key)
		return nil, false
	}

	return value.value, true
}

// Delete a key-value pair from the Cache
func (c *Cache) Delete(key string) {
	log.Warn("Cached resource deleted", "key", key)
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.store, key)
}

// Clear the cache
func (c *Cache) Clear() {
	log.Warn("Cache cleared")
	c.lock.Lock()
	defer c.lock.Unlock()

	c.store = make(map[string]data)
}

// Set default expiration time for Cache
func (c *Cache) SetExpiration(expiration time.Duration) {
	log.Warn("Cache expiration time changed", "expiration", expiration)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.settings.expiration = expiration
}

// Get expiration time of Cache
func (c *Cache) Expiration() time.Duration {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.settings.expiration
}

// Set active status of cache
func (c *Cache) SetActive(active bool) {
	log.Warn("Cache active status changed", "active", active)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.settings.active = active
}

// Get active status of Cache
func (c *Cache) Active() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.settings.active
}
