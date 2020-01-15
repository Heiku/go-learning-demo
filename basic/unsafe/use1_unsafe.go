package main

import (
	"fmt"
	"unsafe"
)

// transform type of value with unsafe, it can ignore the compiler
func main() {
	type MyString string
	ms := []MyString{"C", "C++", "Go"}
	fmt.Printf("%s\n", ms)

	// ss := ([]string)ms	 compile failed
	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Rust"
	fmt.Printf("%s\n", ms)

	ms = *(*[]MyString)(unsafe.Pointer(&ss))
	fmt.Printf("%s\n", ms)
}
