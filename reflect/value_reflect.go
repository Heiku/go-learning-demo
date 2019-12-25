package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false

	vn := vp.Elem()
	fmt.Println(vn.CanSet(), vn.Addr()) // true 0xc00000a0c8
	vn.Set(reflect.ValueOf(789))
	vn.SetInt(222)
	fmt.Println(n) // 222

	// pro
	var s struct {
		x interface{}
		y interface{}
	}
	vp = reflect.ValueOf(&s)
	v := reflect.Indirect(vp) // return the value of pointer

	vx, vy := v.Field(0), v.Field(1)
	fmt.Println(vx.CanSet(), vy.CanAddr())
	fmt.Println(vy.CanSet(), vy.CanAddr())

	vb := reflect.ValueOf(123)
	vx.Set(vb)
	fmt.Println(s)
	fmt.Println(vx.IsNil(), vy.IsNil())
}
