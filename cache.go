package reflectio

import (
	"reflect"
	"sync"
)

// NewCache will generate a new Map cache
func NewCache() *Cache {
	var c Cache
	c.m = make(map[reflect.Type]Map, 4)
	return &c
}

// Cache represents a cache of Maps
type Cache struct {
	mux sync.RWMutex
	m   map[reflect.Type]Map
}

// Get will get a Map for a given value
func (c *Cache) Get(value interface{}, tagKey string) (m Map) {
	var ok bool
	rtype := reflect.TypeOf(value)
	c.mux.RLock()
	m, ok = c.m[rtype]
	c.mux.RUnlock()

	if ok {
		// Map was found, return!
		return
	}

	return c.create(rtype, tagKey)
}

// Get will get a Map for a given value
func (c *Cache) create(rtype reflect.Type, tagKey string) (m Map) {
	var ok bool
	c.mux.Lock()
	if m, ok = c.m[rtype]; !ok {
		// Map still does not exist, create map and associate it to lookup
		m = makeMap(rtype, tagKey)
		c.m[rtype] = m
	}
	c.mux.Unlock()
	return
}
