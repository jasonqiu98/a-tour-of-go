/*
functions and closures
*/

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
// a closure is a function value that references variables from outside its body
// the returned function refereces the variable "sum"
func adder() func(int) int {
	sum := 0 // local (initial) variable cannot be accessed from outside
	return func(x int) int {
		sum += x // operations
		return sum
	}
}

func fibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		temp := n1
		n1, n2 = n2, n1+n2
		return temp
	}
}

func main() {
	// function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81

	// use the closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// Fibonacci closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Println()
}
