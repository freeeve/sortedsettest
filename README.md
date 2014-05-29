sortedsettest
=============

just a test

#### benchmarks
Compare adding to a map-based hash set to a skiplist. Skiplist roughly 3-10x slower, for the benefit of being in order.

~~~
# in order, increasing
(master ✗) wes-macbook:sortedsettest go test -check.b -check.bmem
PASS: main_test.go:17: IdSortedSetBenchSuite.BenchmarkAddSet	10000000	       247 ns/op	      22 B/op	       0 allocs/op
PASS: main_test.go:26: IdSortedSetBenchSuite.BenchmarkAddSortedSet	 1000000	      1094 ns/op	      47 B/op	       2 allocs/op
OK: 2 passed
PASS
~~~

~~~
# in order, decreasing
(master ✗) wes-macbook:sortedsettest go test -check.b -check.bmem
PASS: main_test.go:20: IdSortedSetBenchSuite.BenchmarkAddSet	10000000	       236 ns/op	      22 B/op	       0 allocs/op
PASS: main_test.go:29: IdSortedSetBenchSuite.BenchmarkAddSortedSet	 2000000	       834 ns/op	      47 B/op	       2 allocs/op
OK: 2 passed
PASS
~~~

~~~
# random
PASS: main_test.go:20: IdSortedSetBenchSuite.BenchmarkAddSet	10000000	       263 ns/op	      22 B/op	       0 allocs/op
PASS: main_test.go:29: IdSortedSetBenchSuite.BenchmarkAddSortedSet	 1000000	      2651 ns/op	      47 B/op	       2 allocs/op
OK: 2 passed
PASS
~~~
