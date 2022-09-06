package main

import (
	"fmt"
)

var c, python, java bool = true, false, true

func main() {
	var i int // default value 0

	// 0, true, false, true
	fmt.Println(i, c, python, java)

	// Short variable declarations
	// equivalent to
	// var k int = 3
	k := 3
	fmt.Println(k)

	// Basic types:
	// bool, string, int/uint/uintptr, float32/float64, complex64/complex128
	// byte (uint8), rune (int32, a Unicode code point, default type for a 'char')
	// Note: int, uint and uintptr are 32 bits on 32-bit systems, 64 bits on 64-bit systems

	// a factored statement
	var (
		m float32 = 2.1
		n int     = 5
	)

	// type conversion
	fmt.Println(m + float32(n))

	// equivalent to q := 8
	// without specifying an explicit type
	var q = 8

	fmt.Println(q)

}
