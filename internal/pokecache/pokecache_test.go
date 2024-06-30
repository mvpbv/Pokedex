package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("Cache not created")
	}
}
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{{
		inputKey: "key1",
		inputVal: []byte("val1"),
	}, {
		inputKey: "key2",
		inputVal: []byte("val2"),
	}, {
		inputKey: "",
		inputVal: []byte("val3"),
	},
	}
	for _, cas := range cases {
		cache.Add("key1", []byte("val1"))
		actual, ok := cache.Get("key1")
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != "val1" {
			t.Errorf("Value doesn't match %s", string(actual))
		}
	}

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("Value not found")
	}
	if string(actual) != "val1" {
		t.Error("Value doens't match")
	}
}
func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(time.Millisecond)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
