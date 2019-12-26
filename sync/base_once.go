package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	x := 0
	doSomething := func() {
		x++
		log.Println("Hello")
	}

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// once make sure method be invoked once
			once.Do(doSomething)

			log.Println("world!")
		}()
	}
	wg.Wait()
	log.Println("x = ", x)
}
