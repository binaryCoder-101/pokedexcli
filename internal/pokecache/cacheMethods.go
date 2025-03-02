package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheBody[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, exists := c.cacheBody[key]

	return data.val, exists
}

func (c *Cache) reapLoop(timeInterval time.Duration) {
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()

	timeAt := time.Now().UTC().Add(-timeInterval)

	for range ticker.C {
		for key, cachebody := range c.cacheBody {
			if cachebody.createdAt.Before(timeAt) {
				delete(c.cacheBody, key)
			}
		}
	}
}
