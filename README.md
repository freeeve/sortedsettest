sortedsettest
=============

just a test

#### benchmarks
Compare adding to a map-based hash set to a skiplist. Skiplist roughly 4x slower.

(master ✗) wes-macbook:sortedsettest go test -check.b -check.bmem
PASS: main_test.go:17: IdSortedSetBenchSuite.BenchmarkAddSet	10000000	       247 ns/op	      22 B/op	       0 allocs/op
PASS: main_test.go:26: IdSortedSetBenchSuite.BenchmarkAddSortedSet	 1000000	      1094 ns/op	      47 B/op	       2 allocs/op
OK: 2 passed
PASS
ok  	github.com/wfreeman/sortedsettest	3.846s
