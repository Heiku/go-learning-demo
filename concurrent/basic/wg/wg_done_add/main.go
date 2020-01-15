package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()

		// sync: WaitGroup is reused before previous Wait has returned
		// wait() isn't a atomic operation, may add() happened in wait()
		wg.Add(1)
	}()
	wg.Wait()
}
