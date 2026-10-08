package set

type set[T comparable] map[T]struct{}

// Set returns a new Thread Unsafe set with the given initial values. 
// To create an empty set, use:
//   st := Set[int]()
func Set[T comparable](values ...T) set[T] {
	s := make(set[T], len(values))
	for _, v := range values {
		s[v] = struct{}{}
	}

	return s
}
