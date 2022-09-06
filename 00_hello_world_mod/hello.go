// import, print and functions

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"rsc.io/quote"
)

// single result
func add(x int, y int) int {
	return x + y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	//  Don't communicate by sharing memory, share memory by communicating.
	fmt.Println(quote.Go())

	// The time is 2022-09-03 16:19:53.765988958 +0000 UTC m=+0.000580794
	fmt.Println("The time is", time.Now())

	// My favorite number is 1
	fmt.Println("My favorite number is", rand.Intn(10))

	// Now you have 2.6457513110645907 problems.
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// The value of PI is 3.141592653589793
	// Sprintf formats according to a format specifier and returns the resulting string.
	s := fmt.Sprintf("The value of PI is %g", math.Pi)
	fmt.Println(s)

	// 55
	fmt.Println(add(42, 13))

	// world hello
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// hello world
	a, b = b, a
	fmt.Println(a, b)

	// 7 10
	fmt.Println(split(17))
}
