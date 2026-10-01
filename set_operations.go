package set

import "maps"

/*
	Primary methods.
	All other set operations can be derived from these methods.
*/

// Add inserts item into the set.
func (s set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove deletes item from the set if it exists.
func (s set[T]) Remove(item T) {
	delete(s, item)
}

// Contains reports whether item is present in the set.
func (s set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Size returns the number of elements in the set.
func (s set[T]) Size() int {
	return len(s)
}

/*
	Mathematical set operations.
*/

// Union returns a new set containing all elements present in either s or rhs.
// Neither s nor rhs is modified.
func (s set[T]) Union(rhs set[T]) set[T] {
	tempSet := maps.Clone(s)
	for k := range rhs {
		tempSet.Add(k)
	}
	return tempSet
}

// UnionUpdate modifies s to contain all elements present in either s or rhs.
func (s set[T]) UnionUpdate(rhs set[T]) {
	for k := range rhs {
		s.Add(k)
	}
}

// Intersection returns a new set containing only the elements present in both s and rhs.
// Neither s nor rhs is modified.
func (s set[T]) Intersection(rhs set[T]) set[T] {
	tempSet := make(set[T])

	minSet := s
	if len(rhs) < len(s) {
		minSet = rhs
	}

	for k := range minSet {
		if s.Contains(k) && rhs.Contains(k) {
			tempSet[k] = struct{}{}
		}
	}

	return tempSet
}

// IntersectionUpdate modifies s to contain only the elements that are also present in rhs.
func (s set[T]) IntersectionUpdate(rhs set[T]) {
	for k := range s {
		if !rhs.Contains(k) {
			delete(s, k)
		}
	}
}

// Diff returns a new set containing the elements present in s but not in rhs.
// Neither s nor rhs is modified.
func (s set[T]) Diff(rhs set[T]) set[T] {
	tempSet := maps.Clone(s)

	for k := range rhs {
		tempSet.Remove(k)
	}

	return tempSet
}

// DiffUpdate modifies s by removing all elements that are also present in rhs.
func (s set[T]) DiffUpdate(rhs set[T]) {
	for k := range rhs {
		s.Remove(k)
	}
}

// SymmetricDiff returns a new set containing the elements present in either of the sets,
// but not in both
func (s set[T]) SymmetricDiff(rhs set[T]) set[T] {
	intersection := s.Intersection(rhs)
	return s.Diff(intersection).Union(rhs.Diff(intersection))
}

// SymmetricDiffUpdate modifies s to contain the elements present in either s or rhs,
// but not in both.
func (s set[T]) SymmetricDiffUpdate(rhs set[T]) {
	for k := range rhs {
		if s.Contains(k) {
			s.Remove(k)
		} else {
			s.Add(k)
		}
	}
}