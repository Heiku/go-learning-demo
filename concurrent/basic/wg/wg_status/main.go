package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			wg.Done()
			wg.Add(1)
		}
	}()

	go func() {
		for {
			wg.Wait()
			fmt.Println("go2 free")
		}
	}()

	select {}
}
