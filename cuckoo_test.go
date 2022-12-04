package testcmp

import (
	"github.com/zeebo/wyhash"
	"math/rand"
	"testing"
)

var rng wyhash.SRNG

func BenchmarkRandiGo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.Int31() % 2
	}
}

func BenchmarkRandiWyng(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uint32(uint64(uint32(rng.Uint64())) * uint64(2) >> 32)
	}
}
