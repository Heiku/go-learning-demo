package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// string reflect.StringHeader
	// type StringHeader struct {
	//    Data uintptr (actual string arr)
	//    Len  int		(len of arr)
	//}

	s := "hello world"
	hello := s[:5]
	world := s[6:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println(hello + world)
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
}
