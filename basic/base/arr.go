package main

import "fmt"

func main() {
	a := [3]int{1, 3, 5}
	b := &a
	c := a
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	a[0] = 0

	for i := range b {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}

	for i := range c {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
}
