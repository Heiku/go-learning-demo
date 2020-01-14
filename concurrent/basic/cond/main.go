package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// cond wait()、broadcast()、signal()    must be lock status(just like synchronized in Java, used in Monitor)

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)

	ready := make(chan struct{}, 10)
	isReady := false

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			m.Lock()

			time.Sleep(time.Duration(rand.Int63n(2)) * time.Second)

			ready <- struct{}{}
			for !isReady {
				// release lock monitor
				c.Wait()
			}
			log.Printf("%d run\n", i)
			m.Unlock()
		}()
	}

	// signalAll but all goroutine running in loop, and wait again
	c.Broadcast()

	for i := 0; i < 10; i++ {
		<-ready
	}

	// skip the loop
	isReady = true
	c.Broadcast()
	fmt.Println("ready go....")

	time.Sleep(time.Second)

}
