/*
value receiver vs. pointer receiver
Two reasons to use a pointer receiver:
1. So that the method can modify the value that its receiver points to
2. To avoid copying the value on each method call
!! In general, all methods on a given type should have either value or pointer receivers
!! but not a mixture of both!
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// a method (defined on the struct Vertex)
// the (v Vertex) is called a receiver / a receiver argument
// the method should stay in the same package as the type/struct
// -------------------------------------------------------------
// HERE we have a "value receiver" and the struct is immutable
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// HERE we have a "pointer receiver" and the struct is mutable
// such a method can be called by both the struct pointer
// and the struct itself
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// actually (&v).Scale(2), but no problem
	// when calling a method, value and its pointer are interchangeable
	// BUT, when taking an argument, they are NOT interchangeable (statically typed)
	v.Scale(2)
	fmt.Println(v) // after scaling
	fmt.Println(v.Abs())
}
