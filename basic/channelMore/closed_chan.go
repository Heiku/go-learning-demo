package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case v := <-ch:
		fmt.Println(v)
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c))

	close(c)
	fmt.Println(IsClosed(c))

	_, ok := <-c
	fmt.Println(ok)
}
