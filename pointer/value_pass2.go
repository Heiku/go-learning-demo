package main

import "fmt"

// safe return value, return pointer, user can change it actual value by pointer
func newInt() *int {
	a := 3
	return &a
}

func doubleX(x *int) {
	*x += *x
	x = nil // set pointer = nil, but change just a pointer duplicate
}

func main() {
	var a = 3
	doubleX(&a)
	fmt.Printf("a = %d\n", a)

	p := &a // p -> a
	doubleX(p)
	fmt.Println(a, p == nil) // it don't influence outside param
}
