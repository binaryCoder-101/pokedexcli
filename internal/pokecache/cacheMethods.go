package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheBody[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, exists := c.cacheBody[key]

	if exists {
		return data.val, true
	} else {
		return []byte{}, false
	}
}

func (c *Cache) reapLoop() {

}
