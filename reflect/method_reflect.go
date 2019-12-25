package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A, B int
}

func (t T) AddAndThenScale(n int) (int, int) {
	return n * (t.A + t.B), n * (t.A - t.B)
}

func main() {
	t := T{5, 2}
	vt := reflect.ValueOf(t)
	vm := vt.MethodByName("AddAndThenScale")

	// param is []Value, return []Value
	results := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(results[0].Int(), results[1].Int()) // 21 9

	neg := func(x int) int {
		return -x
	}
	vf := reflect.ValueOf(neg)
	fmt.Println(vf.Call(results[:1])[0].Int()) // -21
	fmt.Println(vf.Call([]reflect.Value{
		vt.FieldByName("A"),
	})[0].Int()) // -5
}
