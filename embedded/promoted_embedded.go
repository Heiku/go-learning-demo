package main

import (
	"fmt"
	"reflect"
)

type F func(int) bool

func (f F) Validate(m int) bool {
	return f(m)
}

func (f *F) Modify(f2 F) {
	*f = f2
}

type B bool

func (b B) IsTrue() bool {
	return bool(b)
}

func (pb *B) Invert() {
	*pb = !*pb
}

type I interface {
	Load()
	Save()
}

func PrintTypeMethods(t reflect.Type) {
	fmt.Println(t, " has ", t.NumMethod(), " methods: ")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, " : ", t.Method(i).Name, "\n")
	}
}

func main() {
	var s struct {
		F
		*B
		I
	}

	PrintTypeMethods(reflect.TypeOf(s))
	fmt.Println()
	PrintTypeMethods(reflect.TypeOf(&s))
}
