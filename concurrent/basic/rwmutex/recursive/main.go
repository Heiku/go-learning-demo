package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()

		fmt.Println("Unlock")
	}()

	go func() {
		defer wg.Done()
		factorial(&mu, 4)
	}()

	wg.Wait()
}

// did't allow to recursive use rwMutex
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n
}
