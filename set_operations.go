package set

import "maps"

/*
	Primary methods.
	All other methods can be derived from these fucntions
*/

// Add inserts the given element in the set if not present
func (s set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove deletes the item from the set if exists. It returns nothing
func (s set[T]) Remove(item T) {
	delete(s, item)
}

// Contains returns true if the given item is present in the set, else returns false
func (s set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Size returns the total count of elements in the set
func (s set[T]) Size() int {
	return len(s)
}

/*
	There are mathematical operations that can be performed in a set.
*/

// Union returns a new set with elements of both the sets
func (s set[T]) Union(rhs set[T]) set[T] {
	tempset := maps.Clone(s)
	for k := range rhs {
		tempset.Add(k)
	}
	return tempset
}

// UnionUpdate performs and stored the result on the exitsing set without returning a new set
func (s set[T]) UnionUpdate(rhs set[T]) {
	for k := range rhs {
		s.Add(k)
	}
}

// Intersection returns a new set with those common elements of both sets
// If no common elements are found, empty set is returned.
func (s set[T]) Intersection(rhs set[T]) set[T] {
	tempset := make(set[T])
	minset := make(set[T])
	if len(s) > len(rhs) {
		minset = rhs
	} else {
		minset = s
	}

	for k := range maps.Keys(minset) {
		if s.Contains(k) == rhs.Contains(k) {
			tempset[k] = struct{}{}
		}
	}

	return tempset
}

func (s set[T]) IntersectionUpdate(rhs set[T]) {
	for k := range maps.Keys(s) {
		if s.Contains(k) == rhs.Contains(k) {
			s[k] = struct{}{}
		}else{
			delete(s, k)
		}
	}
}
