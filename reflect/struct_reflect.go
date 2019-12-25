package main

import (
	"fmt"
	"reflect"
)

type F func(string, int) bool

func (f F) Validate(s string) bool {
	return f(s, 32)
}

func main() {
	var x struct {
		N int
		f F
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // 2
	fmt.Println(tx.Field(1).Name) // f

	tf := tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2, 1
	fmt.Println(tf.NumMethod())          // 1
	fmt.Println(tf.Method(0).Name)       // Validate

	ts, ti, tb := tf.In(0), tf.In(1), tf.Out(0) // string int bool
	fmt.Println(ts.Kind(), ti.Kind(), tb.Kind())
}
