package bloom

import (
	"fmt"
	"math"
)

type BloomFilter struct {
	bitSet      []uint64 // The bit array represented as a slice of bool
	m           uint     // The number of bits in the bit set (shortcut for len(bitSet)
	k           uint     // The number of hash functions to use (shortcut for len(hashes)
	hasher      Hasher   // The hash functions to use
	count       uint     // Number of items added
	scaleFactor float64  // Factor to scale when capacity is reached
}

// Hasher is an interface for hash function providers
type Hasher interface {
	Hash(data []byte) (uint64, uint64)
}

// NewBloomFilterWithHasher creates a new Bloom filter with the given number
// of elements (n) and false positive rate (p).
func New(initialCapacity uint, falsePositiveRate float64, h Hasher) (*BloomFilter, error) {
	if initialCapacity == 0 {
		return nil, fmt.Errorf("initial capacity must be > 0")
	}
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1")
	}
	if h == nil {
		return nil, fmt.Errorf("hasher cannot be nil")
	}

	m, k := optimalParams(initialCapacity, falsePositiveRate)
	return &BloomFilter{
		bitSet:      make([]uint64, (m+63)/64),
		m:           m,
		k:           k,
		hasher:      h,
		scaleFactor: 2,
	}, nil
}

// getOptimalParams calculates the optimal parameters for the Bloom filter,
// the number of bits in the bit set (m) and the number of hash functions (k).
func optimalParams(n uint, p float64) (m, k uint) {
	m = uint(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	k = uint(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
	if m == 0 {
		m = 1
	}
	if k == 0 {
		k = 1
	}
	return
}

// Add adds an item to the Bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	if bf.shouldScale() {
		bf.scale()
	}

	for i := uint(0); i < bf.k; i++ {
		index := bf.getIndex(data, i)
		bf.setBit(index)
	}

	bf.count++
}

// Test checks if an item is in the Bloom filter
func (bf *BloomFilter) Test(data []byte) bool {
	for i := uint(0); i < bf.k; i++ {
		index := bf.getIndex(data, i)
		if !bf.testBit(index) {
			return false
		}
	}
	return true
}
func (bf *BloomFilter) getIndex(data []byte, i uint) uint {
	h1, h2 := bf.hasher.Hash(data)
	return uint((h1 + uint64(i)*h2) % uint64(bf.m))
}

// setBit sets the bit at the specified index
func (bf *BloomFilter) setBit(index uint) {
	bf.bitSet[index/64] |= 1 << (index % 64)
}

// testBit checks if the bit at the specified index is set
func (bf *BloomFilter) testBit(index uint) bool {
	return bf.bitSet[index/64]&(1<<(index%64)) != 0
}

// shouldScale determines if the Bloom filter should be scaled
func (bf *BloomFilter) shouldScale() bool {
	return float64(bf.count) > float64(bf.m)*0.5
}

// scale increases the size of the Bloom filter
func (bf *BloomFilter) scale() {
	newM := uint(float64(bf.m) * bf.scaleFactor)
	newBitSet := make([]uint64, (newM+63)/64)
	copy(newBitSet, bf.bitSet)
	bf.bitSet = newBitSet
	bf.m = newM
}

// EstimatedFalsePositiveRate returns the estimated false positive rate
func (bf *BloomFilter) EstimatedFalsePositiveRate() float64 {
	return math.Pow(1-math.Exp(-float64(bf.k)*float64(bf.count)/float64(bf.m)), float64(bf.k))
}

// Info returns current information about the Bloom filter
func (bf *BloomFilter) Info() string {
	return fmt.Sprintf("Items: %d, Bits: %d, Hash Functions: %d, Est. False Positive Rate: %.4f",
		bf.count, bf.m, bf.k, bf.EstimatedFalsePositiveRate())
}
