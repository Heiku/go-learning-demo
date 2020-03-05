package main

import (
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)

	go func() {
		time.Sleep(time.Second * 2)
		// block current coroutine
		waitGroup.Wait()
	}()

	waitGroup.Wait()
}
