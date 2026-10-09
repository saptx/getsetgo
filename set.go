package set

import (
	"fmt"
	"maps"
	"strings"
)

// Thread Unsafe
type set[T comparable] map[T]struct{}

// Set returns a new Thread Unsafe set with the given initial values.
// To create an empty set, use:
//
//	st := Set[int]()
func Set[T comparable](values ...T) set[T] {
	s := make(set[T], len(values))
	for _, v := range values {
		s[v] = struct{}{}
	}

	return s
}

/*
	Implement fmt.Stringer interface for formatting the output when using methods like fmt.Println() etc.
	This is done by defining a String() method.
*/

func (s set[T]) String() string {
	elements := make([]string, 0, len(s))
	for el := range s {
		elements = append(elements, fmt.Sprintf("%v", el))
	}

	return "Set {" + strings.Join(elements, " ") + "}"
}

/*
	Primary methods.
	All other methods are defined using these 4 methods.
*/

// Add inserts new item into the set.
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
	Methods to perform Mathematical set operations.
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

// Intersection returns a new set containing only the elements present in both s and rhs.
// Neither s nor rhs is modified.
func (s set[T]) Intersection(rhs set[T]) set[T] {
	tempSet := make(set[T])

	minSet := s
	if len(rhs) < len(s) {
		minSet = rhs
	}
	
	// set haeving least elements is used for iteration
	for k := range minSet {
		if s.Contains(k) && rhs.Contains(k) {
			tempSet[k] = struct{}{}
		}
	}

	return tempSet
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

// SymmetricDiff returns a new set containing the elements present in either of the sets,
// but not in both.
// Neither s nor rhs is modified.
func (s set[T]) SymmetricDiff(rhs set[T]) set[T] {
	intersection := s.Intersection(rhs)
	return s.Diff(intersection).Union(rhs.Diff(intersection))
}