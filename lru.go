package LRUCache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	size  int
	list  *list.List
	items map[int]*list.Element
}

type node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.size = capacity
	cache.list = list.New()
	cache.items = make(map[int]*list.Element)
	return cache
}

func (this LRUCache) String() string {
	text := fmt.Sprintf("size = %d\n", this.list.Len())
	text = text + "{"
	for e := this.list.Front(); e != nil; e = e.Next() {
		n := e.Value.(node)
		text = text + fmt.Sprintf("[%d, %d] ", n.key, n.value)
	}
	text = text + "}"
	return text
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.items[key]; ok {
		this.list.MoveToFront(elem)
		n := elem.Value.(node)
		if n.key == key {
			return n.value
		} else {
			return -1
		}
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.items[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = value
		return
	}
	n := node{key, value}
	elem := this.list.PushFront(n)
	this.items[key] = elem

	if this.list.Len() > this.size {
		elem := this.list.Back()
		n := elem.Value.(node)
		delete(this.items, n.key)
		this.list.Remove(elem)
	}
}
