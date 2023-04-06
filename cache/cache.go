package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type EntriesMap map[string]CacheEntry

type Cache struct {
	entries EntriesMap
	mux     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: EntriesMap{},
		mux:     &sync.Mutex{},
	}
	c.ReapLoop(interval)
	return c
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			c.mux.Lock()
			for key := range c.entries {
				if time.Since(c.entries[key].createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mux.Unlock()
		}
	}()
}

func (c Cache) Add(key string, val []byte) {
	if _, exists := c.Get(key); !exists {
		c.mux.Lock()
		c.entries[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
		c.mux.Unlock()
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (c Cache) Entries() EntriesMap {
	e := EntriesMap{}
	c.mux.Lock()
	for k, v := range c.entries {
		e[k] = v
	}
	c.mux.Unlock()
	return e
}
