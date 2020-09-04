package popcount

import (
	"testing"
)

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(4)
	}
}

// benchmark
// BenchmarkPopCount2-4   	19320066	        60.5 ns/op	       0 B/op	       0 allocs/op
