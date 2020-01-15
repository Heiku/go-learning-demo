package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go dosomething(100, &wg)
	go dosomething(110, &wg)
	go dosomething(120, &wg)
	go dosomething(130, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
	wg.Done()
}

func dosomething2(millisecs time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
}
