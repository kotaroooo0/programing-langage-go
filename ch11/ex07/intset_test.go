package ex07_test

import (
	"math/rand"
	"testing"

	"github.com/kotaroooo0/programing-language-go/ch11/ex07"
)

func randomSet(set *ex07.IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(i)
	}
}

func BenchmarkAdd(b *testing.B) {
	set := ex07.IntSet{}
	randomSet(&set, 100)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(10000))
	}
}
func BenchmarkUnionWith(b *testing.B) {
	set1 := ex07.IntSet{}
	randomSet(&set1, 100)

	set2 := ex07.IntSet{}
	randomSet(&set2, 100)
	for i := 0; i < b.N; i++ {
		set1.UnionWith(&set2)
	}
}
