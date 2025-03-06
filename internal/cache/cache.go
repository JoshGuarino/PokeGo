package cache

import (
	"sync"
	"time"
)

// Cache interface
type ICache interface {
	Get(key string) (any, bool)
	Set(key string, value any, expiration time.Duration)
	Clear()
}

// Cache struct
type Cache struct {
	data map[string]Value
	lock sync.Mutex
}

// Value struct
type Value struct {
	value      any
	expiration time.Time
}

// Return an instance of Cache struct
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Value),
	}
}

// Set a key-value pair in the cache
func (c *Cache) Set(key string, value any, expiration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	expirationTime := time.Now().Add(expiration)
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

// Clear the cache
func (c *Cache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data = make(map[string]Value)
}
