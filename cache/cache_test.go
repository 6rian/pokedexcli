package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddGet(t *testing.T) {
	c := NewCache(10 * time.Minute)
	key := "foo"
	val := []byte("bar")

	c.Add(key, val)
	got, ok := c.Get(key)

	assert.Equal(t, 1, len(c.Entries()), "should contain a single key")
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
