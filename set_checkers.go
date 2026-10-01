package set

import "maps"

// IsEmpty returns true if the set is empty, else returns false
func (s set[T]) IsEmpty() bool {
	if len(s) <= 0 {
		return true
	}

	return false
}

// Equals returns true if rhs is equal to s, else returns false
func (s set[T]) Equals(rhs set[T]) bool {
	return maps.Equal(s, rhs)
}

// IsSubset return true if s1 is a subset of s2, else returns false
func IsSubset[T comparable](s1, s2 set[T]) bool {
	for k := range s1 {
		if !s2.Contains(k) {
			return false
		}
	}
	return true
}