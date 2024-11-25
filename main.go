package main

import (
	"container/list"
	"encoding/json"
	"fmt"
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
	val := item.Value.(cacheItem)

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
		item := backElement.Value.(string)
		delete(c.items, item)
		c.list.Remove(backElement)
	}
}

func main() {

	// list.PushFront(8)
	// fmt.Println(list.PushFront(8))

	// fmt.Println(list.PushBack(1))
	// fmt.Println(list.PushBack(4))

	// cache := NewLRUCache(4)
	// fmt.Println("===============1")
	// cache.Add("a", "11")
	// fmt.Printf("cache %v\ncapacity %d\ntail %v\ntail_next %v\ntail_prev %v\nhead %v\n",
	// 	cache.cache,
	// 	cache.capacity,
	// 	cache.tail,
	// 	cache.tail.next,
	// 	cache.tail.prev,
	// 	cache.head)

	// fmt.Println("===============2")
	// cache.Add("b", "22")
	// fmt.Printf("cache %v\ncapacity %d\ntail %v\ntail_next %v\ntail_prev %v\nhead %v\n",
	// 	cache.cache,
	// 	cache.capacity,
	// 	cache.tail,
	// 	cache.tail.next,
	// 	cache.tail.prev,
	// 	cache.head)

	// fmt.Println("===============3")
	// cache.Add("c", "33")
	// fmt.Printf("cache %v\ncapacity %d\ntail %v\ntail_next %v\ntail_prev %v\nhead %v\n",
	// 	cache.cache,
	// 	cache.capacity,
	// 	cache.tail,
	// 	cache.tail.next,
	// 	cache.tail.prev,
	// 	cache.head)

	// fmt.Println("===============4")
	// cache.Add("d", "44")
	// fmt.Printf("cache %v\ncapacity %d\ntail %v\ntail_next %v\ntail_prev %v\nhead %v\n",
	// 	cache.cache,
	// 	cache.capacity,
	// 	cache.tail,
	// 	cache.tail.next,
	// 	cache.tail.prev,
	// 	cache.head)

	// fmt.Println("===============5")
	// cache.Add("e", "55")
	// fmt.Printf("cache %v\ncapacity %d\ntail %v\ntail_next %v\ntail_prev %v\nhead %v\n",
	// 	cache.cache,
	// 	cache.capacity,
	// 	cache.tail,
	// 	cache.tail.next,
	// 	cache.tail.prev,
	// 	cache.head)
}

func PrintJson(v interface{}) {
	beautifulJsonByte, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(beautifulJsonByte))
}
