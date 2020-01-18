package main

import (
	"fmt"
	"time"
)

func receive() {
	ch := make(chan int, 100)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// closed() -> still can take out element(initial value, ok = false)
	close(ch)

	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		time.Sleep(time.Second)
	}
}

func main() {
	receive()
}
