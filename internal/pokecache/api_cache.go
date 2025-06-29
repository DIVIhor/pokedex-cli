package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (rawData []byte, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.data[key]
	if exists {
		rawData = entry.val
	}
	return
}

func (c *Cache) reapLoop(interval time.Duration) {
	timer := time.NewTicker(interval)
	defer timer.Stop()

	// iterate through the cache to check timestamps
	// and delete old values
	for {
		<-timer.C
		if len(c.data) > 0 {
			now := time.Now()
			c.mu.Lock()
			for k, v := range c.data {
				if now.Sub(v.createdAt) >= interval {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	// cache initialization
	// if it's not initialized -> causes panic
	cache := &Cache{
		data: map[string]cacheEntry{},
		mu:   &sync.RWMutex{},
	}
	go cache.reapLoop(interval)

	return cache
}
