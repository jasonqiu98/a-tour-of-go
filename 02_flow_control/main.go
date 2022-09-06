package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	delta := (z*z - x) / (2.0 * z)
	// the exponent can go up to 1e-15
	for math.Abs(delta) > 1e-12 {
		z -= delta
		delta = (z*z - x) / (2.0 * z)
	}
	return z
}

func main() {

	// deferred execution
	// after the function returns
	// LIFO, "stack"
	defer fmt.Print("Jason!\n")
	defer fmt.Print("Bye bye, ")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// "C's while is spelled for in Go"
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// in this "argument list", all items should trail with a comma
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Exercise: Loops and Functions
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	// switch-case
	// we don't need to attach "break" to each case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// can put an variable into each case
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch-case with no condition
	// equivalent to an if-then-else chain
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
