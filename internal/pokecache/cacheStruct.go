package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheBody map[string]cacheEntry
	mu        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeInterval time.Duration) *Cache {
	cache := &Cache{
		cacheBody: make(map[string]cacheEntry),
	}

	cache.reapLoop()
	return cache
}
