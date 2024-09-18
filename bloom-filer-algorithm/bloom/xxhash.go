package bloom

import (
	"github.com/cespare/xxhash/v2"
)

type XXHasher struct{}

func NewXXHasher() *XXHasher {
	return &XXHasher{}
}

func (h *XXHasher) Hash(data []byte) (uint64, uint64) {
	h1 := xxhash.Sum64(data)
	h2 := xxhash.Sum64(append(data, byte(1)))
	return h1, h2
}
