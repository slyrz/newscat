package util

const (
	// 61 is the greatest prime below 64. Depending on the input values,
	// a prime number might lead to a better distribution.
	bitsetCap = 61
)

type Bitset struct {
	Bucket uint64
}

func NewBitset() *Bitset {
	return new(Bitset)
}

// Adds a value to Bitset.
func (f *Bitset) Add(val uint32) {
	f.Bucket |= 1 << (val % bitsetCap)
}

// Adds elements from g to Bitset.
func (f *Bitset) Union(g *Bitset) {
	f.Bucket |= g.Bucket
}

// Return the number of buckets in use.
func (f *Bitset) Len() int {
	return popCount(f.Bucket)
}

// Return the total number of buckets.
func (f *Bitset) Cap() int {
	return bitsetCap
}

// Returns the number of common elements.
func (f *Bitset) Common(g *Bitset) int {
	return popCount(f.Bucket & g.Bucket)
}

// Returns the number of bits set.
// Source:
//	http://graphics.stanford.edu/~seander/bithacks.html
func popCount(x uint64) (n int) {
	x -= (x >> 1) & 0x5555555555555555
	x  = (x >> 2) & 0x3333333333333333 + x & 0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return int(x >> 56)
}

type Stringset struct {
	Bitset
}

func NewStringset() *Stringset {
	return new(Stringset)
}

// Adds a value to Stringset.
func (f *Stringset) Add(val string) {
	f.Bitset.Add(Hash(val))
}
