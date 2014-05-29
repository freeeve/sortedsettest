package main

import (
	"testing"
	. "github.com/wfreeman/check"
)

type IdSortedSetBenchSuite struct{}

var _ = Suite(&IdSortedSetBenchSuite{})

// gocheck link to go test - only needs doing once for whole lib
func Test(t *testing.T) {
	TestingT(t)
}

func (s *IdSortedSetBenchSuite) BenchmarkAdd(c *C) {
	/*
		ss := NewIdSortedSet(func(a, b Id) bool { return a < b })
		i := 0
		for ; i < c.N; i++ {
			ss.Add(Id(i))
		}
		c.Assert(ss.Cardinality(), Equals, i)
	*/
}
