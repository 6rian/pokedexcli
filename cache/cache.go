package cache

import (
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type EntriesMap map[string]CacheEntry

type Cache struct {
	entries EntriesMap
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: EntriesMap{},
	}
	c.ReapLoop(interval)
	return c
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			for key := range c.entries {
				if time.Since(c.entries[key].createdAt) > interval {
					delete(c.entries, key)
				}
			}
		}
	}()
}

func (c Cache) Add(key string, val []byte) {
	if _, exists := c.Get(key); !exists {
		c.entries[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (c Cache) Entries() EntriesMap {
	return c.entries
}
