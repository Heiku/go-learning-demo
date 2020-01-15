package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now().UnixNano()
	a, b := longTimeRequest(), longTimeRequest() // two coroutine are working on concurrent
	fmt.Println(<-a, <-b)
	fmt.Println(time.Now().UnixNano() - start) // about 3s, it depend on the slower coroutine
}
