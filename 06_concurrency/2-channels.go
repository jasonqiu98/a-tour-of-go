package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		// introduce some "sleep"
		// and then observe the "swap" between `x` and `y`
		time.Sleep(100 * time.Millisecond)
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// make a channel of int
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	// sometimes "17 -5 12", and sometimes "-5 17 12"
	fmt.Println(x, y, x+y)
}
