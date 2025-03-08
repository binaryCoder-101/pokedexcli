package pokecache

import (
	"sync"
	"time"
)

// Initialize a new cache
func NewCache(timeInterval time.Duration) *Cache {
	cache := &Cache{
		cacheBody: make(map[string]cacheEntry),
		mu:        &sync.Mutex{},
	}

	go cache.reapLoop(timeInterval)
	return cache
}

// Add new element to cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheBody[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

// Get value from cache for a particular key
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, exists := c.cacheBody[key]

	return data.val, exists
}

// Delete entries from cache after specified timer interval
func (c *Cache) reapLoop(timeInterval time.Duration) {

	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()

	timeAt := time.Now().UTC().Add(-timeInterval)

	for range ticker.C {
		c.mu.Lock()
		for key, cachebody := range c.cacheBody {
			if cachebody.createdAt.Before(timeAt) {
				delete(c.cacheBody, key)
			}
		}
		c.mu.Unlock()
	}
}
