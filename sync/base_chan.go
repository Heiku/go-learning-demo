package main

import "fmt"

func main() {
	var a []int = nil

	// channel make sure the operation before sending data is front of the operation after receiving data
	c := make(chan struct{})

	go func() {
		a = make([]int, 3)
		c <- struct{}{}
	}()

	<-c
	fmt.Println(a[0], a[1], a[2]) // safe
}
