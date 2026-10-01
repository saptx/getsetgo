package set

import (
	"fmt"
	"strings"
)

type set[T comparable] map[T]struct{}

/*
	Constructor to initialise a new Set
*/

// Set is used to initalise a new set with inital values. It returns a set with the given values.
// To create an empty set, use:
//   st := Set[int]()
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
