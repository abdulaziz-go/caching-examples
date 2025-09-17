package write_back

import (
	"fmt"
	"sync"
)

type WriteBackCache struct {
	cache             map[string]string
	persistentStorage map[string]string
	dirty             map[string]bool
	mu                sync.RWMutex
}

func NewWriteBackCache() *WriteBackCache {
	return &WriteBackCache{
		cache:             make(map[string]string),
		persistentStorage: make(map[string]string),
		dirty:             make(map[string]bool),
	}
}

func (wbc *WriteBackCache) Write(key, value string) {
	wbc.mu.Lock()
	defer wbc.mu.Unlock()

	wbc.cache[key] = value
	wbc.dirty[key] = true
}

func (wbc *WriteBackCache) Read(key string) (string, bool) {
	wbc.mu.RLock()
	defer wbc.mu.RUnlock()

	if val, exist := wbc.cache[key]; exist {
		fmt.Printf("loaded from cache %v", val)
		return val, true
	}

	if val, exist := wbc.persistentStorage[key]; exist {
		fmt.Printf("loaded from peristent storage %v", val)
		return val, true
	}
	return "", false
}

func (wbc *WriteBackCache) Flush() {
	wbc.mu.Lock()
	defer wbc.mu.Unlock()

	for key := range wbc.dirty {
		if val, exist := wbc.cache[key]; exist {
			wbc.persistentStorage[key] = val
			delete(wbc.dirty, key)
		}
	}
}
