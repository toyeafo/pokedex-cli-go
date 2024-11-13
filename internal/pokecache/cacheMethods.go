package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.readLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	if c.entry == nil {
		c.entry = make(map[string]cacheEntry)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now().Add(c.interval),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result, prs := c.entry[key]
	return result.val, prs
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, x_val := range c.entry {
			if time.Now().After(x_val.createdAt) {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
