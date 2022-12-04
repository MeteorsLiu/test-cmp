package testcmp

import (
	"encoding/binary"

	metro "github.com/dgryski/go-metro"
	"github.com/zeebo/xxh3"
)

var (
	altHash = [65536]uint{}
)

func init() {
	b := make([]byte, 2)
	endian := binary.LittleEndian
	for i := 0; i < 65536; i++ {
		endian.PutUint16(b, uint16(i))
		altHash[i] = (uint(xxh3.Hash(b)))
	}
}

func getMetroAltIndex(fp uint16) uint {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(fp))
	hash := uint(metro.Hash64(b, 1337))
	return hash
}

func getXXH3AltIndex(fp uint16) uint {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(fp))
	hash := uint(xxh3.Hash(b))
	return hash
}

func getXXH3AltIndexByStatic(fp uint16) uint {
	return altHash[fp]
}
