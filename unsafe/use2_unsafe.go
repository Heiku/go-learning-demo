package main

import (
	"fmt"
	"unsafe"
)

// use uintptr for debug, quick find the obj memory address
func main() {
	type T struct{}
	var t T
	fmt.Println(&t)
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // 596c18
	println(&t)                                     // 0x596c18
}
