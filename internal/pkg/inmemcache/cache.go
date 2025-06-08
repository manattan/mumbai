package inmemcache

import (
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	Clear()
}

type item struct {
	value      interface{}
	expiration int64
}

type cache struct {
	items map[string]*item
	mu    sync.RWMutex
}

func New() Cache {
	c := &cache{
		items: make(map[string]*item),
	}
	
	go c.cleanup()
	return c
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	expiration := int64(0)
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}
	
	c.items[key] = &item{
		value:      value,
		expiration: expiration,
	}
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	
	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		return nil, false
	}
	
	return item.value, true
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.items, key)
}

func (c *cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.items = make(map[string]*item)
}

func (c *cache) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now().UnixNano()
			for key, item := range c.items {
				if item.expiration > 0 && now > item.expiration {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}
}