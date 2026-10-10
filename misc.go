package set

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