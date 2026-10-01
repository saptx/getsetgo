package set

import "maps"

func (s set[T]) IsEmpty() bool {
	if len(s) <= 0 {
		return true
	}

	return false
}

func (s set[T]) Equals(rhs set[T]) bool {
	return maps.Equal(s, rhs)
}
