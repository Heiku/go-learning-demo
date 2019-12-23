package main

import "fmt"

// memory address usually Hex(0x1234CDEF) express.

// what is pointer?
// a pointer can store a memory value
// in fact, we see a pointer as a memory address or a memory address as pointer

func main() {
	p0 := new(int)   // new(type) return *Type, so new(int) return a pointer point to 0 value in memory
	fmt.Println(p0)  // p0 is a pointer(memory address), so print 0xc00000a0c8
	fmt.Println(*p0) // 0

	x := *p0              // x is a copy of 'p0 reference value', copy a other 0 value diff from p0
	p1, p2 := &x, &x      // p1 -> x, p2 -> x
	fmt.Println(p1 == p2) // point to same value, is true
	fmt.Println(p0 == p1) // p0 -> true 0 value, p1 -> copy 0 value(x), is false

	p3 := &*p0            // p3 -> true 0 value
	fmt.Println(p0 == p3) // point to same value, is true

	*p0, *p1 = 123, 789      // p0 -> 123, p1 -> 789
	fmt.Println(*p2, x, *p3) // 789 789 123 , x depend on p0, if p0 changed, x changed.

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}
