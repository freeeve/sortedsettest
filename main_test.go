package main

import (
	"math/rand"
	"testing"
	. "github.com/wfreeman/check"
)

type IdSortedSetBenchSuite struct{}

var _ = Suite(&IdSortedSetBenchSuite{})

// gocheck link to go test - only needs doing once for whole lib
func Test(t *testing.T) {
	TestingT(t)
}

var r = rand.New(rand.NewSource(123123))

func (s *IdSortedSetBenchSuite) BenchmarkAddSet(c *C) {
	//	ss := NewIdSet()
	i := c.N
	for ; i > 0; i-- {
		//		ss.Add(Id(i))
	}
	//c.Assert(ss.Cardinality(), Equals, i)
}

func (s *IdSortedSetBenchSuite) BenchmarkAddSortedSet(c *C) {
	//	ss := NewIdSortedSet(func(a, b Id) bool { return a < b })
	i := c.N
	for ; i > 0; i-- {
		//		ss.Add(Id(i))
	}
	//c.Assert(ss.Cardinality(), Equals, i)
}
