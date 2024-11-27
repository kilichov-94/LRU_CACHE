package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	added := cache.Add("key1", "value1")
	require.True(t, added, "Expected true")

	value, found := cache.Get("key1")
	require.True(t, found, "Expected true")
	require.Equal(t, "value1", value, "Expected value1")

	added = cache.Add("key1", "newValue")
	require.False(t, added, "Expected false")
}

func Test_Evict(t *testing.T) {
	cache := NewLRUCache(2)

	require.True(t, cache.Add("key1", "value1"), "Expected true")
	require.True(t, cache.Add("key2", "value2"), "Expected true")

	require.True(t, cache.Add("key3", "value3"), "Expected true")

	_, found := cache.Get("key1")
	require.False(t, found, "Expected key1 to be evicted")

	value, found := cache.Get("key2")
	require.True(t, found, "Expected key2 to still be in cache")
	require.Equal(t, "value2", value, "Expected value of key2 to be 'value2'")

	value, found = cache.Get("key3")
	require.True(t, found, "Expected key3 to still be in cache")
	require.Equal(t, "value3", value, "Expected value of key3 to be 'value3'")
}

func Test_Remove(t *testing.T) {
	cache := NewLRUCache(2)

	require.True(t, cache.Add("key1", "value1"), "Expected true")

	removed := cache.Remove("key1")
	require.True(t, removed, "Expected true")

	_, found := cache.Get("key1")
	require.False(t, found, "Expected key1 to be removed")

	removed = cache.Remove("key1")
	require.False(t, removed, "Expected false")

}

func Test_UsageOrder(t *testing.T) {
	cache := NewLRUCache(2)

	require.True(t, cache.Add("key1", "value1"), "Expected true")
	require.True(t, cache.Add("key2", "value2"), "Expected true")

	value, found := cache.Get("key1")
	require.True(t, found, "Expected key1 to still be in cache")
	require.Equal(t, "value1", value, "Expected value of key1 to be 'value1'")

	require.True(t, cache.Add("key3", "value3"), "Expected true")

	_, found = cache.Get("key2")
	require.False(t, found, "Expected key2 to be evicted")

	value, found = cache.Get("key1")
	require.True(t, found, "Expected key1 to still be in cache")
	require.Equal(t, "value1", value, "Expected value of key1 to be 'value1'")

	value, found = cache.Get("key3")
	require.True(t, found, "Expected key3 to still be in cache")
	require.Equal(t, "value3", value, "Expected value of key3 to be 'value3'")
}
