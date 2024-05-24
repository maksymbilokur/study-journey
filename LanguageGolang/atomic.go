package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// SafeCounter represents a counter that can be safely accessed concurrently.
type SafeCounter struct {
	value int64
	mu    sync.Mutex // For illustration purposes only
}

// Increment increments the counter using sync.Mutex for synchronization.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// GetValue returns the current value of the counter using sync.Mutex for synchronization.
func (c *SafeCounter) GetValue() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// AtomicCounter represents a counter that can be safely accessed concurrently using atomic operations.
type AtomicCounter struct {
	value int64
}

// Increment increments the counter using atomic operations.
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// GetValue returns the current value of the counter using atomic operations.
func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	// Using SafeCounter with sync.Mutex
	safeCounter := SafeCounter{}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			safeCounter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("SafeCounter value:", safeCounter.GetValue())

	// Using AtomicCounter with atomic operations
	atomicCounter := AtomicCounter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomicCounter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("AtomicCounter value:", atomicCounter.GetValue())
}
