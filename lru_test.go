package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	if !cache.Add("a", "1") {
		t.Error("Failed to add element 'a'")
	}

	if !cache.Add("b", "2") {
		t.Error("Failed to add element 'b'")
	}

	if value, ok := cache.Get("a"); !ok || value != "1" {
		t.Error("Failed to retrieve element 'a'")
	}

	if !cache.Add("c", "3") {
		t.Error("Failed to add element 'c'")
	}

	if _, ok := cache.Get("b"); ok {
		t.Error("Element 'b' should have been evicted")
	}

	if value, ok := cache.Get("a"); !ok || value != "1" {
		t.Error("Element 'a' should still be in the cache")
	}

	if value, ok := cache.Get("c"); !ok || value != "3" {
		t.Error("Element 'c' should be in the cache")
	}
}
