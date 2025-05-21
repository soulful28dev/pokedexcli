package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		data:     make(map[string]cacheEntry),
		mu:       sync.RWMutex{},
		interval: interval,
	}
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheEntry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			c.data = map[string]cacheEntry{}
			c.mu.Unlock()
		}
	}
}
