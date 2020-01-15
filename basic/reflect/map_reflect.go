package main

import (
	"fmt"
	"reflect"
)

func main() {
	valueOf := reflect.ValueOf
	fmt.Printf("%T valueOf", valueOf)

	m := map[string]int{
		"Unix":    1973,
		"Windows": 1985,
	}
	v := valueOf(m)

	// when value = nil, is means delete that entry
	v.SetMapIndex(valueOf("Windows"), reflect.Value{})
	v.SetMapIndex(valueOf("Linux"), valueOf(1991))
	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t: ", i.Value())
	}
}
