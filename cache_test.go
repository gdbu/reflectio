package reflectio

import (
	"reflect"
	"testing"
)

var cacheSink *Cache

func TestNewCache(t *testing.T) {
	cacheSink = NewCache()
}

func TestCache_Get(t *testing.T) {
	var test testStruct
	c := NewCache()
	m := c.Get(test, "reflectio")
	target := reflect.ValueOf(&test)

	if err := m.SetValueAsString(target, "int1", "13"); err != nil {
		t.Fatal(err)
	}
}
