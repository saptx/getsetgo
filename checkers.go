package set

import "maps"

/*
	All these methods returns a boolean value after checking the necessary conditions.
*/

// IsEmpty returns true if the set is empty, else returns false
func (s set[T]) IsEmpty() bool {
	if len(s) <= 0 {
		return true
	}

	return false
}

// Equals returns true if set1 is equal to set2, else returns false.
//
// set1 & set2 must conatain data of same type.
func Equal[T comparable](set1, set2 set[T]) bool {
	return maps.Equal(set1, set2)
}

// IsSubset returns true if s1 is a subset of s2, else returns false.
//
// s1 & s2 must contain data of same type.
//
// A set s1 is a subset of s2 if every element of s1 is also an element of s2.
// For example, if A = {1, 2, 3} and B = {2, 3}, then B is a subset of A.
func IsSubset[T comparable](s1, s2 set[T]) bool {
	for k := range s1 {
		if !s2.Contains(k) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s1 is a superset of s2, else returns false.
//
// s1 & s2 must contain data of same type.
//
// A set s1 is a superset of s2 if every element of s2 is also an element of s1.
// For example, if A = {1, 2, 3} and B = {2, 3}, then A is a superset of B.
func IsSuperset[T comparable](s1, s2 set[T]) bool {
	return IsSubset(s2, s1)
}