package cache

import (
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache map[string]CacheEntry

func NewCache(interval time.Duration) Cache {
	c := Cache{}
	c.ReapLoop(interval)
	return c
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			for key := range c {
				if time.Since(c[key].createdAt) > interval {
					delete(c, key)
				}
			}
		}
	}()
}

func (c Cache) Add(key string, val []byte) {
	if _, exists := c.Get(key); !exists {
		c[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c[key]; ok {
		return entry.val, true
	}
	return []byte{}, false
}
