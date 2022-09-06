package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// only the sender should close a channel
	// never the receiver!
	// sending on a closed channel will cause a panic!
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// the loop receives values from the channel repeatedly
	// after the channel is closed
	// the remaining elements can be received
	// but no further elements will be sent (since it is closed)
	for i := range c {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
	v, ok := <-c
	// 0 false
	// implies the channel is closed
	fmt.Println(v, ok)
}
