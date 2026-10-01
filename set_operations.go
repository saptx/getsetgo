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
	There are mathematical operations that can be performed in a Set.
*/

// Union returns a new set with elements of both the sets
func (s set[T]) Union(rhs set[T]) set[T] {
	tempSet := maps.Clone(s)
	for k := range rhs {
		tempSet.Add(k)
	}
	return tempSet
}

func (s set[T]) UnionUpdate(rhs set[T]) {
	for k := range rhs {
		s.Add(k)
	}
}
