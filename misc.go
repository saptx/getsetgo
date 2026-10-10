package set

import (
	"cmp"
	"slices"
)

// todo:
// add Clone(), Sort()

/*
	Other useful functions
*/

// ToSlice returns slice containing the elements of the set
func (s set[T]) ToSlice() []T {
	slc := make([]T, 0)
	for k := range s {
		slc = append(slc, k)
	}
	return slc
}

// Sort returns set with elements sorted in ascending order
func Sort[E cmp.Ordered](set set[E]) []E {
	s := set.ToSlice()
	slices.Sort(s)
	return s
}
