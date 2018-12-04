// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package lru_test

import (
	"math/rand"

	gc "gopkg.in/check.v1"

	"github.com/juju/lru"
)

type BenchmarkSuite struct{}

var _ = gc.Suite(&BenchmarkSuite{})

func (*BenchmarkSuite) BenchmarkAddAndEvict0000100(c *gc.C) {
	benchAddAndEvict(c, 100)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0000200(c *gc.C) {
	benchAddAndEvict(c, 200)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0000500(c *gc.C) {
	benchAddAndEvict(c, 500)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0001000(c *gc.C) {
	benchAddAndEvict(c, 1000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0002000(c *gc.C) {
	benchAddAndEvict(c, 2000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0005000(c *gc.C) {
	benchAddAndEvict(c, 5000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0010000(c *gc.C) {
	benchAddAndEvict(c, 10000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0020000(c *gc.C) {
	benchAddAndEvict(c, 20000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0050000(c *gc.C) {
	benchAddAndEvict(c, 50000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0100000(c *gc.C) {
	benchAddAndEvict(c, 100000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0200000(c *gc.C) {
	benchAddAndEvict(c, 200000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict0500000(c *gc.C) {
	benchAddAndEvict(c, 500000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict1000000(c *gc.C) {
	benchAddAndEvict(c, 1000000)
}

func (*BenchmarkSuite) BenchmarkAddAndEvict2000000(c *gc.C) {
	benchAddAndEvict(c, 2000000)
}

func benchAddAndEvict(c *gc.C, size int) {

	cache := lru.New(size)
	for i := 0; i < c.N; i++ {
		cache.Add(i, i)
	}
	expectLen := size
	if c.N < expectLen {
		expectLen = c.N
	}
	c.Assert(cache.Len(), gc.Equals, expectLen)
}

func (*BenchmarkSuite) BenchmarkGet0000010(c *gc.C) {
	benchGet(c, 10)
}

func (*BenchmarkSuite) BenchmarkGet0000020(c *gc.C) {
	benchGet(c, 20)
}

func (*BenchmarkSuite) BenchmarkGet0000050(c *gc.C) {
	benchGet(c, 50)
}

func (*BenchmarkSuite) BenchmarkGet0000100(c *gc.C) {
	benchGet(c, 100)
}

func (*BenchmarkSuite) BenchmarkGet0000200(c *gc.C) {
	benchGet(c, 200)
}

func (*BenchmarkSuite) BenchmarkGet0000500(c *gc.C) {
	benchGet(c, 500)
}

func (*BenchmarkSuite) BenchmarkGet0001000(c *gc.C) {
	benchGet(c, 1000)
}

func (*BenchmarkSuite) BenchmarkGet0002000(c *gc.C) {
	benchGet(c, 2000)
}

func (*BenchmarkSuite) BenchmarkGet0005000(c *gc.C) {
	benchGet(c, 5000)
}

func (*BenchmarkSuite) BenchmarkGet0010000(c *gc.C) {
	benchGet(c, 10000)
}

func (*BenchmarkSuite) BenchmarkGet0020000(c *gc.C) {
	benchGet(c, 20000)
}

func (*BenchmarkSuite) BenchmarkGet0050000(c *gc.C) {
	benchGet(c, 50000)
}

func (*BenchmarkSuite) BenchmarkGet0100000(c *gc.C) {
	benchGet(c, 100000)
}

func (*BenchmarkSuite) BenchmarkGet0200000(c *gc.C) {
	benchGet(c, 200000)
}

func (*BenchmarkSuite) BenchmarkGet0500000(c *gc.C) {
	benchGet(c, 500000)
}

func (*BenchmarkSuite) BenchmarkGet1000000(c *gc.C) {
	benchGet(c, 1000000)
}

func (*BenchmarkSuite) BenchmarkGet2000000(c *gc.C) {
	benchGet(c, 2000000)
}

func (*BenchmarkSuite) BenchmarkGet5000000(c *gc.C) {
	benchGet(c, 5000000)
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
