package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// initialize with the {key:value} format
	c := SafeCounter{v: make(map[string]int)}

	// starts 1000 goroutines
	for i := 0; i < 1000; i++ {
		// with the lock, the increase is safe
		go c.Inc("somekey")
	}

	// sleep 20ms
	time.Sleep(20 * time.Millisecond)

	// check the results
	// with the lock, the access is safe
	fmt.Println(c.Value("somekey"))
}
