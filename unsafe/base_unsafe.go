package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func createInt() *int {
	return new(int)
}

// uintptr is a int value, don't point at anything, just store obj actual address, so is can be calculate
func main() {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x) // 16 32
	fmt.Println(M, N)

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16

	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y)          // y -> int, p1 -> int, with same point side
	var p2 = uintptr(unsafe.Pointer(z)) // uintptr() means z -> int will not be used any more, it may be recycled

	// uintptr can be calculated
	p2 += 2
	p2--
	p2--

	runtime.GC()
	*p0 = 1
	*(*int)(p1) = 2
	*(*int)(unsafe.Pointer(p2)) = 3

	runtime.KeepAlive(z) // make sure z be used

	xx := 123                // int
	p := unsafe.Pointer(&xx) // unsafe.Pointer
	pp := &p                 // &unsafe.Pointer
	p = unsafe.Pointer(pp)
	pp = (*unsafe.Pointer)(p)
}
