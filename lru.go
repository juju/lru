// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

// LRU implements a least-recently-used cache, which tracks what items have been accessed,
// and evicts them when adding a new item if it hasn't been used recently.
package lru

import (
	"fmt"
)

// maxLRUSize is the largest we can fit in a buffer with a 32-bit unsigned offset
const maxLRUSize = (1<<32 - 1)

// LRU implements a least-recently-used cache, evicting items from the cache if they have not been accessed in a while.
type LRU struct {
	size     int
	maxSize  int
	buf      []cacheEntry
	elements map[interface{}]uint32
	root     *cacheEntry
}

type cacheEntry struct {
	prev, next uint32
	key        interface{}
	value      interface{}
}

// Create a new LRU cache that will hold no more than the given number of items,
func New(size int) *LRU {
	if size > maxLRUSize || size <= 0 {
		panic("size must not be <= 0 or >= 2^32")
	}
	initialBufSize := size + 1
	if initialBufSize > 100 {
		initialBufSize = 101
	}
	lru := &LRU{
		size:    0,
		maxSize: size,
		// TODO: dynamically allocate the size array
		buf:      make([]cacheEntry, initialBufSize),
		elements: make(map[interface{}]uint32, initialBufSize),
	}
	lru.root = &lru.buf[0]
	return lru
}

// Len gives the number of items in the cache
func (lru *LRU) Len() int {
	return lru.size
}

// Add a new entry into the LRU cache
func (lru *LRU) Add(key, value interface{}) {
	elem, exists := lru.elements[key]
	if exists {
		entry := &lru.buf[elem]
		lru.moveToFront(elem, entry)
		// Update the value
		entry.value = value
	} else {
		// We are adding an element, make sure there is room
		if lru.size < lru.maxSize {
			// grab the next element
			lru.size++
			elem = uint32(lru.size)
			if lru.size >= len(lru.buf) {
				lru.realloc()
			}
		} else {
			// reuse the least recently used element
			elem = lru.root.prev
			delete(lru.elements, lru.buf[elem].key)
		}
		if elem >= uint32(len(lru.buf)) {
			panic(fmt.Sprintf("element %d outside of buffer range: %d", elem, len(lru.buf)))
		}
		entry := &lru.buf[elem]
		entry.key = key
		entry.value = value
		lru.elements[key] = elem
		lru.moveToFront(elem, entry)
	}
}

// Get returns the Value associated with key, and a boolean as to whether it actually exists in the cache.
// If it does exist in the cache, then it is treated as recently accessed.
func (lru *LRU) Get(key interface{}) (interface{}, bool) {
	elem, exists := lru.elements[key]
	if exists {
		entry := &lru.buf[elem]
		lru.moveToFront(elem, entry)
		return entry.value, true
	} else {
		return nil, false
	}
}

// Peek is just like Get() except it doesn't affect if it was 'recently accessed'
func (lru *LRU) Peek(key interface{}) (interface{}, bool) {
	if elem, exists := lru.elements[key]; exists {
		return lru.buf[elem].value, true
	}
	return nil, false
}

func (lru *LRU) realloc() {
	// We save 1 slot at the beginning for root, this makes 'offset = 0' an invalid value
	// which makes debugging much easier, and we need start and end pointers anyway.
	nextSize := (len(lru.buf) - 1) * 2
	if nextSize > lru.maxSize {
		nextSize = lru.maxSize
	}
	newBuf := make([]cacheEntry, nextSize+1)
	copy(newBuf, lru.buf)
	lru.buf = newBuf
	lru.root = &newBuf[0]
	if nextSize == lru.maxSize {
		// We let the map grow using normal go growth, but when we hit maxSize,
		// we know that we won't ever hold more entries than that, so we don't
		// want to have it grow arbitrarily larger.
		elements := make(map[interface{}]uint32, nextSize)
		for k, v := range lru.elements {
			elements[k] = v
		}
		lru.elements = elements
	}
}

func (lru *LRU) moveToFront(elem uint32, entry *cacheEntry) {
	if lru.root.next == elem {
		// we're already at the front
		return
	}
	if entry.prev != 0 {
		// remove it from its current spot
		lru.buf[entry.prev].next = entry.next
		lru.buf[entry.next].prev = entry.prev
	}
	next := lru.root.next
	entry.prev = 0
	entry.next = next
	lru.root.next = elem
	lru.buf[next].prev = elem
}
