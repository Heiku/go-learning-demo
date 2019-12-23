package main

import "fmt"

func main() {

	// define map[string]T
	type T struct {
		age int
	}
	mt := map[string]T{}
	mt["John"] = T{age: 29}

	// define map[int]int[]
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789}

	// wrong change
	// it only can be change whole value, can't be changed alone
	// mt["John"].age = 30
	// ma[1][1] = 123
	t := mt["John"] // use tmp variable
	t.age = 30
	mt["John"] = t // whole changed

	a := ma[1]
	a[1] = 123
	ma[1] = a // whole change

	// but it can be access for reading
	fmt.Println(ma[1][1])
	fmt.Println(mt["John"].age)
}

// this function may cause memory leak, the front of slice not be used, and no pointer point to it
func f() []int {
	s := make([]int, 10, 100)
	return s[50:60]
}
