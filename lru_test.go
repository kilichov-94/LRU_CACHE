package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cache := NewLRUCache(2)

	// Add an element
	if !cache.Add("key1", "value1") {
		t.Errorf("Add() failed for key1")
	}

	// Add another element
	if !cache.Add("key2", "value2") {
		t.Errorf("Add() failed for key2")
	}

	// Adding duplicate key should return false
	if cache.Add("key1", "value3") {
		t.Errorf("Add() should return false for duplicate key")
	}

	// Adding a third element should evict the least recently used (key1)
	if !cache.Add("key3", "value3") {
		t.Errorf("Add() failed for key3")
	}

	if _, ok := cache.Get("key1"); ok {
		t.Errorf("key1 should have been evicted")
	}
}

func TestGet(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")

	// Get an existing element
	val, ok := cache.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Get() failed for key1: got %v, expected value1", val)
	}

	// Accessing key1 should make it most recently used
	cache.Add("key3", "value3") // This should evict key2
	if _, ok := cache.Get("key2"); ok {
		t.Errorf("key2 should have been evicted")
	}

	// Get a non-existing element
	if _, ok := cache.Get("key4"); ok {
		t.Errorf("Get() should return false for non-existing key")
	}
}

func TestRemove(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")

	// Remove an existing element
	if !cache.Remove("key1") {
		t.Errorf("Remove() failed for key1")
	}

	// Key1 should no longer exist
	if _, ok := cache.Get("key1"); ok {
		t.Errorf("key1 should not exist after removal")
	}

	// Remove a non-existing element
	if cache.Remove("key3") {
		t.Errorf("Remove() should return false for non-existing key")
	}
}

func TestLRUBehavior(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	// Access key1 to make it most recently used
	cache.Get("key1")

	// Adding another element should evict the least recently used (key2)
	cache.Add("key4", "value4")

	if _, ok := cache.Get("key2"); ok {
		t.Errorf("key2 should have been evicted")
	}

	// key1, key3, and key4 should still be in the cache
	if _, ok := cache.Get("key1"); !ok {
		t.Errorf("key1 should still be in the cache")
	}
	if _, ok := cache.Get("key3"); !ok {
		t.Errorf("key3 should still be in the cache")
	}
	if _, ok := cache.Get("key4"); !ok {
		t.Errorf("key4 should still be in the cache")
	}
}
