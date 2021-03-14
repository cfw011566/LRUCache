package LRUCache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1) // cache is {1=1}
	fmt.Println(lruCache)
	lruCache.Put(2, 2) // cache is {1=1, 2=2}
	fmt.Println(lruCache)
	if lruCache.Get(1) != 1 {
		t.Errorf("Get(1) Expect 1")
	}
	lruCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lruCache)
	// returns -1 (not found)
	if lruCache.Get(2) != -1 {
		t.Errorf("Get(2) Expect -1")
	}
	lruCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lruCache)
	// return -1 (not found)
	if lruCache.Get(1) != -1 {
		t.Errorf("Get(1) Expect -1")
	}
	fmt.Println(lruCache)
	// return 3
	if lruCache.Get(3) != 3 {
		t.Errorf("Get(3) Expect 3")
	}
	fmt.Println(lruCache)
	// return 4
	if lruCache.Get(4) != 4 {
		t.Errorf("Get(4) Expect 4")
	}
	fmt.Println(lruCache)
}
