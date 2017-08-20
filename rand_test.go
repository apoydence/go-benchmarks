package bench_test

import (
	"math/rand"
	"testing"
)

func BenchmarkRand(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = rand.Intn(1000)
	}
	_ = x
}
