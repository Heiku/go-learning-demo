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

	reflect.ValueOf(&s).Elem().SetLen(5)
	fmt.Println(len(s), cap(s))
}
