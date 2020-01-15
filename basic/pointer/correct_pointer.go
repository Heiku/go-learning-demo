package main

import "fmt"

// pointer don't participate in calculate
func main() {
	a := int64(5)
	p := &a

	// can't pass compile
	// p++
	// p = (&a) + 8

	*p++                 // actual value ++
	fmt.Println(*p, a)   // 6 6
	fmt.Println(p == &a) // true, point to same value

	*&a++
	*&*&a++
	**&p++
	*&*p++
	fmt.Println(*p, a) // 10 10
}
