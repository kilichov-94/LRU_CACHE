// package main

// import (
// 	"fmt"
// 	"testing"
// )

// func TestLRUCache(t *testing.T) {
// 	cache := NewLRUCache(2)
// 	fmt.Println("===============1")
// 	if !cache.Add("a", "1") {
// 		t.Error("Failed to add element 'a'")
// 	}
// 	fmt.Printf("cache %v\ncapacity %d\ntail %v\nhead %v\n",
// 		cache.cache,
// 		cache.capacity,
// 		cache.tail,
// 		cache.head)

// 	fmt.Println("===============2")
// 	if !cache.Add("b", "2") {
// 		t.Error("Failed to add element 'b'")
// 	}

// 	fmt.Printf("cache %v\ncapacity %d\ntail %v\nhead %v\n",
// 		cache.cache,
// 		cache.capacity,
// 		cache.tail,
// 		cache.head)

// 	fmt.Println("===============3")
// 	if value, ok := cache.Get("a"); !ok || value != "1" {
// 		t.Error("Failed to retrieve element 'a'")
// 	}

// 	fmt.Printf("cache %v\ncapacity %d\ntail %v\nhead %v\n",
// 		cache.cache,
// 		cache.capacity,
// 		cache.tail,
// 		cache.head)

// 	fmt.Println("===============4")
// 	if !cache.Add("c", "3") {
// 		t.Error("Failed to add element 'c'")
// 	}

// 	fmt.Printf("cache %v\ncapacity %d\ntail %v\nhead %v\n",
// 		cache.cache,
// 		cache.capacity,
// 		cache.tail,
// 		cache.head)

// 	if _, ok := cache.Get("b"); ok {
// 		t.Error("Element 'b' should have been evicted")
// 	}

// 	if value, ok := cache.Get("a"); !ok || value != "1" {
// 		t.Error("Element 'a' should still be in the cache")
// 	}

// 	if value, ok := cache.Get("c"); !ok || value != "3" {
// 		t.Error("Element 'c' should be in the cache")
// 	}
// }
