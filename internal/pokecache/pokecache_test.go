package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestCache_AddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/1/",
			val: []byte("pokeAPIID1Data"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/2/",
			val: []byte("pokeAPIID2Data"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Testing case %d", i), func(t *testing.T) {
			cache := NewCache(5)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected value to be present in cache")
			}
			if !bytes.Equal(val, c.val) {
				t.Errorf("Expected value %s, got %s", c.val, val)
			}
		})
	}
}

func TestCache_ReapLoop(t *testing.T) {
	cases := []struct {
		keys []string
		vals [][]byte
	}{
		{
			keys: []string{"https://pokeapi.co/api/v2/location-area/1/", "https://pokeapi.co/api/v2/location-area/2/"},
			vals: [][]byte{[]byte("pokeAPIID1Data"), []byte("pokeAPIID1Data")},
		},
		{
			keys: []string{"https://pokeapi.co/api/v2/location-area/1/", "https://pokeapi.co/api/v2/location-area/2/", "https://example.com", "https://exchange.com"},
			vals: [][]byte{[]byte("pokeAPIID1Data"), []byte("pokeAPIID2Data"), []byte("3"), []byte("4")},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Testing case %d", i), func(t *testing.T) {
			cache := NewCache(5 * time.Millisecond)
			for j, key := range c.keys {
				cache.Add(key, c.vals[j])
			}
			time.Sleep(2 * time.Millisecond)
			for j, key := range c.keys {
				val, ok := cache.Get(key)
				if !ok {
					t.Errorf("Expected value to be present in cache")
				}
				if !bytes.Equal(val, c.vals[j]) {
					t.Errorf("Expected value %s, got %s", c.vals[j], val)
				}
			}
			time.Sleep(8 * time.Millisecond)
			for _, key := range c.keys {
				_, ok := cache.Get(key)
				if ok {
					t.Errorf("Expected value to be deleted from cache")
				}
			}
		})
	}

}
