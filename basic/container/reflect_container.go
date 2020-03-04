package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s))

	reflect.ValueOf(&s).Elem().SetLen(3)
	fmt.Println(len(s), cap(s))

	// can't out of the cap when we make the slice
	reflect.ValueOf(&s).Elem().SetCap(7)
	fmt.Println(len(s), cap(s))
}
