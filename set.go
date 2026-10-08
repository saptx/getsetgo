package set

import (
	"fmt"
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
