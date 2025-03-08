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
