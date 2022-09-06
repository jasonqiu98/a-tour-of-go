/*
1. Interfaces are implemented implicitly by methods.
2. There is no null pointer exception in Go.
   Null pointers are handled by nil receivers.
3. However, calling a method on a nil interface is a run-time error.
4. An empty interface `interface{}` is used hold values of any/unknown type.
5. type assertions: `t, ok := i.(T)`, used on interfaces
   used to make sure the type is correct
   if not correct: (1) with `ok`, then ok = false
                   (2) without `ok`: panic error

6. fmt.Stringer interface
   ```
   type Stringer interfacd {
       String() string
   }
   ```
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	// (value, type)
	fmt.Printf("(%v, %T)\n", a, a)
	s := a.(MyFloat)
	fmt.Println(s)

	s, ok := a.(MyFloat)
	fmt.Println(s, ok)

	a = &v // a *Vertex implements Abser
	// (value, type)
	fmt.Printf("(%v, %T)\n", a, a)

	r, ok := a.(*Vertex)
	fmt.Println(r, ok)

	// type switch
	// b := a.(type)
	// b returns the value; case followed by type
	switch b := a.(type) {
	case MyFloat:
		fmt.Println("MyFloat:", b)
	case *Vertex:
		fmt.Println("*Vertex:", b)
	default:
		fmt.Println("Others:", b)
	}

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

// Interfaces are implemented implicitly
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
