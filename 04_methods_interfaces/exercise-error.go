/*
type error interface {
	Error() string
}
*/

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// those structs implementing the Error() method
// can be directly used as an implementation of the "error" interface
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func CalcSqrt(x float64) (float64, error) {
	if x < 0 {
		// here ErrNegativeSqrt(x) is an implementation of the "error" interface
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	delta := (z*z - x) / (2.0 * z)
	// the exponent can go up to 1e-15
	for math.Abs(delta) > 1e-12 {
		z -= delta
		delta = (z*z - x) / (2.0 * z)
	}
	return z, nil
}

func RunSqrt(x float64) {
	if sqrt, err := CalcSqrt(x); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt)
	}
}

func main() {
	RunSqrt(2)
	RunSqrt(-2)
}
