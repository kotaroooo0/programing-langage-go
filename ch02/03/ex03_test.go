package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(4)
	}
}

// benchmark
// BenchmarkPopCount-4   	1000000000	         0.350 ns/op	       0 B/op	       0 allocs/op

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFor(4)
	}
}

// benchmark
// BenchmarkPopCountFor-4   	 2053081	       598 ns/op	       0 B/op	       0 allocs/op
