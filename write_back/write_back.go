package write_back

import (
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

}

func (wbc *WriteBackCache) Flush() {

}
