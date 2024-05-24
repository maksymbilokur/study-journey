/*
go run -race channels2.go
gcc must be installed

*/

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.m[key]
	return value, ok
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		safeMap.Set("key1", 1)
	}()
	go func() {
		defer wg.Done()

		safeMap.Set("key2", 2)
	}()

	wg.Wait()

	value1, found1 := safeMap.Get("key1")
	value2, found2 := safeMap.Get("key2")

	fmt.Println("Key1 value:", value1, "found:", found1)
	fmt.Println("Key2 value:", value2, "found:", found2)
}
