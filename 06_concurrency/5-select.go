package main

import "fmt"

// A select blocks until one of its cases can run, then it executes that case.
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// this can happen only if someone tries to receive from `c`
		case c <- x:
			x, y = y, x+y
		// this can happen only if someone tries to send to `quit`
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
--------------------------------------
fibonacci()
-- tries to receive from `c` 10 times
-- and sends to `quit`
--------------------------------------
main() starts a goroutine to
-- 1. wait to send to `c`
-- 2. wait to receive from `quit`
--------------------------------------
*/
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// receive from c
			fmt.Println(<-c)
		}
		// send to quit
		quit <- 0
	}()
	fibonacci(c, quit)
}
