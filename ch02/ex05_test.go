package popcount

import (
	"testing"
)

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(4)
	}
}

// benchmark
// BenchmarkPopCount3-4   	342246792	         3.06 ns/op	       0 B/op	       0 allocs/op
