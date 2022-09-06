/*
Type parameters
Generic types
*/

package main

import "fmt"

// Generics / Generic types
// any is an alias for interface{}
type List[T any] struct {
	next *List[T]
	val  T
}

// Type parameters
// Index returns the index of x in s, or -1 if not found
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	node := List[int]{nil, 3}
	node2 := List[int]{&node, 2}
	fmt.Println(node2)

}
