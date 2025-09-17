package main

import "sync"

type DataStore struct {
	data      map[string]string
	observers []Observer
	mu        sync.RWMutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		data:      make(map[string]string),
		observers: make([]Observer, 0),
	}
}

func (ds *DataStore) subscribe(observer Observer) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.observers = append(ds.observers, observer)
}

func (ds *DataStore) unsubscribe(observer Observer) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for i, obs := range ds.observers {
		if obs == observer {
			ds.observers = append(ds.observers[:i], ds.observers[i+1:]...)
		}
	}
}

func (ds *DataStore) notify(event string, data interface{}) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	for _, observer := range ds.observers {
		go observer.Update(event, data)
	}
}
