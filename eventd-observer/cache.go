package main

type Cache struct {
	name  string
	cache map[string]string
}

func NewCache(name string) *Cache {
	return &Cache{
		name:  name,
		cache: make(map[string]string),
	}
}

func (c *Cache) Update(event string, data interface{}) {
	if event == "data_changed" {
		if eventData, ok := data.(map[string]string); ok {
			key := eventData["key"]
			value := eventData["value"]
			c.cache[key] = value
		}
	}
}
