package testcmp

import (
	"github.com/zeebo/wyhash"
	"math/rand"
	"testing"
)

var rng wyhash.SRNG

func TestHash(t *testing.T) {
	for i := 0; i < 65536; i++ {
		if getXXH3AltIndexByStatic(uint16(i)) != getXXH3AltIndex(uint16(i)) {
			t.Errorf("incorrect hash: %d", i)
		}
	}
}
func BenchmarkRandiGoMod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.Int31() % 2
	}
}

func BenchmarkRandiGo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uint32(uint64(uint32(rand.Int31())) * uint64(2) >> 32)
	}
}

func BenchmarkRandiWyrng(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uint32(uint64(uint32(rng.Uint64())) * uint64(2) >> 32)
	}
}

func BenchmarkRandiWyrngMod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Uint64() % 2
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
