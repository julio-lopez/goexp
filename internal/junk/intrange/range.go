package intrange

import (
	"errors"
	"fmt"
)

type closedIntRange struct {
	lo, hi int
}

// constants from the standard math package.
const (
	//nolint:mnd
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	maxInt = 1<<(intSize-1) - 1
	minInt = -1 << (intSize - 1)
)

var errNonContiguousRange = errors.New("non-contiguous range")

func (r closedIntRange) length() uint {
	// any range where lo > hi is empty. The canonical empty representation
	// is {lo:0, hi: -1}
	if r.lo > r.hi {
		return 0
	}

	return uint(r.hi - r.lo + 1)
}

func (r closedIntRange) isEmpty() bool {
	return r.length() == 0
}

// Returns a range for the keys in m. It returns an empty range when m is empty.
func getKeyRange[E any](m map[int]E) closedIntRange {
	if len(m) == 0 {
		return closedIntRange{lo: 0, hi: -1}
	}

	lo, hi := maxInt, minInt
	for k := range m {
		if k < lo {
			lo = k
		}

		if k > hi {
			hi = k
		}
	}

	return closedIntRange{lo: lo, hi: hi}
}

// Returns a contiguous range for the keys in m.
// When the range is not continuous an error is returned.
func getContiguousKeyRange[E any](m map[int]E) (closedIntRange, error) {
	r := getKeyRange(m)

	// r.hi and r.lo are from unique map keys, so for the range to be continuous
	// then the range length must be exactly the same as the size of the map.
	// For example, if lo==2, hi==4, and len(m) == 3, the range must be
	// contiguous => {2, 3, 4}
	if r.length() != uint(len(m)) {
		return closedIntRange{-1, -2}, fmt.Errorf("[lo: %d, hi: %d], length: %d: %w", r.lo, r.hi, len(m), errNonContiguousRange)
	}

	return r, nil
}
