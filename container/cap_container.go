package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(len(a), cap(a))

	var s []int // define 0
	fmt.Println(len(s), cap(s))

	s, s2 := []int{2, 3, 5}, []bool{}
	fmt.Println(len(s), cap(s), len(s2), cap(s2))

	var m map[int]bool // define 0
	fmt.Println(len(m))

	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m), len(m2))

}
