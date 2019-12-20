package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var latch sync.WaitGroup

func sayGreetings(greeting string, times int){
	for i := 0; i < times; i++{
		log.Println(greeting)

		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	latch.Done()
}

func main() {
	latch.Add(2)

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	go sayGreetings("hi!", 10)
	go sayGreetings("hello!", 10)

	latch.Wait()
}
