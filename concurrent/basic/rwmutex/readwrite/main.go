package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(12)

	var rw sync.RWMutex
	go func() {
		defer wg.Done()

		rw.RLock()
		fmt.Println("before lock, reader0 RLock")
		time.Sleep(10 * time.Second)
		rw.RUnlock()

		fmt.Println("before lock, reader0 RUnlock")
	}()

	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()

		rw.Lock()
		fmt.Println("writer1 lock")
		time.Sleep(10 * time.Second)
		rw.Unlock()

		fmt.Println("writer1 UnLock")
	}()

	time.Sleep(time.Second)

	for i := 1; i <= 10; i++ {
		i := i
		go func() {
			defer wg.Done()

			rw.RLock()
			fmt.Printf("after lock, reader%d RLock\n", i)
			time.Sleep(5 * time.Second)
			rw.RUnlock()

			fmt.Printf("after lock, reader%d RUnlock\n", i)
		}()
	}

	time.Sleep(time.Second)
	wg.Wait()
}
