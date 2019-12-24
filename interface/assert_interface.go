package main

import "fmt"

type Writer interface {
	Write2(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write2(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	var w Writer
	var ok bool

	w, ok = x.(Writer)
	fmt.Println(w, ok)
	x, ok = w.(interface{})
	fmt.Println(x, ok)

	w, ok = y.(Writer)
	fmt.Println(w, ok)
	w = y.(Writer) // panic
}
