package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mutex    *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new cache and set a time duration
func NewCache(interval time.Duration) Cache {

	cache := Cache{
		data:     make(map[string]cacheEntry),
		mutex:    &sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

// Add method to create new entry to the new cache
func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = entry
}

// Get method to read data from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, ok := c.data[key]
	return value.val, ok
}

// Reaploop method to close cache if interval too long
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, value := range c.data {
				if time.Since(value.createdAt) > c.interval {
					delete(c.data, key)
				}
			}
			c.mutex.Unlock()
		}

	}

}
