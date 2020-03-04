package main

import (
	"fmt"
	"time"
)

// chan T : a dequeue
// chan<- T: only enter chan, don't take out
// <-chan T: only take out, don't enter

// channel
// 1.wait queue: all coroutine are waiting for data in blocking state
// 2.send queue: al coroutine are waiting for sending data in blocking state
// 3.data buffer queue: circle queue storing actual data, it has cap and len

// sumw
// 1. if chan is closed, wait and send queue are all empty, but buffer queue may has data
// 2. in anytime, if buffer queue isn't empty, wait queue must be empty(still have data)
// 3. in anytime, if buffer queue isn't full, send queue must be empty

// most important is the mutex lock for racing data
func main() {
	c := make(chan int)

	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		fmt.Println("producer wake up and put data")

		ch <- x * x // block
	}(c, 3)

	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch // block, wait in wait queue, wait for data in buffer queue
		fmt.Println(n)

		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)

	<-done // main coroutine block, wait data in chan
	fmt.Println("byte")
}
