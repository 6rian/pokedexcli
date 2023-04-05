package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_Add(t *testing.T) {
	c := NewCache(10 * time.Minute)
	key := "foo"
	val := []byte("bar")

	c.Add(key, val)

	assert.Equal(t, 1, len(c), "should contain a single key")
	assert.Contains(t, c, key, "should contain the key")
	assert.NotContains(t, c, "asdf", "should not contain the key")
	assert.Equal(t, val, c[key].val, "should contain an entry with the value")
}

func TestCache_Get(t *testing.T) {
	c := NewCache(10 * time.Minute)
	key := "testkey"
	val := []byte("testvalue")

	c.Add(key, val)

	got, ok := c.Get(key)
	assert.True(t, ok, "should be true")
	assert.Equal(t, val, got, "should match the value")
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
