package main

import "fmt"

func main() {
	// a buffered channel of size 2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
