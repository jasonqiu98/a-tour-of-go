/*
pointer, array, slice, map
*/

package main

import (
	"fmt"
	"strings"
)

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// Exercise: Maps
func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, ch := range strings.Fields(s) {
		res[string(ch)] += 1
	}
	return res
}

func main() {
	// pointers
	var i, j int = 42, 2701
	var p *int = &i
	fmt.Println("the address of i", p)
	fmt.Println("the value of i", *p)
	fmt.Println("Changing the value of i...")
	*p = 21 // changing the value of i
	fmt.Println("the address of i", p)
	fmt.Println("the value of i", *p)

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j) // new value of j

	// Go has no pointer arithemetic
	// i.e., cannot do the op like (p + 1)

	// print a struct
	var v Vertex = Vertex{1, 2}
	fmt.Println("v: ", v)
	fmt.Println("v.X: ", v.X)
	// v.X = 4
	fmt.Println("v: ", v)

	// struct fields can be accessed through a struct pointer
	var p_v *Vertex = &v
	fmt.Println("p_v.X: ", p_v.X) // a shorthand for (*p_v).X

	var (
		v1   = Vertex{1, 2}  // has type Vertex
		v2   = Vertex{X: 1}  // Y:0 is implicit
		v3   = Vertex{}      // X:0 and Y:0
		p_v4 = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, v2, v3, *p_v4)

	// array
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// array literal
	// cannot omit [6]int before the literal here!
	// var primes [6]int = [6]int{2, 3, 5, 7, 11, 13}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// a slice type
	// -- a dynamically sized, flexible view
	// cannot put the size in a slice type
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	slice1 := names[:2] // equivalent to names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)

	// a slice "describes" a section of an underlying array
	// slices are like "references"/pointers to arrays
	slice2[0] = "XXX"
	fmt.Println(names)

	// slice literals
	slice3 := []int{2, 3, 5, 7, 11, 13}
	slice4 := []bool{true, false, true, true, false, true}
	slice5 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(slice3, slice4, slice5)

	// length: length of the slice
	// capacity: length of the underlying array
	// the capacity depends on where the slice starts

	// the slice starts at index 1 of the array
	// so the capacity is (length of the array - 1) = (6 - 1) = 5
	s = s[:]
	fmt.Printf("[1:], len=%d cap=%d %v\n", len(s), cap(s), s)

	// the slice starts at index 1
	s = s[:0]
	fmt.Printf("[1:1], len=%d cap=%d %v\n", len(s), cap(s), s)

	// the slice starts at index 1
	s = s[:4]
	fmt.Printf("[1:5], len=%d cap=%d %v\n", len(s), cap(s), s)

	// the slice starts at index (1 + 2) = 3
	// the capacity is 6 - 3 = 3
	s = s[2:]
	fmt.Printf("[3:], len=%d cap=%d %v\n", len(s), cap(s), s)

	// the slice starts at index 2
	// the capacity is 6 - 2 = 4
	s = primes[2:]
	fmt.Printf("[2:], len=%d cap=%d %v\n", len(s), cap(s), s)

	// nil is the zero value of a slice
	// don't initialize the slice like this; use make() instead
	// otherwise every append will lead to auto increase of the size
	var slice6 []int
	fmt.Println(slice6, len(slice6), cap(slice6))
	if slice6 == nil {
		fmt.Println("nil!")
	}

	// creating a slice with make
	// allocates a zeroed array
	slice7 := make([]int, 5) // type, len // equivalent to make([]int, 5, 5)
	printSlice("slice7", slice7)

	slice8 := make([]int, 0, 5) // type, len, cap
	printSlice("slice8", slice8)

	// slices of slices, 2D arrays
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// auto increase capacity
	// we should try to avoid auto increasing capacity by allocating enough space ourselves
	slice7 = append(slice7, 0)
	printSlice("slice7", slice7)

	// append multiple elements
	slice8 = append(slice8, 1, 2, 3)
	printSlice("slice8", slice8)

	// range ~ like Python's enumerate
	var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow2 {
		fmt.Printf("2**%d=%d ", i, v)
	}
	fmt.Println()

	// range ~ like Python's range(len())
	for i := range pow2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// range ~ like Python's in
	for _, v := range pow2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// map
	m := make(map[string]Vertex)
	m["Hello World"] = Vertex{
		5, 10,
	}
	fmt.Println(m)

	// if the key doesn't exist
	// return a zero value of the value
	fmt.Println(m["Hello"])

	// map literals
	// if the top level type is just a type name
	// you can omit it from the elements of the literal
	m1 := map[string]Vertex{
		"Hello": {
			1, 2,
		},
		"World": {
			2, 3,
		},
	}
	fmt.Println(m1)

	// Mutating Maps
	m1["World"] = Vertex{3, 5}
	fmt.Println(m1)

	// delete a key
	delete(m1, "Hello")
	fmt.Println(m1)

	// v, ok
	// v - value
	// ok - exist or not
	v_hello, ok_hello := m1["Hello"]
	fmt.Println("The value:", v_hello, "Present?", ok_hello)

}
