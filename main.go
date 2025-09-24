package main

import "log"

type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{make(map[K]V)}
}

func (c *Cache[K, V]) Add(key K, value V) {
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return  value, ok
}

func main() {
	cache := NewCache[any, any]()

	cache.Add("", "value")
	value, ok := cache.Get("")
	if !ok {
		log.Fatal("value not found")
	}
	log.Println(value)
}
