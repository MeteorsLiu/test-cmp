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

func BenchmarkMetroHash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getMetroAltIndex(uint16(i))
	}
}
func BenchmarkXXH3Hash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getXXH3AltIndex(uint16(i))
	}
}

func BenchmarkStaticHash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getXXH3AltIndexByStatic(uint16(i))
	}
}
