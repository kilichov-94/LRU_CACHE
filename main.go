package main

import (
	"container/list"
)

type LRUCache interface {
	Add(key, value string) bool
	Get(key string) (value string, ok bool)
	Remove(key string) (ok bool)
}

type cacheItem struct {
	key   string
	value string
}

type lruCache struct {
	capacity int
	items    map[string]*list.Element
	list     *list.List
}

func NewLRUCache(n int) LRUCache {
	return &lruCache{
		capacity: n,
		items:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *lruCache) Add(key, value string) bool {
	if _, exists := c.items[key]; exists {
		return false
	}

	if c.list.Len() >= c.capacity {
		c.evict()
	}

	item := &cacheItem{key: key, value: value}
	element := c.list.PushFront(item)
	c.items[key] = element
	return true
}

func (c *lruCache) Get(key string) (string, bool) {
	item, exists := c.items[key]
	if !exists {
		return "", false
	}

	c.list.MoveToFront(item)
	val := item.Value.(*cacheItem)

	return val.value, true
}

func (c *lruCache) Remove(key string) bool {
	item, exists := c.items[key]
	if !exists {
		return false
	}

	c.list.Remove(item)
	
	delete(c.items, key)
	
	return true
}

func (c *lruCache) evict() {
	if c.list.Len() == 0 {
		return
	}
	
	backElement := c.list.Back()
	
	if backElement != nil {
		item := backElement.Value.(*cacheItem)
		delete(c.items, item.key)
		c.list.Remove(backElement)
	}
}
