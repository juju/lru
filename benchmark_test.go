// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package lru_test

import (
	"fmt"
	"math/rand"

	gc "gopkg.in/check.v1"

	"github.com/juju/lru"
)

type BenchmarkLRUSuite struct{}

var _ = gc.Suite(&BenchmarkLRUSuite{})

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000010(c *gc.C) {
	benchAddAndEvictInt(c, 10)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000020(c *gc.C) {
	benchAddAndEvictInt(c, 20)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000050(c *gc.C) {
	benchAddAndEvictInt(c, 50)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000100(c *gc.C) {
	benchAddAndEvictInt(c, 100)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000200(c *gc.C) {
	benchAddAndEvictInt(c, 200)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0000500(c *gc.C) {
	benchAddAndEvictInt(c, 500)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0001000(c *gc.C) {
	benchAddAndEvictInt(c, 1000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0002000(c *gc.C) {
	benchAddAndEvictInt(c, 2000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0005000(c *gc.C) {
	benchAddAndEvictInt(c, 5000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0010000(c *gc.C) {
	benchAddAndEvictInt(c, 10000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0020000(c *gc.C) {
	benchAddAndEvictInt(c, 20000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0050000(c *gc.C) {
	benchAddAndEvictInt(c, 50000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0100000(c *gc.C) {
	benchAddAndEvictInt(c, 100000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0200000(c *gc.C) {
	benchAddAndEvictInt(c, 200000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt0500000(c *gc.C) {
	benchAddAndEvictInt(c, 500000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt1000000(c *gc.C) {
	benchAddAndEvictInt(c, 1000000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictInt2000000(c *gc.C) {
	benchAddAndEvictInt(c, 2000000)
}

func (*BenchmarkLRUSuite) BenchmarkIntMemSize(c *gc.C) {
	keys := make([]int, c.N)
	for i := 0; i < c.N; i++ {
		keys[i] = i + 1e7
	}
	c.ResetTimer()
	rand.Shuffle(c.N, func(i, j int) { keys[j], keys[i] = keys[i], keys[j] })
	cache := lru.New(c.N)
	for i := 0; i < c.N; i++ {
		cache.Add(keys[i], i)
	}
}

func (*BenchmarkLRUSuite) BenchmarkStrMemSize(c *gc.C) {
	keys := make([]string, c.N)
	for i := 0; i < c.N; i++ {
		keys[i] = fmt.Sprint(i + 1e7)
	}
	rand.Shuffle(c.N, func(i, j int) { keys[j], keys[i] = keys[i], keys[j] })
	c.ResetTimer()
	cache := lru.New(c.N)
	for i := 0; i < c.N; i++ {
		cache.Add(keys[i], i)
	}
}

func benchAddAndEvictInt(c *gc.C, size int) {
	keys := make([]int, c.N)
	for i := 0; i < c.N; i++ {
		keys[i] = i + 1e7
	}
	rand.Shuffle(c.N, func(i, j int) { keys[j], keys[i] = keys[i], keys[j] })
	cache := lru.New(size)
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		cache.Add(keys[i], i)
	}
	expectLen := size
	if c.N < expectLen {
		expectLen = c.N
	}
	c.Assert(cache.Len(), gc.Equals, expectLen)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000010(c *gc.C) {
	benchGet(c, 10)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000020(c *gc.C) {
	benchGet(c, 20)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000050(c *gc.C) {
	benchGet(c, 50)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000100(c *gc.C) {
	benchGet(c, 100)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000200(c *gc.C) {
	benchGet(c, 200)
}

func (*BenchmarkLRUSuite) BenchmarkGet0000500(c *gc.C) {
	benchGet(c, 500)
}

func (*BenchmarkLRUSuite) BenchmarkGet0001000(c *gc.C) {
	benchGet(c, 1000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0002000(c *gc.C) {
	benchGet(c, 2000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0005000(c *gc.C) {
	benchGet(c, 5000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0010000(c *gc.C) {
	benchGet(c, 10000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0020000(c *gc.C) {
	benchGet(c, 20000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0050000(c *gc.C) {
	benchGet(c, 50000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0100000(c *gc.C) {
	benchGet(c, 100000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0200000(c *gc.C) {
	benchGet(c, 200000)
}

func (*BenchmarkLRUSuite) BenchmarkGet0500000(c *gc.C) {
	benchGet(c, 500000)
}

func (*BenchmarkLRUSuite) BenchmarkGet1000000(c *gc.C) {
	benchGet(c, 1000000)
}

func (*BenchmarkLRUSuite) BenchmarkGet2000000(c *gc.C) {
	benchGet(c, 2000000)
}

func benchGet(c *gc.C, size int) {
	cache := lru.New(size)
	lookups := make([]int, size)
	// Fill the cache:
	for i := 0; i < size; i++ {
		cache.Add(i, i)
		lookups[i] = i
	}
	rand.Shuffle(len(lookups), func(i, j int) { lookups[i], lookups[j] = lookups[j], lookups[i] })
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		cache.Get(lookups[i%size])
	}
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000010(c *gc.C) {
	benchAddAndEvictStr(c, 10)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000020(c *gc.C) {
	benchAddAndEvictStr(c, 20)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000050(c *gc.C) {
	benchAddAndEvictStr(c, 50)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000100(c *gc.C) {
	benchAddAndEvictStr(c, 100)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000200(c *gc.C) {
	benchAddAndEvictStr(c, 200)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0000500(c *gc.C) {
	benchAddAndEvictStr(c, 500)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0001000(c *gc.C) {
	benchAddAndEvictStr(c, 1000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0002000(c *gc.C) {
	benchAddAndEvictStr(c, 2000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0005000(c *gc.C) {
	benchAddAndEvictStr(c, 5000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0010000(c *gc.C) {
	benchAddAndEvictStr(c, 10000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0020000(c *gc.C) {
	benchAddAndEvictStr(c, 20000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0050000(c *gc.C) {
	benchAddAndEvictStr(c, 50000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0100000(c *gc.C) {
	benchAddAndEvictStr(c, 100000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0200000(c *gc.C) {
	benchAddAndEvictStr(c, 200000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr0500000(c *gc.C) {
	benchAddAndEvictStr(c, 500000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr1000000(c *gc.C) {
	benchAddAndEvictStr(c, 1000000)
}

func (*BenchmarkLRUSuite) BenchmarkAddAndEvictStr2000000(c *gc.C) {
	benchAddAndEvictStr(c, 2000000)
}

func benchAddAndEvictStr(c *gc.C, size int) {
	keys := make([]string, c.N)
	for i := 0; i < c.N; i++ {
		keys[i] = fmt.Sprint(i + 1e7)
	}
	rand.Shuffle(c.N, func(i, j int) { keys[j], keys[i] = keys[i], keys[j] })
	cache := lru.New(size)
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		cache.Add(keys[i], i)
	}
	expectLen := size
	if c.N < expectLen {
		expectLen = c.N
	}
	c.Assert(cache.Len(), gc.Equals, expectLen)

}
