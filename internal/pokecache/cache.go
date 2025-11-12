package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.cacheMap[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cutoff := time.Now().Add(-cache.interval)
		cache.mu.Lock()
		for key, t := range cache.cacheMap {
			if t.createdAt.Before(cutoff) {
				delete(cache.cacheMap, key)
			}
		}
		cache.mu.Unlock()
	}
}
