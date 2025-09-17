package main

import "fmt"

func main() {
	datastore := NewDataStore()
	firstCache := NewCache("replica1")
	secondCache := NewCache("replica2")

	datastore.subscribe(firstCache)
	datastore.subscribe(secondCache)

	datastore.notify("data_changed", map[string]string{
		"key":   "hello",
		"value": "world",
	})

	fmt.Println(firstCache.cache)
	fmt.Println(secondCache.cache)
}
