// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package lru

import (
	"fmt"
)

// StringCache tracks a limited number of strings.
// Use Intern() to get a saved version of the string, such that
//   x := cache.Intern(s1)
//   y := cache.Intern(s2)
// Now x and y will use the same underlying memory if s1 == s2.
// We track a map into a doubly linked list, moving accessed (or recently
// added) strings to the front of the list, and using the expiry at the end of
// the list to maintain the size of the cache.
// Note that StringCache is *not* thread safe, some form of mutex is necessary
// if you want to access it from multiple threads.
type StringCache struct {
	maxSize   int
	size      int
	hitCount  int64
	missCount int64
	buf       []stringElem
	values    map[string]uint32
	root      *stringElem
}

const invalidElem = uint32((1 << 32) - 1)

// NewStringCache creates a cache for string objects that will hold no-more
// than 'size' strings.
func NewStringCache(size int) *StringCache {
	cache := &StringCache{
		maxSize: size,
	}
	if size > 1<<32 {
		// maxSize cannot be > 2**32
		panic("cannot set maxSize bigger than an unsigned 32bit integer")
	}
	cache.init()
	return cache
}

// stringElem represents a doubly linked list of elements, which allows us to
// move an entry from anywhere in the list to the 'front' of the list whenever
// it is accessed.
// For a complete double-linked list implementation, see the golang standard
// library "container/list". However, by inlining some of the functionality,
// we are able to save on storage size, and remove levels of indirection.
// The only function we need is 'MoveToFront', as we also never remove items,
// instead replacing the content when the old value is expired.
type stringElem struct {
	value      string
	prev, next uint32
}

func (sc *StringCache) init() {
	initialSize := sc.maxSize + 1
	if initialSize > 100 {
		initialSize = 101
	}
	sc.values = make(map[string]uint32, initialSize)
	sc.buf = make([]stringElem, initialSize)
	sc.size = 0
	sc.root = &sc.buf[0]
	sc.root.next = 0
	sc.root.prev = 0
}

// Len returns how many strings are currently cached
func (sc *StringCache) Len() int {
	return sc.size
}

// HitCounts is used to track how well this cache is working
type HitCounts struct {
	Hit, Miss int64
}

// HitCounts gives information about accesses to the cache. The total number of
// calls to Intern can be computed by adding Hit and Miss.
func (sc *StringCache) HitCounts() HitCounts {
	return HitCounts{
		Hit:  sc.hitCount,
		Miss: sc.missCount,
	}
}

// Validate checks invariants to make sure the double-linked list is properly
// linked, and that the values map to the correct element.
func (sc *StringCache) Validate() error {
	count := 0
	if sc.root != &sc.buf[0] {
		return fmt.Errorf("error, root=%p, not buf[0]=%p", sc.root, &sc.buf[0])
	}
	for cur := sc.root.next; cur != 0; cur = sc.buf[cur].next {
		count++
		curS := &sc.buf[cur]
		if curS.prev > uint32(len(sc.buf)) {
			return fmt.Errorf("error at %#v, the value prev %d > len(buf) %d", curS, curS.prev, len(sc.buf))
		}
		prev := &sc.buf[curS.prev]
		if prev.next != cur {
			return fmt.Errorf("error at %#v, the prev.next (%d) is not this %d", curS, prev.next, cur)
		}
		if curS.next > uint32(len(sc.buf)) {
			return fmt.Errorf("error at %#v, the value next %d > len(buf) %d", curS, curS.next, len(sc.buf))
		}
		next := &sc.buf[curS.next]
		if next.prev != cur {
			return fmt.Errorf("error at %#v, the next.prev (%d) is not this %d", curS, next.prev, cur)
		}
		v := curS.value
		if sc.values[v] != cur {
			if sc.values[v] > uint32(len(sc.buf)) {
				return fmt.Errorf("error at %q, %d %# v, the map doesn't point to cur it points to: %d (outside of buf)",
					v, cur, curS, sc.values[v])
			} else {
				return fmt.Errorf("error at %q, %d %# v, the map doesn't point to cur it points to: %d %# v",
					v, cur, curS, sc.values[v], sc.buf[sc.values[v]])
			}
		}
	}
	if count != sc.size {
		return fmt.Errorf("incorrect count, expected %d got %d", sc.size, count)
	}
	if len(sc.values) != sc.size {
		return fmt.Errorf("value map has wrong count, expected %d got %d", sc.size, len(sc.values))
	}
	return nil
}

func (sc *StringCache) realloc(nextSize int) {
	if nextSize == 0 {
		// We save 1 slot at the beginning for root, this makes 'offset = 0' an invalid value
		// which makes debugging much easier, and we need start and end pointers anyway.
		nextSize = (len(sc.buf) - 1) * 3
		if nextSize > sc.maxSize {
			nextSize = sc.maxSize
		}
		nextSize++ // reserve root = buf[0]
	}
	newBuf := make([]stringElem, nextSize)
	copy(newBuf, sc.buf)
	sc.buf = newBuf
	sc.root = &newBuf[0]
}

// Intern takes a string, and returns either the cached copy of the string, or
// caches the string and returns it back.  It also updates how recently the
// string was seen, so that strings aren't cached forever.
func (sc *StringCache) Intern(v string) string {
	if elem, ok := sc.values[v]; ok {
		sc.moveToFront(elem)
		value := sc.buf[elem].value
		sc.hitCount++
		return value
	}
	sc.missCount++
	var elem uint32
	if sc.size < sc.maxSize {
		sc.size++
		elem = uint32(sc.size)
		if sc.size >= len(sc.buf) {
			sc.realloc(0)
		}
		sc.buf[elem].value = v
	} else {
		elem = sc.root.prev
		e := &sc.buf[elem]
		delete(sc.values, e.value)
		e.value = v
	}
	sc.moveToFront(elem)
	sc.values[v] = elem
	return v
}

// Contains returns true if the string is in the cache. It does not change
// information about recently-used.
func (sc *StringCache) Contains(v string) bool {
	_, ok := sc.values[v]
	return ok
}

func (sc *StringCache) moveToFront(elem uint32) {
	if sc.root.next == elem {
		// we're already at the front
		return
	}
	e := &sc.buf[elem]
	if e.prev != 0 {
		// remove it from its current spot
		sc.buf[e.prev].next = e.next
		sc.buf[e.next].prev = e.prev
	}
	next := sc.root.next
	e.prev = 0
	e.next = next
	sc.root.next = elem
	sc.buf[next].prev = elem
}

// Prealloc allocates a maxSize buffer immediately, rather than slowly growing
// the buffer to maxSize.
func (sc *StringCache) Prealloc() {
	values := make(map[string]uint32, sc.maxSize)
	for k, v := range sc.values {
		values[k] = v
	}
	sc.values = values
	sc.realloc(sc.maxSize + 1)
}
