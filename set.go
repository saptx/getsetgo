package set

type set[T comparable] map[T]struct{}

func Set[T comparable](values... T) set[T] {
	s := make(set[T], len(values))
	for _, v := range values {
		s[v] = struct{}{}
	}

	return s
}