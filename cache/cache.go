package cache

import "time"

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache map[string]CacheEntry

func NewCache(interval time.Duration) Cache {
	return Cache{}
}

func (c Cache) Add(key string, val []byte) {
	// TODO: don't allow empty string as key
	c[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c[key]; ok {
		return entry.val, true
	}
	return []byte{}, false
}
