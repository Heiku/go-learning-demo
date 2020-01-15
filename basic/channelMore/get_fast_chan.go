package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Int31()+1
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ { // simulate 5 coroutine put data
		go source(c)
	}

	rnd := <-c // just take one
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
