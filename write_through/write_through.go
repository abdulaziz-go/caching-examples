package main

import (
	"sync"
)

type WriteThrough struct {
	cache   map[string]string
	storage map[string]string
	mu      sync.RWMutex
}

func NewWriteThroughCache() *WriteThrough {
	return &WriteThrough{
		cache:   make(map[string]string),
		storage: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}

func (c *WriteThrough) Write(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = value
	c.storage[key] = value
}

func (c *WriteThrough) Read(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, exists := c.cache[key]; exists {
		return val, true
	}

	return "", false
}
