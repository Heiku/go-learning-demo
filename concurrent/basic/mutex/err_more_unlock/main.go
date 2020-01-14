package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		time.Sleep(5 * time.Second)

		// fatal error: sync: unlock of unlocked mutex
		mu.Unlock()
	}()

	time.Sleep(time.Second)
	mu.Unlock()
	fmt.Println("mutex had unlocked")

	select {}
}
